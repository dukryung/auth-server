// Package auth to implement for authentication server.
package auth

import (
	"github.com/dukryung/microservice/server/types"
	"github.com/gin-gonic/gin"
	"github.com/dukryung/microservice/server/auth/rest"
)

type Server struct {
	router      *gin.Engine
	restHandler *rest.Handler
}

func NewServer() *Server {
	s := &Server{}

	router := gin.Default()
	s.router = router
	s.restHandler = rest.NewHandler()
	HandleManager := types.NewHandlerManager(s.restHandler)
	HandleManager.RegisterRoute(s.router)

	return s
}

func (s *Server) Run() {
	s.router.Run(":13579")
}

func (s *Server) Close() {
	s.Close()
}
