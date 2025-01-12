package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/viniciusfal/erp/factories"
)

func SafeRoutes(router *gin.RouterGroup) {
	CreateSafeController := factories.MakeSafe()
	ListSafesController := factories.MakeListSafe()
	GetSafesController := factories.MakeGetSafesByDate()
	SetSafesController := factories.MakeSetSafe()
	SetActiveSafeController := factories.MakeSetActive()

	router.POST("/safe", CreateSafeController.CreateSafe)
	router.GET("/safe", ListSafesController.GetSafes)
	router.GET("/safe/:startDate/:endDate", GetSafesController.GetSafesByDate)
	router.PUT("/safe/:id", SetSafesController.SetSafe)
	router.PATCH("/safe/:id", SetActiveSafeController.SetActiveSafe)
}
