package middleware

import (
	"MaintenanceSystem/pkg"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		tokenStr := c.GetHeader("Authorization")
		if tokenStr == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "未登录"})
			return
		}

		tokenStr = strings.TrimPrefix(tokenStr, "Bearer ")

		claims, err := pkg.ParseToken(tokenStr)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "token无效"})
			return
		}

		// 存用户信息
		c.Set("user_id", claims["user_id"])

		c.Next()
	}
}
