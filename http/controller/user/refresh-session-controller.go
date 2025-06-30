package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/viniciusfal/erp/middleware"
)

type RefreshRequest struct {
	RefreshToken string `json:"refreshToken"`
}

type RefreshTokenController struct {
	jwtSecret string
}

func NewRefreshTokenController(jwtSecret string) *RefreshTokenController {
	return &RefreshTokenController{
		jwtSecret: jwtSecret,
	}
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refreshToken" binding:"required"`
}

type RefreshTokenResponse struct {
	AccessToken string `json:"accessToken"`
	ExpiresIn   int64  `json:"expiresIn"`
	TokenType   string `json:"tokenType"`
}

func (rtc *RefreshTokenController) HandleRefreshToken(c *gin.Context) {
	var request RefreshRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	// 1. Validar a requisição
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Requisição inválida",
			"details": err.Error(),
		})
		return
	}

	// 2. Validar o refresh token
	claims, err := middleware.ValidarToken(request.RefreshToken, rtc.jwtSecret)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Token de refresh inválido",
			"details": err.Error(),
		})
		return
	}

	// 3. Verificar se é realmente um refresh token
	if tokenType, ok := claims["type"].(string); !ok || tokenType != "refresh" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Tipo de token inválido",
			"details": "Esperado: refresh token",
		})
		return
	}

	// 4. Extrair claims necessários
	userID, ok1 := claims["sub"].(string)
	role, ok2 := claims["role"].(string)

	if !ok1 || !ok2 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Claims inválidas no token",
		})
		return
	}

	// 5. Gerar novo access token
	newAccessToken, err := middleware.GerarAccessToken(userID, role, rtc.jwtSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Falha ao gerar novo token",
			"details": err.Error(),
		})
		return
	}

	// 6. Retornar resposta padronizada
	response := RefreshTokenResponse{
		AccessToken: newAccessToken,
		ExpiresIn:   int64(middleware.AccessTokenExpiry.Seconds()),
		TokenType:   "Bearer",
	}

	c.JSON(http.StatusOK, response)
}
