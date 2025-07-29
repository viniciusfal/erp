package controller

import (
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
)

type DownloadController struct{}

func NewDownloadController() DownloadController {
	return DownloadController{}
}

func (dc *DownloadController) DownloadFile(ctx *gin.Context) {
	const uploadDir = "uploads"

	fileName := ctx.Param("filename")

	// Validação profunda do nome do arquivo
	if fileName == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Nome de arquivo inválido"})
		return
	}
	// Não permitir path traversal, barras, barras invertidas, espaços ou caracteres especiais
	invalidName := strings.Contains(fileName, "..") || strings.Contains(fileName, "/") || strings.Contains(fileName, "\\") || strings.Contains(fileName, " ")
	// Permitir apenas letras, números, ponto, underline e hífen
	allowedPattern := regexp.MustCompile(`^[a-zA-Z0-9._-]+$`)
	if invalidName || !allowedPattern.MatchString(fileName) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Nome de arquivo inválido"})
		return
	}

	filePath := filepath.Join(uploadDir, fileName)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Arquivo não encontrado"})
		return
	}

	fileInfo, err := os.Stat(filePath)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao acessar arquivo"})
		return
	}

	if fileInfo.IsDir() {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Caminho inválido"})
		return
	}

	ctx.Header("Content-Description", "File Transfer")
	ctx.Header("Content-Disposition", "attachment; filename="+fileName)
	ctx.Header("Content-Type", "application/octet-stream")
	ctx.Header("Content-Transfer-Encoding", "binary")
	ctx.Header("Expires", "0")
	ctx.Header("Cache-Control", "must-revalidate")
	ctx.Header("Pragma", "public")

	ctx.File(filePath)
} 