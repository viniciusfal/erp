package main

import (
	"github.com/gin-gonic/gin"
	"github.com/viniciusfal/erp/db"
	"github.com/viniciusfal/erp/http/controller"
	"github.com/viniciusfal/erp/infra/repository"
	"github.com/viniciusfal/erp/infra/usecase"
)

func main() {
	router := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	TransactionRepository := repository.NewTransactionRepository(dbConnection)
	TransactionUseCase := usecase.NewTransactionUseCase(TransactionRepository)
	TransactionController := controller.NewTransactionController(TransactionUseCase)

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.POST("/transaction", TransactionController.CreateTransaction)

	router.Run(":8000")
}
