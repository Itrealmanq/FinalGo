package controller

import (
	"FinalGo/config"
	"FinalGo/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProfile(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(http.StatusOK, user)
}

func UpdateAddress(c *gin.Context) {
	var request struct {
		Address string `json:"address"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	user, _ := c.Get("user")
	u := user.(models.User)

	config.DB.Model(&u).Update("address", request.Address)

	c.JSON(http.StatusOK, gin.H{"message": "Address updated", "address": request.Address})
}
