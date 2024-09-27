package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/viniciusfal/erp/factories"
	"github.com/viniciusfal/erp/http/middleware"
	"github.com/viniciusfal/erp/lib"
)

func TransactionRoutes(router *gin.Engine) {
	ListTransactionsController := factories.MakeListTransactions()
	TransactionController := factories.MakeTransactions()
	GetTransactionByIdController := factories.MakeGetTransactionById()
	GetTransactionByDateController := factories.MakeGetTransactionByDate()
	SetTransactionController := factories.MakeSetTransaction()
	RemoveTransactionController := factories.MakeRemoveTransaction()

	transactionGroup := router.Group("/transaction", middleware.Auth())
	{
		transactionGroup.GET("/", ListTransactionsController.GetTransactions)
		transactionGroup.POST("/", TransactionController.CreateTransaction)
		transactionGroup.GET("/:transactionId", GetTransactionByIdController.GetTransactionById)
		transactionGroup.GET("/by-date", GetTransactionByDateController.GetTransactionByDate)
		transactionGroup.PUT("/:transactionId", SetTransactionController.SetTransaction)
		transactionGroup.DELETE("/:transactionId", RemoveTransactionController.RemoveTransaction)
		transactionGroup.POST("/upload", lib.Upload)
	}
}
