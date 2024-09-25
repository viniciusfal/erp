package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/viniciusfal/erp/factories"
	"github.com/viniciusfal/erp/lib"
)

func TransactionRoutes(router *gin.Engine) {
	ListTransactionsController := factories.MakeListTransactions()
	TransactionController := factories.MakeTransactions()
	GetTransactionByIdController := factories.MakeGetTransactionById()
	GetTransactionByDateController := factories.MakeGetTransactionByDate()
	SetTransactionController := factories.MakeSetTransaction()
	RemoveTransactionController := factories.MakeRemoveTransaction()

	router.GET("/transactions", ListTransactionsController.GetTransactions)
	router.POST("/transaction", TransactionController.CreateTransaction)
	router.GET("/transaction/:transactionId", GetTransactionByIdController.GetTransactionById)
	router.GET("/transactions/by-date", GetTransactionByDateController.GetTransactionByDate)
	router.PUT("/transaction/:transactionId", SetTransactionController.SetTransaction)
	router.DELETE("/transaction/:transactionId", RemoveTransactionController.RemoveTransaction)
	router.POST("/transaction/upload", lib.Upload)
}
