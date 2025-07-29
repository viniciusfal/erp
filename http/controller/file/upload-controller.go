package controller

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UploadController struct{}

func NewUploadController() UploadController {
	return UploadController{}
}

type UploadResponse struct {
	FilePath string `json:"file_path"`
	FileName string `json:"file_name"`
	FileSize int64  `json:"file_size"`
}

func (uc *UploadController) UploadFile(ctx *gin.Context) {
	// Configurações de upload
	const (
		maxFileSize = 10 * 1024 * 1024 // 10MB
		uploadDir   = "uploads"
	)

	// Criar diretório de upload se não existir
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar diretório de upload"})
		return
	}

	// Obter arquivo do request
	file, header, err := ctx.Request.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Arquivo não encontrado"})
		return
	}
	defer file.Close()

	// Validar tamanho do arquivo
	if header.Size > maxFileSize {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Arquivo muito grande. Tamanho máximo: 10MB"})
		return
	}

	// Validar tipo de arquivo
	allowedTypes := []string{".pdf", ".jpg", ".jpeg", ".png", ".doc", ".docx", ".xls", ".xlsx"}
	fileExt := strings.ToLower(filepath.Ext(header.Filename))
	
	isAllowed := false
	for _, allowedType := range allowedTypes {
		if fileExt == allowedType {
			isAllowed = true
			break
		}
	}
	
	if !isAllowed {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Tipo de arquivo não permitido"})
		return
	}

	// Gerar nome único para o arquivo
	fileName := fmt.Sprintf("%s_%s%s", 
		time.Now().Format("20060102_150405"),
		uuid.New().String()[:8],
		fileExt)
	
	filePath := filepath.Join(uploadDir, fileName)

	// Salvar arquivo
	dst, err := os.Create(filePath)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao salvar arquivo"})
		return
	}
	defer dst.Close()

	// Copiar conteúdo do arquivo
	if _, err := ctx.Request.MultipartForm.File["file"][0].Open(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao processar arquivo"})
		return
	}

	response := UploadResponse{
		FilePath: filePath,
		FileName: fileName,
		FileSize: header.Size,
	}

	ctx.JSON(http.StatusOK, response)
} 