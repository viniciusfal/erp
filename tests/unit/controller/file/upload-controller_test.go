package file

import (
	"bytes"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"sync"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	file "github.com/viniciusfal/erp/http/controller/file"
)

func TestUploadController_UploadFile(t *testing.T) {
	// Configurar Gin para modo de teste
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name           string
		fileName       string
		fileContent    string
		expectedStatus int
		expectedError  bool
	}{
		{
			name:           "Upload arquivo PDF válido",
			fileName:       "test.pdf",
			fileContent:    "conteudo do pdf",
			expectedStatus: http.StatusOK,
			expectedError:  false,
		},
		{
			name:           "Upload arquivo JPG válido",
			fileName:       "image.jpg",
			fileContent:    "conteudo da imagem",
			expectedStatus: http.StatusOK,
			expectedError:  false,
		},
		{
			name:           "Upload arquivo tipo inválido",
			fileName:       "test.exe",
			fileContent:    "conteudo executavel",
			expectedStatus: http.StatusBadRequest,
			expectedError:  true,
		},
		{
			name:           "Upload sem arquivo",
			fileName:       "",
			fileContent:    "",
			expectedStatus: http.StatusBadRequest,
			expectedError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			controller := file.NewUploadController()

			// Criar request
			req := httptest.NewRequest("POST", "/upload", nil)
			w := httptest.NewRecorder()

			if tt.fileName != "" {
				body := &bytes.Buffer{}
				writer := multipart.NewWriter(body)
				part, err := writer.CreateFormFile("file", tt.fileName)
				assert.NoError(t, err)
				_, err = part.Write([]byte(tt.fileContent))
				assert.NoError(t, err)
				writer.Close()
				req = httptest.NewRequest("POST", "/upload", body)
				req.Header.Set("Content-Type", writer.FormDataContentType())
			}

			c, _ := gin.CreateTestContext(w)
			c.Request = req

			controller.UploadFile(c)

			assert.Equal(t, tt.expectedStatus, w.Code)

			if !tt.expectedError {
				responseBody := w.Body.String()
				assert.Contains(t, responseBody, "file_path")
				assert.Contains(t, responseBody, "file_name")
				assert.Contains(t, responseBody, "file_size")
			}
		})
	}
}

func TestUploadController_ConcurrentUploads(t *testing.T) {
	gin.SetMode(gin.TestMode)
	controller := file.NewUploadController()

	const numUploads = 10
	var wg sync.WaitGroup
	results := make(chan int, numUploads)

	uploadFile := func(id int) {
		defer wg.Done()

		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		part, err := writer.CreateFormFile("file", fmt.Sprintf("test_%d.pdf", id))
		if err != nil {
			results <- http.StatusInternalServerError
			return
		}
		part.Write([]byte(fmt.Sprintf("conteudo do arquivo %d", id)))
		writer.Close()

		req := httptest.NewRequest("POST", "/upload", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = req

		controller.UploadFile(c)
		results <- w.Code
	}

	for i := 0; i < numUploads; i++ {
		wg.Add(1)
		go uploadFile(i)
	}

	wg.Wait()
	close(results)

	successCount := 0
	for status := range results {
		if status == http.StatusOK {
			successCount++
		}
	}

	assert.Equal(t, numUploads, successCount, "Todos os uploads concorrentes devem ter sucesso")
}

func TestUploadController_FileSizeLimit(t *testing.T) {
	gin.SetMode(gin.TestMode)
	controller := file.NewUploadController()

	largeContent := make([]byte, 11*1024*1024) // 11MB

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", "large.pdf")
	assert.NoError(t, err)
	part.Write(largeContent)
	writer.Close()

	req := httptest.NewRequest("POST", "/upload", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	controller.UploadFile(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "Arquivo muito grande")
}

func TestUploadController_Cleanup(t *testing.T) {
	uploadDir := "uploads"
	if _, err := os.Stat(uploadDir); err == nil {
		os.RemoveAll(uploadDir)
	}
} 