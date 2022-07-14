// Package auth to implement for authentication server.
package auth

import (
	"github.com/dukryung/microservice/server/types"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine

}

func NewServer() *Server {
	s := &Server{}

	router := gin.Default()
	s.router = router

	HandleManager := types.NewHandlerManager()
	HandleManager.RegisterRoute(s.router)

	return s
}

func (s *Server) Run() {
	s.router.Run(":13579")
}

func (s *Server) Close() {
	s.Close()
}

