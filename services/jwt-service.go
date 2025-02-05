package services

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type JWTService struct {
	secretKey string
}

func NewJWTService() *JWTService {
	return &JWTService{
		secretKey: "Pr073!n@22", // Substitua por uma chave segura
	}
}

func (service *JWTService) GenerateToken(userID string) (string, string, error) {
	// Define a expiração do token
	accessTokenExpires := time.Now().Add(time.Hour * 1).Unix()
	refreshTokenExpires := time.Now().Add(time.Hour * 24 * 7).Unix()

	// Criando o token de acesso
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     accessTokenExpires,
	})

	accessTokenString, err := accessToken.SignedString([]byte(service.secretKey))
	if err != nil {
		return "", "", err
	}

	// Criando o token de refresh
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     refreshTokenExpires,
	})

	refreshTokenString, err := refreshToken.SignedString([]byte(service.secretKey))
	if err != nil {
		return "", "", err
	}

	return accessTokenString, refreshTokenString, nil
}

func (service *JWTService) SetTokenInCookie(c *gin.Context, refreshToken string) {
	c.SetSameSite(http.SameSiteNoneMode)
	c.SetCookie("refresh_token", refreshToken, 3600*24*7, "/", "erpnet.up.railway.app", true, true) // Cookies seguros (secure: false para localhost)

	c.JSON(http.StatusOK, gin.H{
		"message": "Login bem-sucedido",
	})
}
