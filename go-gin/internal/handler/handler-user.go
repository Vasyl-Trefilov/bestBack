package handler

import (
	models "github.com/Vasyl-Trefilov/bestBack/internal"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) ListUsers(ctx *gin.Context) {
	// list users
}

func (h *Handler) CreateUser(ctx *gin.Context) {
	name := uuid.New().String()
	email := name + "@example.com"

	user := &models.User{
		Name:  name,
		Email: email,
	}

	user, err := h.srv.CreateUser(user)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(201, user)
}

func (h *Handler) GetUser(ctx *gin.Context) {
	// get user by id
}

func (h *Handler) UpdateUser(ctx *gin.Context) {
	// update user by id
}

func (h *Handler) DeleteUser(ctx *gin.Context) {
	// delete user by id
}
