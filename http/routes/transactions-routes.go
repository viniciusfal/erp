package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/viniciusfal/erp/factories"
	"github.com/viniciusfal/erp/middleware"
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
		transactionGroup.POST("/", middleware.RBAC("transaction.create"), TransactionController.CreateTransaction)
		transactionGroup.POST("/installment",  middleware.RBAC("transaction.create"), MakeCreateInstallmentTransactionController.CreateInstallmentTransactions)
		transactionGroup.POST("/low", middleware.RBAC("transaction.create"), MakeCreateLowTo0Controller.CreateLowTo0)
		transactionGroup.GET("/", middleware.RBAC("transaction.view"), ListTransactionsController.GetTransactions)
		transactionGroup.GET("/:transactionId", middleware.RBAC("transaction.view"), GetTransactionByIdController.GetTransactionById)
		transactionGroup.GET("/byDate/:status/:startDate/:endDate", middleware.RBAC("transaction.view"), GetTransactionByDateController.GetTransactionByDate)
		transactionGroup.GET("/byDate/low/today", middleware.RBAC("transaction.view"), MakeLowTodayTransaction.GetLowBydateToday)
		transactionGroup.GET("/byDate/low/current", middleware.RBAC("transaction.view"), MakeGetLowCurrentMonth.GetTransactionCurrent)
		transactionGroup.GET("/byDate/low/lastSeven", middleware.RBAC("transaction.view"), MakeLowSevenDays.GetTransactionByDate)
		transactionGroup.GET("/byDate/low/lastThird", middleware.RBAC("transaction.view"), MakeLowThirdyDays.GetTransactionByDate)
		transactionGroup.GET("/byDate/due/today", middleware.RBAC("transaction.view"), MakeDueGetToday.GetTransactionToday)
		transactionGroup.GET("/byDate/due/current", middleware.RBAC("transaction.view"), MakeDueGetCurrentMonth.GetDueCurrent)
		transactionGroup.GET("/byDate/due/lastSeven", middleware.RBAC("transaction.view"), MakeDueGetLasrSevenDays.GetTransactionByDate)
		transactionGroup.GET("/byDate/due/lastThird", middleware.RBAC("transaction.view"), MakeDueGetLastThirdDays.GetDueTransactions)
		transactionGroup.GET("/analitics", middleware.RBAC("transaction.view"), MakeAnalysesTransactionController.GetTransactionByDate)
		transactionGroup.PUT("/:transactionId", middleware.RBAC("transaction.edit"), SetTransactionController.SetTransaction)
		transactionGroup.PUT("/markLow/:transactionId", middleware.RBAC("transaction.edit"), MakeMarkLowTransactionController.MarkLow)
		transactionGroup.PATCH("/:transactionId", middleware.RBAC("transaction.edit"), MarkPaymentController.MarkPayment)
		transactionGroup.POST("/upload", middleware.RBAC("transaction.create"), MakeImportTransactions.ImportTransactions)
		transactionGroup.DELETE("/:transactionId", middleware.RBAC("transacton.remove"), RemoveTransactionController.RemoveTransaction)
	}
}
