package routes

import (
	"github.com/gin-gonic/gin"
	factories "github.com/viniciusfal/erp/factories/supplier"
)

func SupplierRoutes(router *gin.RouterGroup) {
	CreateSupplierController := factories.MakeSupplier()
	ListSupplierController := factories.MakeListSupplier()
	ListSupplierWithTransactions := factories.MakeListSupplierWithTransactions()
	SetSupplierController := factories.MakeSetSupplier()
	FindOneSupplierController := factories.MakeFindOneSupplier()
	DesactiveSupplierController := factories.MakeDesactiveSupplier()

	router.POST("/supplier", CreateSupplierController.CreateSupplier)
	router.GET("/supplier", ListSupplierController.ListSupplier)
	router.GET("/supplier/:id", FindOneSupplierController.FindOne)
	router.PUT("/supplier/:id", SetSupplierController.SetSupplier)
	router.PATCH("/supplier/:id", DesactiveSupplierController.Desactive)
	router.GET("/supplier/with-transactions", ListSupplierWithTransactions.ListSupplier)
}
