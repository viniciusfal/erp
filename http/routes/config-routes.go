package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/viniciusfal/erp/factories"
	"github.com/viniciusfal/erp/middleware"
)

func ConfigRoutes(router *gin.RouterGroup) {
	GetConfigController := factories.MakeGetConfig()
	UpdateConfigController := factories.MakeUpdateConfig()

	router.GET("/config", middleware.RBAC("config.view"), GetConfigController.GetConfig)
	router.PUT("/config", middleware.RBAC("config.update"), UpdateConfigController.UpdateConfig)
} 