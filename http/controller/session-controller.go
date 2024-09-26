package controller

import (
	"log"
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
	var user model.User

	err := ctx.BindJSON(&session)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	if user.Password != services.SHA256Encoder(session.Password) {
		log.Fatal("invalid credentials")
		return
	}

	token, err := services.NewJWTService().GenerateToken(user.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(201, token)

}
