// middleware/rbac.go
package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/viniciusfal/erp/rbac"
)

func RBAC(requiredPermission string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. Obtém as claims do JWT (já validado pelo JWTMiddleware)
		claims, exists := c.Get("claims")
		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token inválido ou ausente"})
			return
		}

		// 2. Converte as claims para jwt.MapClaims
		jwtClaims, ok := claims.(jwt.MapClaims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Erro ao processar permissões"})
			return
		}

		// 3. Obtém o role do usuário do token
		userRole, roleExists := jwtClaims["role"].(string)
		if !roleExists {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Papel do usuário não encontrado"})
			return
		}

		// 4. Verifica se o role tem a permissão necessária
		permissions, roleExists := rbac.RolePermissions[userRole]
		if !roleExists {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Papel não encontrado"})
			return
		}

		hasPermission := false
		for _, perm := range permissions {
			if perm == requiredPermission {
				hasPermission = true
				break
			}
		}

		if !hasPermission {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error":               "Acesso negado",
				"detail":              "Você não tem permissão para acessar este recurso",
				"required_permission": requiredPermission,
				"user_rope":           userRole,
			})
			return
		}

		c.Next()
	}
}
