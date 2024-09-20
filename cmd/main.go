package main

import (
	"github.com/gin-gonic/gin"
	"github.com/viniciusfal/erp/http/routes"
)

func main() {
	server := gin.Default()

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	routes.TransactionRoutes()

	server.Run(":8000")
}
