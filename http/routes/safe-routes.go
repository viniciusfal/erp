package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/viniciusfal/erp/factories"
)

func SafeRoutes(router *gin.Engine) {
	CreateSafeController := factories.MakeSafe()
	ListSafesController := factories.MakeListSafe()
	GetSafesController := factories.MakeGetSafesByDate()
	SetSafesController := factories.MakeSetSafe()

	router.POST("/safe", CreateSafeController.CreateSafe)
	router.GET("/safe", ListSafesController.GetSafes)
	router.GET("/safe/:startDate/:endDate", GetSafesController.GetSafesByDate)
	router.PUT("/safe/:id", SetSafesController.SetSafeController)
}
