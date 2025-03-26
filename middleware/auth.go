package middleware

import (
	"FinalGo/config"
	"FinalGo/models"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4" // ใช้ /v4 แทน
)

// AuthMiddleware ตรวจสอบและยืนยัน JWT Token
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// อ่าน JWT Token จาก Header
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is missing"})
			c.Abort()
			return
		}

		// ตัด "Bearer " ออกจาก Token
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		// ตรวจสอบและยืนยัน JWT Token
		token, err := jwt.ParseWithClaims(tokenString, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
			// ตรวจสอบความถูกต้องของ token
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				// แก้ไขตรงนี้: หากไม่ใช่ SigningMethodHMAC ให้คืนค่า error ที่เป็นข้อความ
				return nil, fmt.Errorf("unexpected signing method %v", token.Header["alg"])
			}
			return []byte("secret_key"), nil // เปลี่ยนเป็น key ที่คุณใช้ในการเซ็น JWT
		})
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// ดึงข้อมูลผู้ใช้จาก JWT Token
		claims, ok := token.Claims.(*jwt.MapClaims) // Cast claims to *MapClaims
		if !ok || (*claims)["email"] == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}

		// ค้นหาผู้ใช้ในฐานข้อมูล
		email := (*claims)["email"].(string)
		var user models.User
		if err := config.DB.Where("email = ?", email).First(&user).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
			c.Abort()
			return
		}

		// ส่งข้อมูลผู้ใช้ไปยัง context
		c.Set("user", user)
		c.Next()
	}
}
