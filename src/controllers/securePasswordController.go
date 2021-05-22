package controllers

import (
	"password-share-server-golang/src/models"

	"github.com/gin-gonic/gin"
)

type CreateSecurePassword struct {
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

	if form.EncryptedData == "" {
		c.JSON(400, gin.H{"message": "You must send an encrypted_data field"})
		return
	}

	userContext, _ := c.Get("user")
	user := userContext.(*models.User)

	newSecurePassword, err := h.securePasswordRepo.CreateSecurePassword(user.ID, form.EncryptedData)

	if err != nil {
		c.JSON(500, gin.H{"status": "error", "message": "An error occurred while creating the securePassword", "err": err})
		return
	}

	c.JSON(200, gin.H{"status": "success", "message": "congrats", "newSecurePassword": newSecurePassword})
}

func (h *SecurePasswordHandler) GetUsersSecurePasswords(c *gin.Context) {
	userContext, _ := c.Get("user")
	user := userContext.(*models.User)

	passwords, err := h.securePasswordRepo.FindAllByUserId(user.ID)

	if err != nil {
		c.JSON(500, gin.H{"status": "error", "message": "Error occurred while fetching user's passwords", "err": err})
	}

	c.JSON(200, gin.H{"status": "success", "message": "congrats", "securePasswords": passwords})
}
