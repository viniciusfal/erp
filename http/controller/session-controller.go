package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/viniciusfal/erp/infra/model"
	"github.com/viniciusfal/erp/infra/usecase"
	"github.com/viniciusfal/erp/services"
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
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	user, err := uc.sessionUseCase.CreateSession(session.Email, session.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, err.Error())
		return
	}

	token, err := services.NewJWTService().GenerateToken(user.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(201, token)

}
