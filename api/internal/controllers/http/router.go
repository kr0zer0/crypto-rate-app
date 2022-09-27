package http

import (
	"api/internal/usecases"
	"github.com/gin-gonic/gin"
)

//go:generate mockgen -source=router.go -destination=mocks/serviceMock.go

type Handler struct {
	useCases *usecases.UseCases
}

func NewHandler(useCases *usecases.UseCases) *Handler {
	return &Handler{
		useCases: useCases,
	}
}

func (h *Handler) InitRouter() *gin.Engine {
	router := gin.Default()

	base := router.Group("/api")
	base.GET("/rate", h.getCurrentExchangeRate)
	base.POST("/subscribe", h.subscribe)
	base.POST("/sendEmails", h.sendEmails)

	return router
}
