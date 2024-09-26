package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/viniciusfal/erp/factories"
)

func UserRoutes(router *gin.Engine) {
	CreateUserController := factories.MakeUser()
	ListUsers := factories.MakeListUsers()

	router.POST("/user", CreateUserController.CreateUser)
	router.GET("/users", ListUsers.GetUsers)
}
