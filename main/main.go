package main

import (
	model "FinalGo/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func connectDatabase() (*gorm.DB, error) {
	dsn := "cp_65011212022:65011212022@csmsu@tcp(202.28.34.197:3306)/cp_65011212157?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func main() {
	r := gin.Default()

	db, err := connectDatabase()
	if err != nil {
		panic("Failed to connect to database!")
	}

	r.POST("/login", func(c *gin.Context) {
		var input struct {
			Email    string `json:"email" binding:"required"`
			Password string `json:"password" binding:"required"`
		}
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		var customer model.Customer
		if err := db.Where("email = ?", input.Email).First(&customer).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
			return
		}

		var passwordMatch bool
		if len(customer.Password) < 60 {
			if customer.Password == input.Password {
				passwordMatch = true
			}
		} else {

			err := bcrypt.CompareHashAndPassword([]byte(customer.Password), []byte(input.Password))
			if err == nil {
				passwordMatch = true
			}
		}
		if passwordMatch {
			if len(customer.Password) < 60 {
				hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
					return
				}
				customer.Password = string(hashedPassword)
				if err := db.Save(&customer).Error; err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update password"})
					return
				}
			}
			c.JSON(http.StatusOK, gin.H{
				"CustomerID":  customer.CustomerID,
				"FirstName":   customer.FirstName,
				"LastName":    customer.LastName,
				"Email":       customer.Email,
				"PhoneNumber": customer.PhoneNumber,
				"Address":     customer.Address,
				"CreatedAt":   customer.CreatedAt.Format(time.RFC3339),
				"UpdatedAt":   customer.UpdatedAt.Format(time.RFC3339),
			})

		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		}
	})
	r.Run(":8080")
}
