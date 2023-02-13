package main

import (
	"hsr/loto/internal/application"
	"hsr/loto/internal/exithandler"
	"hsr/loto/internal/logger"
)

func main() {
	app, err := application.Get()
	if err != nil {
		logger.Error.Fatal(err.Error())
	}

	go func() {
		logger.Info.Printf("starting GRPC server at %s", app.Cfg.GrpcPort)
		if err := app.Api.Run(); err != nil {
			logger.Error.Fatal(err.Error())
		}
	}()

	exithandler.Init(func() {
		if err := app.Stop(); err != nil {
			logger.Error.Println(err)
		}

		if err := app.Api.Close(); err != nil {
			logger.Error.Println(err.Error())
		}
	})
}
