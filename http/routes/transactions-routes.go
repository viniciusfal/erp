package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/viniciusfal/erp/factories"
)

func TransactionRoutes(router *gin.Engine) {
	ListTransactionsController := factories.MakeListTransactions()
	TransactionController := factories.MakeTransactions()
	GetTransactionByIdController := factories.MakeGetTransactionById()
	SetTransactionController := factories.MakeSetTransaction()

	router.GET("/transactions", ListTransactionsController.GetTransactions)
	router.POST("/transaction", TransactionController.CreateTransaction)
	router.GET("/transaction/:transactionId", GetTransactionByIdController.GetTransactionById)
	router.PUT("/transaction/:transactionId", SetTransactionController.SetTransaction)
}
