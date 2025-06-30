package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/viniciusfal/erp/infra/model"
	usecase "github.com/viniciusfal/erp/infra/usecase/user"
	"github.com/viniciusfal/erp/middleware"
)

type CreateSessionController struct {
	sessionUseCase usecase.SessionUseCase
	jwtSecret      string
}

func NewCreateSessionController(usecase usecase.SessionUseCase, jwtSecret string) CreateSessionController {
	return CreateSessionController{
		sessionUseCase: usecase,
		jwtSecret:      jwtSecret,
	}
}

type SessionResponse struct {
	User         *model.User `json:"user"`
	AccessToken  string      `json:"accessToken"`
	RefreshToken string      `json:"refreshToken,omitempty"`
	ExpiresIn    int64       `json:"expiresIn"`
}

func (uc *CreateSessionController) CreateSession(ctx *gin.Context) {
	var session model.Session

	if err := ctx.ShouldBindJSON(&session); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	user, err := uc.sessionUseCase.CreateSession(session.Email, session.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Usuário ou senha inválidos"})
		return
	}

	// Gerar access token (1 hora de expiração)
	accessToken, err := middleware.GerarAccessToken(user.ID, user.Role, uc.jwtSecret)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao gerar access token"})
		return
	}

	// Gerar refresh token (7 dias de expiração)
	refreshToken, err := middleware.GerarRefreshToken(user.ID, user.Role, uc.jwtSecret)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao gerar refresh token"})
		return
	}

	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie(
		"access_token",
		accessToken,
		int(time.Hour.Seconds()),
		"/",
		".erpnet.tech", // ou "localhost" em dev
		true,           // secure
		true,           // httpOnly
	)
	ctx.SetCookie(
		"refresh_token",
		refreshToken,
		int((time.Hour * 24 * 7).Seconds()),
		"/",
		".erpnet.tech",
		true,
		true,
	)

	ctx.JSON(http.StatusOK, gin.H{
		"user": gin.H{
			"id":    user.ID,
			"email": user.Email,
			"name":  user.Name,
			"role":  user.Role,
		},
		"accessToken":  accessToken,
		"refreshToken": refreshToken, // Adicionado
		"expiresIn":    int64(middleware.AccessTokenExpiry.Seconds()),
	})
}
