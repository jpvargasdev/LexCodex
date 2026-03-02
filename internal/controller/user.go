package controller

import (
	"lexcodex/internal/models"
	"lexcodex/internal/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Controller) DeleteUserController(c *gin.Context) {
	userUID, exists := c.Get("userUID")
	if !exists {
		c.JSON(401, gin.H{"error": "User not authenticated"})
		return
	}
	uid, ok := userUID.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user UID"})
		return
	}

	err := models.DeleteUser(uid)
	if err != nil {
		log.Printf("Error deleting user: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}
	c.JSON(http.StatusOK, "OK")
}

func (h *Controller) CreateUserController(c *gin.Context) {
	userUID, exists := c.Get("userUID")
	if !exists {
		c.JSON(401, gin.H{"error": "User not authenticated"})
		return
	}
	uid, ok := userUID.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user UID"})
		return
	}

	var newUser models.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate email
	if !utils.IsValidEmail(newUser.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email format"})
		return
	}

	newUser.ID = uid

	err := models.CreateUser(newUser)
	if err != nil {
		log.Printf("Error creating user: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}
	c.JSON(http.StatusCreated, "OK")
}
