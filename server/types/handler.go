package types

import "github.com/gin-gonic/gin"

type Handler interface {
	RegisterRoute(engine *gin.Engine)
}

type HandlerManager struct{
	handlers []Handler
}

func NewHandlerManager(handlers ...Handler) *HandlerManager {
	handleManager := &HandlerManager{}

	return handleManager
}

func (m *HandlerManager) Run() {

}

func (m *HandlerManager) RegisterRoute(router *gin.Engine) {
	for _, handler := range m.handlers {
		handler.RegisterRoute(router)
	}

}