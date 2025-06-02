package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/signintech/gopdf"
)

func PdfGenerate(ctx *gin.Context) {
	// Criar o PDF
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	pdf.AddPage()

	// Adicionar a fonte
	err := pdf.AddTTFFont("Arial", "./arial.ttf")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "Erro ao carregar a fonte do PDF.")
		return
	}

	// Adicionar Linhas com os dados
	err = pdf.SetFont("Arial", "", 16)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "Erro ao definir a fonte do PDF.")
		return
	}

}
