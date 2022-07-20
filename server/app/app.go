package app

import (
	"fmt"
	"github.com/dukryung/microservice/server/auth"
	"github.com/dukryung/microservice/server/types"
	"github.com/dukryung/microservice/server/types/configs"
	_ "github.com/lib/pq"
)

type App struct {
	servers []types.Server
}

func NewApp() *App {
	app := &App{}

	authConfig, err := configs.LoadAuthConfig()
	if err != nil {
		fmt.Println("err : ", err)
	}

	authServer := auth.NewServer(authConfig)

	app.servers = append(app.servers, authServer)

	return app
}

func (app *App) RunServer() error {
	if len(app.servers) == 0 {
		return fmt.Errorf("empty servers")
	}

	for _, server := range app.servers {
		server.Run()
	}

	return nil
}

func (app *App) CloseServer() {
	for _, server := range app.servers {
		server.Close()
	}
}
