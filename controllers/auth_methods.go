package controllers

import (
	"RestLecture/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// RegisterUser POST /auth/register
func RegisterUser(c *gin.Context) {
	var input models.RegisterUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create user
	user := models.User{
		Username: input.Username,
		Password: input.Password,
	}

	models.DB.Create(&user)

	c.JSON(http.StatusCreated, gin.H{
		"id":       user.ID,
		"username": user.Username,
	})
}
