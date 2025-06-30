package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

const (
	AccessTokenExpiry  = time.Hour * 1
	RefreshTokenExpiry = time.Hour * 24 * 7
)

func GerarAccessToken(userID, role, secretKey string) (string, error) {
	return GerarToken(userID, role, secretKey, "access", AccessTokenExpiry)
}

func GerarRefreshToken(userID, role, secretKey string) (string, error) {
	return GerarToken(userID, role, secretKey, "refresh", RefreshTokenExpiry)
}

// GerarToken cria um token JWT para um usuário
func GerarToken(userID, role, secretKey, tokenType string, expiry time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  userID,
		"exp":  time.Now().Add(expiry).Unix(),
		"role": role,
		"type": tokenType,
	})
	return token.SignedString([]byte(secretKey))
}

func JWTMiddleware(secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := extrairToken(c)
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Token de acesso ausente",
				"details": "Use o header Authorization: Bearer <token>",
			})
			return
		}

		claims, err := ValidarToken(tokenString, secretKey)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Token inválido",
				"details": err.Error(),
			})
			return
		}

		// Verificar se é um token de acesso
		if claims["type"] != "access" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Tipo de token inválido",
				"details": "Esperado: access token",
			})
			return
		}

		c.Set("claims", claims)
		c.Set("userID", claims["sub"])
		c.Set("userRole", claims["role"])
		c.Next()
	}
}

func extrairToken(c *gin.Context) string {
	// 1. Tenta do Header Authorization
	if authHeader := c.GetHeader("Authorization"); authHeader != "" {
		return strings.TrimPrefix(authHeader, "Bearer ")
	}

	// 2. Tenta do Cookie
	if token, err := c.Cookie("access_token"); err == nil {
		return token
	}

	return ""
}
func ValidarToken(tokenString, secretKey string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("método de assinatura inesperado: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("token inválido")
}
