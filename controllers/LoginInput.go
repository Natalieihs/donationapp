package controllers

import (
	"donationapp/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LoginInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// 注册
func Register(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
}

// 登录
func Login(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := models.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(401, gin.H{"error": "User not found"})
		} else {
			c.JSON(500, gin.H{"error": "Server error"})
		}
		return
	}

	if input.Password != user.Password {
		c.JSON(401, gin.H{"error": "Password wrong"})
		return
	}
	//实现用户登录逻辑
}
func NewDonation(c *gin.Context) {
	//实现捐款逻辑
}

func GetDonation(c *gin.Context) {
	//实现获取捐款记录逻辑
}
