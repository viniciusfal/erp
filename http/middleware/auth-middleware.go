package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/viniciusfal/erp/services"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		const Bearer_schema = "Bearer "
		header := ctx.GetHeader("Authorization")
		if header == "" {
			ctx.AbortWithStatus(401)
		}

		if !strings.HasPrefix(header, Bearer_schema) {
			ctx.AbortWithStatusJSON(401, gin.H{"error": "invalid authorization schema"})
			return
		}

		token := header[len(Bearer_schema):]

		if !services.NewJWTService().ValidateToken(token) {
			ctx.AbortWithStatus(401)
		}

		ctx.Next()
	}

}
