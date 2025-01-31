package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/viniciusfal/erp/http/middleware"
	"github.com/viniciusfal/erp/infra/model"
	usecase "github.com/viniciusfal/erp/infra/usecase/user"
)

type CreateSessionController struct {
	sessionUseCase usecase.SessionUseCase
}

func NewCreateSessionController(usecase usecase.SessionUseCase) CreateSessionController {
	return CreateSessionController{
		sessionUseCase: usecase,
	}
}

func (uc *CreateSessionController) CreateSession(ctx *gin.Context) {
	var session model.Session

	err := ctx.BindJSON(&session)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	user, err := uc.sessionUseCase.CreateSession(session.Email, session.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Usuário ou senha inválidos"})
		return
	}

	token, err := middleware.GerarToken(user.ID)
	if err != nil {
		fmt.Println("Erro ao gerar token", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao gerar token"})
		return
	}

	fmt.Println("token JWT", token)

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
		"user": gin.H{
			"id":    user.ID,
			"email": user.Email,
		},
	})
}
