package controllers

import (
	"password-share-server-golang/src/models"

	"github.com/gin-gonic/gin"
)

type BaseHandler struct {
	userRepo models.UserRepository
}

func NewUserController(userRepo models.UserRepository) *BaseHandler {
	return &BaseHandler{
		userRepo: userRepo,
	}
}

func (h *BaseHandler) SuccessRoute(c *gin.Context) {
	c.JSON(200, gin.H{"status": "success"})
}
