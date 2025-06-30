package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/viniciusfal/erp/factories"
	controller "github.com/viniciusfal/erp/http/controller/user"
)

func UserRoutes(router *gin.RouterGroup, jwtSecret string) {
	CreateUserController := factories.MakeUser()
	ListUsers := factories.MakeListUsers()
	CreateSessionController := factories.MakeSession(jwtSecret)
	refreshController := controller.NewRefreshTokenController(jwtSecret)

	router.POST("/auth/refresh", refreshController.HandleRefreshToken)
	router.POST("/user", CreateUserController.CreateUser)
	router.GET("/user", ListUsers.GetUsers)
	router.POST("/session", CreateSessionController.CreateSession)
}
