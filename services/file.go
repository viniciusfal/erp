package services

import (
	"io"
	"mime/multipart"
	"os"
)

// SaveTempFile salva um arquivo temporário e retorna o caminho
func SaveTempFile(file multipart.File, filePath string) error {
	outFile, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, file)
	return err
}

// DeleteTempFile remove o arquivo temporário após o processamento
func DeleteTempFile(filePath string) {
	_ = os.Remove(filePath) // Ignora erro caso o arquivo já tenha sido deletado
}
