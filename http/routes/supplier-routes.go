package routes

import (
	"github.com/gin-gonic/gin"
	factories "github.com/viniciusfal/erp/factories/supplier"
	"github.com/viniciusfal/erp/middleware"
)

func SupplierRoutes(router *gin.RouterGroup) {
	CreateSupplierController := factories.MakeSupplier()
	ListSupplierController := factories.MakeListSupplier()
	ListSupplierWithTransactions := factories.MakeListSupplierWithTransactions()
	SetSupplierController := factories.MakeSetSupplier()
	FindOneSupplierController := factories.MakeFindOneSupplier()
	DesactiveSupplierController := factories.MakeDesactiveSupplier()

	router.POST("/supplier",middleware.RBAC("supplier.create"), CreateSupplierController.CreateSupplier)
	router.GET("/supplier",middleware.RBAC("supplier.view"), ListSupplierController.ListSupplier)
	router.GET("/supplier/:id",middleware.RBAC("supplier.view"), FindOneSupplierController.FindOne)
	router.PUT("/supplier/:id",middleware.RBAC("supplier.edit"), SetSupplierController.SetSupplier)
	router.PATCH("/supplier/:id", middleware.RBAC("supplier.remove"), DesactiveSupplierController.Desactive)
	router.GET("/supplier/with-transactions", middleware.RBAC("supplier.view"), ListSupplierWithTransactions.ListSupplier)
}
