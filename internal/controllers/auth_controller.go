package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/phantom5803/ticket-booking-project/internal/models"
	"github.com/phantom5803/ticket-booking-project/internal/utils"
	"github.com/phantom5803/ticket-booking-project/pkg/database"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var req models.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	hashed, _ := bcrypt.GenerateFromPassword([]byte(req.Password), 10)

	user := models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: string(hashed),
	}

	database.DB.Create(&user)

	c.JSON(200, gin.H{"message": "User created"})
}

func Login(c *gin.Context) {
	var req models.RegisterRequest

	c.ShouldBindJSON(&req)

	var user models.User
	database.DB.Where("email = ?", req.Email).First(&user)

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		c.JSON(401, gin.H{"error": "Invalid credentials"})
		return
	}

	token, _ := utils.GenerateToken(user.ID)

	c.JSON(200, gin.H{"token": token})
}
