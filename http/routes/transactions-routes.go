package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/viniciusfal/erp/factories"
)

func TransactionRoutes(router *gin.Engine) {
	ListTransactionsController := factories.MakeListTransactions()
	TransactionController := factories.MakeTransactions()

	router.GET("/transactions", ListTransactionsController.GetTransactions)
	router.POST("/transaction", TransactionController.CreateTransaction)
}
