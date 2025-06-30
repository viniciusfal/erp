package routes

import (
	"github.com/gin-gonic/gin"
	factories "github.com/viniciusfal/erp/factories/cashier"
)

func CashierRoutes(router *gin.RouterGroup) {
	CreateCashierController := factories.MakeCreateCashier()
	getCashiersControllers := factories.MakeGetAllCashier()
	getCashierByIdController := factories.MakeGetcashierById()
	getCashierByUsernameController := factories.MakeGetCashierByUsername()
	setCashierController := factories.MakeSetCashier()
	alterCashierController := factories.MakeAlterCashierStatus()

	router.POST("/cashier", CreateCashierController.CreateCashier)
	router.GET("/cashier", getCashiersControllers.GetAllCashier)
	router.GET("/cashier/:id", getCashierByIdController.GetCashierByID)
	router.GET("/cashier/user/:username", getCashierByUsernameController.GetCashierByUsername)
	router.PATCH("/cashier/:id", setCashierController.SetCashier)
	router.PATCH("/cashier/user/:id", alterCashierController.AlterCashierStatus)
}
