package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/viniciusfal/erp/factories"
)

func TransactionRoutes(router *gin.RouterGroup) {
	ListTransactionsController := factories.MakeListTransactions()
	TransactionController := factories.MakeTransactions()
	GetTransactionByIdController := factories.MakeGetTransactionById()
	GetTransactionByDateController := factories.MakeGetTransactionByDate()
	SetTransactionController := factories.MakeSetTransaction()
	RemoveTransactionController := factories.MakeRemoveTransaction()
	MarkPaymentController := factories.MakeMarkPayment()
	MakeAnalysesTransactionController := factories.MakeAnalysesTransactionsByMonth()
	MakeImportTransactions := factories.MakeImportTransactions()
	MakeCreateInstallmentTransactionController := factories.MakeCreateInstallmentTransactionController()
	MakeCreateLowTo0Controller := factories.MakeCreateLowTo0Controller()
	MakeLowTodayTransaction := factories.MakeLowTodayTransaction()
	MakeMarkLowTransactionController := factories.MakeMarkLowTransactionController()
	MakeGetLowCurrentMonth := factories.MakeGetLowCurrentMonth()
	MakeLowSevenDays := factories.MakeLowGetSevenDays()
	MakeLowThirdyDays := factories.MakeLowGetThirdDays()
	MakeDueGetToday := factories.MakeDueTodayTransaction()
	MakeDueGetCurrentMonth := factories.MakeDueCurrentMonth()
	MakeDueGetLasrSevenDays := factories.MakeDueGetLastSevenDays()
	MakeDueGetLastThirdDays := factories.MakeDueLastThirdDays()
	transactionGroup := router.Group("/transaction")

	{
		transactionGroup.POST("/", TransactionController.CreateTransaction)
		transactionGroup.POST("/installment", MakeCreateInstallmentTransactionController.CreateInstallmentTransactions)
		transactionGroup.POST("/low", MakeCreateLowTo0Controller.CreateLowTo0)
		transactionGroup.GET("/", ListTransactionsController.GetTransactions)
		transactionGroup.GET("/:transactionId", GetTransactionByIdController.GetTransactionById)
		transactionGroup.GET("/byDate/:status/:startDate/:endDate", GetTransactionByDateController.GetTransactionByDate)
		transactionGroup.GET("/byDate/low/today", MakeLowTodayTransaction.GetLowBydateToday)
		transactionGroup.GET("/byDate/low/current", MakeGetLowCurrentMonth.GetTransactionCurrent)
		transactionGroup.GET("/byDate/low/lastSeven", MakeLowSevenDays.GetTransactionByDate)
		transactionGroup.GET("/byDate/low/lastThird", MakeLowThirdyDays.GetTransactionByDate)
		transactionGroup.GET("/byDate/due/today", MakeDueGetToday.GetTransactionToday)
		transactionGroup.GET("/byDate/due/current", MakeDueGetCurrentMonth.GetDueCurrent)
		transactionGroup.GET("/byDate/due/lastSeven", MakeDueGetLasrSevenDays.GetTransactionByDate)
		transactionGroup.GET("/byDate/due/lastThird", MakeDueGetLastThirdDays.GetDueTransactions)
		transactionGroup.GET("/analitics", MakeAnalysesTransactionController.GetTransactionByDate)
		transactionGroup.PUT("/:transactionId", SetTransactionController.SetTransaction)
		transactionGroup.PUT("/markLow/:transactionId", MakeMarkLowTransactionController.MarkLow)
		transactionGroup.PATCH("/:transactionId", MarkPaymentController.MarkPayment)
		transactionGroup.POST("/upload", MakeImportTransactions.ImportTransactions)
		transactionGroup.DELETE("/:transactionId", RemoveTransactionController.RemoveTransaction)
	}
}
