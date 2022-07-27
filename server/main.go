package main

import (
	"github.com/dukryung/microservice/server/app"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	app := app.NewApp()
	quit := make(chan os.Signal)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	err := app.RunServer()
	if err != nil {
		panic(err)
	}

	<-quit
	app.CloseServer()

}
