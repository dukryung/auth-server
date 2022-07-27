// Package auth to implement for authentication server.
package auth

import (
	"fmt"
	"github.com/dukryung/microservice/server/auth/rest"
	"github.com/dukryung/microservice/server/types"
	"github.com/dukryung/microservice/server/types/configs"
	"github.com/dukryung/microservice/server/types/log"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router      *gin.Engine
	restHandler *rest.Handler
	authConfig  *configs.AuthConfig
	logger      *log.Logger

	handlerManager *types.HandlerManager
}

func NewServer(authConfig *configs.AuthConfig) *Server {
	s := &Server{
		authConfig: authConfig,
		logger:     log.NewLogger("auth/server", authConfig.Log),
	}

	router := gin.Default()
	s.router = router

	database, err := authConfig.DB.GetDBConnection()
	if err != nil {
		fmt.Println("err : ", err)
	}

	s.restHandler = rest.NewHandler(router, authConfig, database)
	s.handlerManager = types.NewHandlerManager(s.restHandler)
	s.handlerManager.RegisterRoute(s.router)

	return s
}

func (s *Server) Run() {
	s.handlerManager.Run()
}

func (s *Server) Close() {
	s.Close()
}
