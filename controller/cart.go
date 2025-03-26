package controller

import (
	"FinalGo/config"
	"FinalGo/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CartItemResp represents the response for cart items
type CartItemResp struct {
	ProductName string  `json:"product_name"`
	Quantity    int     `json:"quantity"`
	Price       float64 `json:"price"`
	TotalPrice  float64 `json:"total_price"`
}

// GetAllCarts retrieves all carts for a specific customer
func GetAllCarts(c *gin.Context) {
	customerID := c.DefaultQuery("customer_id", "")
	if customerID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "customer_id is required"})
		return
	}

	var carts []models.Cart
	if err := config.DB.Where("customer_id = ?", customerID).Preload("CartItems.Product").Find(&carts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch carts"})
		return
	}

	// Define response structure
	type CartResponse struct {
		CartName string         `json:"cart_name"`
		Items    []CartItemResp `json:"items"`
	}

	var cartResponses []CartResponse

	// Process cart items and calculate prices
	for _, cart := range carts {
		var items []CartItemResp
		for _, item := range cart.CartItems {
			totalPrice := float64(item.Quantity) * item.Product.Price
			items = append(items, CartItemResp{
				ProductName: item.Product.ProductName, // Changed from Name to ProductName
				Quantity:    item.Quantity,
				Price:       item.Product.Price,
				TotalPrice:  totalPrice,
			})
		}
		cartResponses = append(cartResponses, CartResponse{
			CartName: cart.CartName, // Changed from Name to CartName
			Items:    items,
		})
	}

	c.JSON(http.StatusOK, gin.H{"carts": cartResponses})
}
