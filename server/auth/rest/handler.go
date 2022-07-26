package rest

import (
	"database/sql"
	"github.com/dukryung/microservice/server/auth/rest/endpoint"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	EndPoint *endpoint.EndPoint
}

func NewHandler(db *sql.DB) *Handler {
	h := &Handler{}
	h.EndPoint = endpoint.NewEndPoint(db)

	return h
}

func (h *Handler) Run() {

}

func (h *Handler) RegisterRoute(router *gin.Engine) {
	router.GET("/seed", h.EndPoint.GetMnemonic)
	router.POST("/account/register", h.EndPoint.RegisterAccount)
	router.GET("/account/login", h.EndPoint.LoginAccount)
	router.PUT("/account/import",h.EndPoint.ImportAccount)
}
