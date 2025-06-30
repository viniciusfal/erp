package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/viniciusfal/erp/db"
	"github.com/viniciusfal/erp/http/routes"
	"github.com/viniciusfal/erp/middleware"
)

func main() {
	godotenv.Load() // Carrega .env
	secret := os.Getenv("JWT_SECRET")

	dbConnection := db.RunDB()
	server := gin.Default()

	// Configuração do CORS para permitir acesso tanto do frontend local quanto da produção
	server.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000", "https://erpnet.tech"}, // Domínios permitidos
		AllowMethods: []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
		AllowHeaders: []string{
			"Origin",
			"Content-Type",
			"Accept",
			"Authorization",
			"X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length", "Set-cookie"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	server.Use(gin.Logger())   // Log de requisições
	server.Use(gin.Recovery()) // Recupera de panics
	server.MaxMultipartMemory = 20 << 20

	// Servir a pasta 'uploads' como estática
	server.Static("/uploads", "./uploads")

	apiPublic := server.Group("/api")
	{
		routes.UserRoutes(apiPublic, secret)
	}

	api := server.Group("/api")
	api.Use(middleware.JWTMiddleware(secret))
	{
		routes.TransactionRoutes(api)
		routes.MetaRoutes(api)
		routes.SafeRoutes(api)
		routes.SupplierRoutes(api)
		routes.CashierRoutes(api)
		routes.AccountabilityRoutes(api)
	}

	fmt.Printf("Banco de dados conectado com sucesso: %v\n", dbConnection)

	port := os.Getenv("DBPORT")
	if port == "" {
		port = "8000"
	}

	server.Run("0.0.0.0:" + port) // Escuta em todos os IPs
}
