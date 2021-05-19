package controllers

import (
	"password-share-server-golang/src/models"
	"password-share-server-golang/src/util"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type BaseHandler struct {
	userRepo models.UserRepository
}

type RegisterStruct struct {
	Username string
	Password string
}

func (h *BaseHandler) RegisterRoute(c *gin.Context) {
	var form RegisterStruct

	c.BindJSON(&form)

	if form.Username == "" {
		c.JSON(400, gin.H{"message": "You must send a username"})
		return
	}

	if form.Password == "" {
		c.JSON(400, gin.H{"message": "You must send a password"})
		return
	}

	foundUser, _ := h.userRepo.FindByUsername(form.Username)

	if foundUser != (models.User{}) {
		c.JSON(400, gin.H{"message": "Username already exists"})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(form.Password), 14)

	if err != nil {
		c.JSON(400, gin.H{"message": "Fatal error hashing password", "err": err})
		return
	}

	hashedPassword := string(hash)

	newUser, err := h.userRepo.CreateUser(form.Username, hashedPassword)

	if err != nil {
		c.JSON(500, gin.H{"status": "error", "message": "Error occured while creating and inserting user", "error": err})
		return
	}

	tokenString, err := util.CreateJWTTokenString(newUser)

	if err != nil {
		c.JSON(500, gin.H{"message": "Error creating JWT token", "err": err})
		return
	}

	c.JSON(200, gin.H{"status": "success", "message": "congrats", "user": newUser, "jwt": tokenString})
}

func NewUserController(userRepo models.UserRepository) *BaseHandler {
	return &BaseHandler{
		userRepo: userRepo,
	}
}

func (h *BaseHandler) SuccessRoute(c *gin.Context) {
	c.JSON(200, gin.H{"status": "success"})
}
