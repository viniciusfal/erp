package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/viniciusfal/erp/factories"
)

func MetaRoutes(router *gin.Engine) {
	CreateMetaController := factories.MakeMeta()

	router.POST("/meta", CreateMetaController.CreateMeta)
}
