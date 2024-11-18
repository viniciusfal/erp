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
		AllowOrigins:     []string{"http://localhost:5173", "https://erp-1xqz.onrender.com", "https://cornflowerblue-pony-878669.hostingersite.com"}, // Domínios permitidos
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	routes.TransactionRoutes(server)
	routes.MetaRoutes(server)
	routes.SafeRoutes(server)
	routes.UserRoutes(server)

	fmt.Printf("Banco de dados conectado com sucesso: %v\n", dbConnection)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" // Porta padrão para desenvolvimento
	}

	server.Run(":" + port)
}
