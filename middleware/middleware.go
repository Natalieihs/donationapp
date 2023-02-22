package middleware

import "github.com/gin-gonic/gin"

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//实现用户登录认证逻辑
		c.Next()
	}
}
