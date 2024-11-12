package main

import (
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/viniciusfal/erp/http/routes"
)

func main() {
	server := gin.Default()

	// Configuração do CORS para permitir acesso tanto do frontend local quanto da produção
	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "erp.railway.internal", "https://erpamazoniainter.vercel.app"}, // Domínios permitidos
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Configuração das rotas
	routes.TransactionRoutes(server)
	routes.UserRoutes(server)
	routes.MetaRoutes(server)
	routes.SafeRoutes(server)

	// Obter a porta via variável de ambiente para compatibilidade com o Render
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" // Porta padrão para desenvolvimento
	}

	// Rodar o servidor na porta especificada
	server.Run(":" + port)
}
