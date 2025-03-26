package controller

import (
	"FinalGo/config"
	"FinalGo/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SearchProducts ค้นหาสินค้าตามรายละเอียดและช่วงราคา
func SearchProducts(c *gin.Context) {
	var products []models.Product
	minPrice := c.DefaultQuery("min_price", "0")
	maxPrice := c.DefaultQuery("max_price", "1000000")

	if err := config.DB.Where("price BETWEEN ? AND ?", minPrice, maxPrice).Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch products"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"products": products})
}
