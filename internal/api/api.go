package api

import (
	"context"
	grpcPrometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	pb "hsr/loto/api/proto/loto"
	"hsr/loto/internal/config"
	"hsr/loto/internal/logger"
	"hsr/loto/internal/server"
	"net"
	"net/http"
	"time"
)

type Api struct {
	c  *config.Config
	s  *server.Server
	m  *runtime.ServeMux
	gs *grpc.Server
	hs *http.Server
}

type wrappedStream struct {
	grpc.ServerStream
	ctx context.Context
}

var errMissingMetadata = status.Errorf(codes.InvalidArgument, "no incoming metadata in rpc context")

func unaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errMissingMetadata
	}

	md.Append("Access-Control-Allow-Origin", "*")
	md.Append("123", "123")
	ctx = metadata.NewIncomingContext(ctx, md)

	return handler(ctx, req)
}

func streamInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	md, ok := metadata.FromIncomingContext(ss.Context())
	if !ok {
		return errMissingMetadata
	}

	md.Append("Access-Control-Allow-Origin", "*")
	md.Append("123", "123")
	ctx := metadata.NewIncomingContext(ss.Context(), md)

	return handler(srv, &wrappedStream{ss, ctx})
}

func Get(c *config.Config, s *server.Server, m *runtime.ServeMux) *Api {
	return &Api{
		c: c,
		s: s,
		m: m,
		gs: grpc.NewServer(
			grpc.StreamInterceptor(streamInterceptor),
			grpc.UnaryInterceptor(unaryInterceptor),
		),
		hs: &http.Server{
			Addr:         c.HttpPort,
			ErrorLog:     logger.Error,
			ReadTimeout:  30 * time.Second,
			WriteTimeout: 30 * time.Second,
		},
	}
}

func (a *Api) Run() error {
	lis, err := net.Listen("tcp", a.c.GrpcPort)
	if err != nil {
		logger.Error.Fatalln("Failed to listen:", err)
	}

	pb.RegisterLotoServer(a.gs, a.s)
	reflection.Register(a.gs)
	grpcPrometheus.Register(a.gs)
	logger.Info.Println("Serving gRPC on 0.0.0.0" + a.c.GrpcPort)

	go func() {
		err := a.gs.Serve(lis)
		if err != nil {
			logger.Error.Fatal(err)
		}
	}()

	conn, err := grpc.DialContext(
		context.Background(),
		"0.0.0.0"+a.c.GrpcPort,
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		logger.Error.Fatalln("Failed to dial server:", err)
	}

	err = pb.RegisterLotoHandler(context.Background(), a.m, conn)
	if err != nil {
		logger.Error.Fatalln("Failed to register gateway:", err)
	}

	a.hs.Handler = cors(a.m)

	logger.Info.Printf("Serving gRPC-Gateway on http://localhost%s\n", a.c.GrpcPort)
	logger.Info.Printf("Serving Swagger-UI on http://localhost%s/swagger-ui/", a.c.HttpPort)
	go func() {
		err := a.hs.ListenAndServe()
		if err != nil {
			logger.Error.Fatal(err)
		}
	}()

	return err
}

func (a *Api) Close() error {
	a.gs.GracefulStop()
	return a.hs.Close()
}

func cors(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PATCH,UPDATE,DELETE,OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-CSRF-Token, Authorization")

		h.ServeHTTP(w, r)
	})
}
