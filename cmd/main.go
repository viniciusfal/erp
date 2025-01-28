package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/viniciusfal/erp/db"
	"github.com/viniciusfal/erp/http/routes"
)

func main() {
	dbConnection := db.RunDB()
	server := gin.Default()

	// Configuração do CORS para permitir acesso tanto do frontend local quanto da produção
	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "https://erpnet.tech", "https://www.erpnet.tech"}, // Domínios permitidos
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	server.MaxMultipartMemory = 12 << 20 // 20 MB

	// Servir a pasta 'uploads' como estática
	server.Static("/uploads", "./uploads")

	api := server.Group("/api")
	{
		routes.TransactionRoutes(api)
		routes.MetaRoutes(api)
		routes.SafeRoutes(api)
		routes.UserRoutes(api)
	}

	fmt.Printf("Banco de dados conectado com sucesso: %v\n", dbConnection)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" // Porta padrão para desenvolvimento
	}

	server.Run(":" + port)
}
