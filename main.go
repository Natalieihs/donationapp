package main

import (
	"donationapp/controllers"
	"donationapp/models"

	"github.com/gin-gonic/gin"
)

func main() {
	err := models.InitDatabase()
	if err != nil {
		panic("failed to connect database")
	}

	r := gin.Default()
	//注册API路由
	user := r.Group("/api/user")
	user.POST("/register", controllers.Register)
	user.POST("/login", controllers.Login)
	donation := r.Group("/api/donation")
	donation.POST("/new", controllers.NewDonation)
}
