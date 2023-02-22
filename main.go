package main

import (
	"donationapp/controllers"
	"donationapp/middleware"
	"donationapp/models"

	"github.com/gin-gonic/gin"
)

func main() {
	err := models.InitDatabase()
	if err != nil {
		panic("failed to connect database")
	}

	r := gin.Default()

	// 注册中间件
	r.Use(middleware.AuthMiddleware())
	//注册API路由
	user := r.Group("/api/user")
	user.POST("/register", controllers.Register)
	user.POST("/login", controllers.Login)
	donation := r.Group("/api/donation")
	donation.POST("/new", controllers.NewDonation)
}
