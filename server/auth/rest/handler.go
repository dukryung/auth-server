package rest

import (
	"github.com/dukryung/microservice/server/auth/rest/endpoint"
	"github.com/gin-gonic/gin"
)

type Handler struct {
}

func NewHandler() *Handler {
	h := &Handler{}

	return h
}

func(h *Handler) Run() {

}

func (h *Handler) RegisterRoute(router *gin.Engine) {
	router.GET("/seed", endpoint.GetMnemonic)
	router.POST("/account", endpoint.RegisterAccount)
}
