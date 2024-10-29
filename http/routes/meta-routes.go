package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/viniciusfal/erp/factories"
)

func MetaRoutes(router *gin.Engine) {
	CreateMetaController := factories.MakeMeta()
	GetMetasController := factories.MakeMetas()
	GetMetaByMonth := factories.MakeGetMetaByMonth()
	SetMetaController := factories.MAkeSetMeta()

	router.POST("/meta", CreateMetaController.CreateMeta)
	router.GET("/metas", GetMetasController.GetMetas)
	router.GET("meta/:month", GetMetaByMonth.GetMetaByMonth)
	router.PATCH("meta/:id", SetMetaController.SetMeta)
}
