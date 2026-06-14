package handler
<<<<<<< HEAD

import (
	"github.com/Vasyl-Trefilov/bestBack/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	srv *service.Service
}

func NewHandler(srv *service.Service) *Handler {
	return &Handler{
		srv: srv,
	}
}

func (h *Handler) HelloWorld(ctx *gin.Context) {
	ctx.String(200, "Hello, World!")
}

func (h *Handler) HealthCheck(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"status": "ok"})
}

func (h *Handler) HaavyRoute(ctx *gin.Context) {
	// heavy math
}
=======
>>>>>>> bd3a36c591bd590609f477457d6e266c02a6b81c
