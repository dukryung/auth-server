package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/dukryung/microservice/server/auth/rest/endpoint"
)

type Handler struct {
}

func NewHandler() *Handler {
	h := &Handler{}

	return h
}

func (h *Handler) RegisterRoute(router *gin.Engine) {
	router.GET("/seed", endpoint.GetMnemonic)

}
