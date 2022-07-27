package rest

import (
	"database/sql"
	"fmt"
	"github.com/dukryung/microservice/server/auth/rest/endpoint"
	"github.com/dukryung/microservice/server/types/configs"
	"github.com/dukryung/microservice/server/types/log"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	EndPoint   *endpoint.EndPoint
	logger     *log.Logger
	router     *gin.Engine
	authConfig *configs.AuthConfig
}

func NewHandler(router *gin.Engine, authConfig *configs.AuthConfig, db *sql.DB) *Handler {
	h := &Handler{
		logger:     log.NewLogger("handler", authConfig.Log),
		router:     router,
		authConfig: authConfig,
	}
	h.EndPoint = endpoint.NewEndPoint(authConfig.Log, db)

	return h
}

func (h *Handler) Run() {
	err := h.router.Run(fmt.Sprintf(":%d", h.authConfig.Port))
	if err != nil {
		h.logger.Err(err)
	}
}

func (h *Handler) RegisterRoute(router *gin.Engine) {
	router.GET("/seed", h.EndPoint.GetMnemonic)
	router.POST("/account/register", h.EndPoint.RegisterAccount)
	router.GET("/account/login", h.EndPoint.LoginAccount)
	router.PUT("/account/import", h.EndPoint.ImportAccount)
}
