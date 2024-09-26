package main

import (
	"github.com/gin-gonic/gin"
	"github.com/viniciusfal/erp/http/routes"
)

func main() {
	server := gin.Default()

	routes.TransactionRoutes(server)
	routes.UserRoutes(server)

	server.Run(":8000")
}
