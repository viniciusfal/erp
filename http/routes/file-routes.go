package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/viniciusfal/erp/factories"
	"github.com/viniciusfal/erp/middleware"
)

func FileRoutes(router *gin.RouterGroup) {
	UploadController := factories.MakeUploadController()
	DownloadController := factories.MakeDownloadController()

	// Upload de arquivos (apenas usuários autenticados com permissão)
	router.POST("/upload", middleware.RBAC("file.upload"), UploadController.UploadFile)
	
	// Download de arquivos (apenas usuários autenticados com permissão)
	router.GET("/download/:filename", middleware.RBAC("file.download"), DownloadController.DownloadFile)
} 