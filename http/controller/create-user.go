package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/viniciusfal/erp/infra/model"
	"github.com/viniciusfal/erp/infra/usecase"
	"github.com/viniciusfal/erp/services"
)

type CreateUserController struct {
	createUserUseCase usecase.UserUseCase
}

func NewCreateUserController(usecase usecase.UserUseCase) CreateUserController {
	return CreateUserController{
		createUserUseCase: usecase,
	}
}

func (uc *CreateUserController) CreateUser(ctx *gin.Context) {
	var user model.User

	err := ctx.BindJSON(&user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	user.Password = services.SHA256Encoder(user.Password)

	isertedUser, err := uc.createUserUseCase.CreateUser(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, isertedUser)

}
