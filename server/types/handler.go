package types

import "github.com/gin-gonic/gin"

type Handler interface {
	Run()
	RegisterRoute(engine *gin.Engine)
}

type HandlerManager struct {
	handlers []Handler
}

func NewHandlerManager(handlers ...Handler) *HandlerManager {
	handleManager := &HandlerManager{
		handlers: handlers,
	}

	return handleManager
}

func (m *HandlerManager) Run() {
	for _, handler := range m.handlers {
		handler.Run()
	}
}

func (m *HandlerManager) RegisterRoute(router *gin.Engine) {
	for _, handler := range m.handlers {
		handler.RegisterRoute(router)
	}
}
