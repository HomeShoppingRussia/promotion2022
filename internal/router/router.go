package router

import (
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"hsr/loto/internal/handler"
	"hsr/loto/internal/logger"
	"hsr/loto/pkg/myhttp"
	"hsr/loto/pkg/swagger"
	"net/http"
)

type Router struct {
	hh  *handler.Healthcheck
	sh  *swagger.Swagger
	hth *myhttp.MyHttp
}

func NewRouter(hh *handler.Healthcheck, sh *swagger.Swagger) *Router {
	return &Router{
		hh: hh,
		sh: sh,
	}
}

func (ro *Router) Get() *runtime.ServeMux {
	router := runtime.NewServeMux()

	/*err := router.HandlePath(http.MethodGet, "/api/getDraw2", func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		fmt.Println("123")
		ro.hth.GetHandler(w, r)
		//http.StripPrefix("/", http.FileServer(http.Dir("./third_party/swagger-ui"))).ServeHTTP(w, r)
	})
	if err != nil {
		logger.Error.Fatal(err)
	}*/

	err := router.HandlePath(http.MethodGet, "/health", func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		ro.hh.HealthcheckHandler(w, r)
	})
	if err != nil {
		logger.Error.Fatal(err)
	}

	err = router.HandlePath(http.MethodGet, "/swagger.json", func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		ro.sh.GetSwaggerJsonHandler(w, r)
	})
	if err != nil {
		logger.Error.Fatal(err)
	}

	/*err = router.HandlePath(http.MethodGet, "/api2", func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		ro.hth.GetHandler(w, r)
	})
	if err != nil {
		logger.Error.Fatal(err)
	}*/

	err = router.HandlePath(http.MethodGet, "/swagger-ui/*", func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		http.StripPrefix("/swagger-ui/", http.FileServer(http.Dir("./third_party/swagger-ui"))).ServeHTTP(w, r)
	})
	if err != nil {
		logger.Error.Fatal(err)
	}

	return router
}
