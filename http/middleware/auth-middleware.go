package middleware

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("Pr073!n@22") // Troque por um valor seguro

// GerarToken cria um token JWT para um usuário
func GerarToken(userID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(time.Hour * 24).Unix(), // Expira em 24 horas
		"iat": time.Now().Unix(),
		"iss": "erp-api",
	})

	return token.SignedString(secretKey)
}

// JWTMiddleware valida o token JWT
func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var tokenString string

		// Verificar se o token está no cabeçalho Authorization
		authHeader := c.GetHeader("Authorization")
		if authHeader != "" {
			tokenString = strings.TrimPrefix(authHeader, "Bearer ")
		} else {
			// Caso não esteja no cabeçalho, verificar se está no cookie
			var err error
			tokenString, err = c.Cookie("refresh_token") // Ou outro nome para o cookie
			if err != nil || tokenString == "" {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token ausente"})
				return
			}
		}

		// Parse do token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return secretKey, nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			return
		}

		// Pegando o userID do token e armazenando no contexto
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Erro ao ler token"})
			return
		}

		c.Set("userID", claims["sub"])
		c.Next()
	}
}
