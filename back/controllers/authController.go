package controllers

import (
	"back/db"
	"back/models"
	"back/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct{}

var AuthControllerInstance = &AuthController{}

func (a *AuthController) Register(c *gin.Context) {
	log.Println("Register endpoint hit")
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// Check if username or email already exists
	var existingUser models.User
	if err := db.DB.Where("username = ?", user.Username).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"message": "Username already exists"})
		return
	}

	if err := db.DB.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"message": "Email already exists"})
		return
	}

	if err := user.HashPassword(user.Password); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error hashing password"})
		return
	}

	// Save user to the database
	if err := db.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error saving user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func (a *AuthController) Login(c *gin.Context) {
	log.Println("Login endpoint hit")
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	// Find user by username
	if err := db.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Username doesn't exist"})
		return
	}

	if err := user.CheckPassword(input.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Password doesn't match"})
		return
	}

	token, err := utils.GenerateToken(user.ID, user.Username, user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error generating token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (a *AuthController) FindUser(c *gin.Context) {
	log.Println("FindUser endpoint hit")

	var users []models.User

	// Find all users in the database
	if err := db.DB.Find(&users).Error; err != nil {
		log.Println("Error finding users:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users"})
		return
	}

	// Create a slice to hold the users' data that will be returned
	var userResponses []gin.H

	// Loop through the users and append their details (excluding sensitive fields like password)
	for _, user := range users {
		userResponse := gin.H{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
			// Add other fields you want to expose
		}
		userResponses = append(userResponses, userResponse)
	}

	// Return the user data as a JSON response
	c.JSON(http.StatusOK, gin.H{
		"users": userResponses,
	})
}
