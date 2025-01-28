package services

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

func Savefile(file multipart.File, header *multipart.FileHeader, dest string) (string, error) {
	if file == nil {
		return "", fmt.Errorf("arquivo vazio")
	}
	err := os.MkdirAll(dest, os.ModePerm)
	if err != nil {
		return "", fmt.Errorf("erro ao criar diretório: %v", err)
	}

	filepath := filepath.Join(dest, header.Filename)

	out, err := os.Create(filepath)
	if err != nil {
		return "", fmt.Errorf("erro ao criar arquivo: %v", err)
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		return "", fmt.Errorf("erro ao copiar arquivo: %v", err)
	}

	return filepath, nil
}
