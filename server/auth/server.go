// Package auth to implement for authentication server.
package auth

import (
	"fmt"
	"github.com/dukryung/microservice/server/auth/rest"
	"github.com/dukryung/microservice/server/types"
	"github.com/dukryung/microservice/server/types/configs"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router      *gin.Engine
	restHandler *rest.Handler
	authConfig  *configs.AuthConfig
}

func NewServer(authConfig *configs.AuthConfig) *Server {
	s := &Server{
		authConfig: authConfig,
	}

	router := gin.Default()
	s.router = router

	database, err := authConfig.DB.GetDBConnection()
	if err != nil {
		fmt.Println("err : ", err)
	}

	s.restHandler = rest.NewHandler(database)
	HandleManager := types.NewHandlerManager(s.restHandler)
	HandleManager.RegisterRoute(s.router)

	return s
}

func (s *Server) Run() {
	s.router.Run(fmt.Sprintf(":%d",s.authConfig.Port))
}

func (s *Server) Close() {
	s.Close()
}
