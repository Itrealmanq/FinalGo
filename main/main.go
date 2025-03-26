package main

import (
	"FinalGo/config"
	"FinalGo/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitDB()
	r := gin.Default()

	// Route สำหรับการดึงข้อมูลรถเข็นทั้งหมด
	r.GET("/carts", controller.GetAllCarts)

	r.Run(":8080")
}
