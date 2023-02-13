package application

import (
	"context"
	//"context"
	//es "hsr/loto/api/proto/external/external_service"
	"hsr/loto/internal/api"
	"hsr/loto/internal/config"
	"hsr/loto/internal/db"
	"hsr/loto/internal/handler"
	"hsr/loto/internal/logger"
	"hsr/loto/internal/repository"
	"google.golang.org/grpc"
	//"hsr/loto/internal/logger"
	"hsr/loto/internal/router"
	"hsr/loto/internal/server"
	"hsr/loto/internal/service"
	"hsr/loto/pkg/swagger"
	//"google.golang.org/grpc/credentials/insecure"
	"time"
)

type Application struct {
	DB  *db.DB
	Cfg *config.Config
	Api *api.Api
	//Es  *es.ExternalServiceClient
	Ext Ext
}

type Ext struct {
	C *grpc.ClientConn
}

func Get() (*Application, error) {
	cfg := config.Get()
	conn, err := db.Get(context.Background(), cfg.DSN)
	if err != nil {
		logger.Error.Print(err)
	}

	//conn, err := grpc.DialContext(
	//	context.Background(),
	//	cfg.NotificationServiceUrl,
	//	grpc.WithBlock(),
	//	grpc.WithTransportCredentials(insecure.NewCredentials()))
	//if err != nil {
	//	logger.Error.Fatalln("Failed to dial server:", err)
	//}

	//ext := &Ext{
	//	Ns: conn,
	//}

	//c := ns.NewExternalServiceClient(ext.C)

	tr := repository.NewTicketRepository(conn)
	pr := repository.NewPrizeRepository(conn)
	ts := service.NewTicket(tr)
	ps := service.NewPrize(pr)

	hh := handler.GetHealthHandler(ts)
	hh.SetApplicationState("1.0.0", time.Now().Format("2006-01-02 15:04:05"))

	s := &server.Server{
		TicketService: ts,
		PrizeService:  ps,
		//ExternalServiceClient: c,
	}

	ro := router.NewRouter(hh, swagger.GetSwaggerHandler(cfg.SwaggerPath))

	a := api.Get(cfg, s, ro.Get())

	return &Application{
		DB:  nil,
		Cfg: cfg,
		Api: a,
	}, nil
}

func (a *Application) Stop() error {
	//defer a.Ext.C.Close()

	return nil
}
