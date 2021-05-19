package controllers

import (
	"password-share-server-golang/src/models"

	"github.com/gin-gonic/gin"
)

type CreateSecurePassword struct {
	UserId        string `json:"user_id"`
	EncryptedData string `json:"encrypted_data"`
}

type SecurePasswordHandler struct {
	userRepo           models.UserRepository
	securePasswordRepo models.SecurePasswordRepository
}

func NewSecurePasswordController(userRepo models.UserRepository, securePasswordRepo models.SecurePasswordRepository) *SecurePasswordHandler {
	return &SecurePasswordHandler{
		userRepo:           userRepo,
		securePasswordRepo: securePasswordRepo,
	}
}

func (h *SecurePasswordHandler) CreateSecurePasswordRoute(c *gin.Context) {
	var form CreateSecurePassword

	c.BindJSON(&form)

	if form.UserId == "" {
		c.JSON(400, gin.H{"message": "You must send a user_id field"})
		return
	}

	if form.EncryptedData == "" {
		c.JSON(400, gin.H{"message": "You must send an encrypted_data field"})
		return
	}

	_, err := h.userRepo.FindByID(form.UserId)

	if err != nil {
		c.JSON(400, gin.H{"message": "User id not found"})
		return
	}

	newSecurePassword, err := h.securePasswordRepo.CreateSecurePassword(form.UserId, form.EncryptedData)

	if err != nil {
		c.JSON(500, gin.H{"status": "error", "message": "An error occurred while creating the securePassword", "err": err})
		return
	}

	c.JSON(200, gin.H{"status": "success", "message": "congrats", "newSecurePassword": newSecurePassword})
}
