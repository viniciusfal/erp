package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/viniciusfal/erp/http/routes"
)

func main() {
	server := gin.Default()

	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},                 // Permita apenas esta origem
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE"}, // Corrigido
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	routes.TransactionRoutes(server)
	routes.UserRoutes(server)

	server.Run(":8000")
}
