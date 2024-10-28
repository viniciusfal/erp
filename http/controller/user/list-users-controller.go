package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	usecase "github.com/viniciusfal/erp/infra/usecase/user"
)

type ListUserController struct {
	listUserUseCase usecase.ListUserUseCase
}

func NewListUserController(usecase usecase.ListUserUseCase) ListUserController {
	return ListUserController{
		listUserUseCase: usecase,
	}
}

func (uc *ListUserController) GetUsers(ctx *gin.Context) {
	users, err := uc.listUserUseCase.GetUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, users)

}
