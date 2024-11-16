package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/viniciusfal/erp/services"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cookie, err := ctx.Cookie("auth_token")
		if err != nil {
			ctx.AbortWithStatusJSON(401, gin.H{"error": "authentication required"})
			return
		}

		if !services.NewJWTService().ValidateToken(cookie) {
			ctx.AbortWithStatusJSON(401, gin.H{"error": "invalid token"})
			return
		}

		ctx.Next()
	}
}
