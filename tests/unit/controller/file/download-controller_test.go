package file

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path/filepath"
	"sync"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	file "github.com/viniciusfal/erp/http/controller/file"
)

func TestDownloadController_DownloadFile(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Criar arquivo de teste
	testFileName := "test_download.pdf"
	testContent := "conteudo do arquivo de teste"

	uploadDir := "uploads"
	os.MkdirAll(uploadDir, 0755)
	testFilePath := filepath.Join(uploadDir, testFileName)
	err := os.WriteFile(testFilePath, []byte(testContent), 0644)
	assert.NoError(t, err)

	tests := []struct {
		name           string
		filename       string
		expectedStatus int
		expectedError  bool
	}{
		{
			name:           "Download arquivo válido",
			filename:       testFileName,
			expectedStatus: http.StatusOK,
			expectedError:  false,
		},
		{
			name:           "Download arquivo inexistente",
			filename:       "arquivo_inexistente.pdf",
			expectedStatus: http.StatusNotFound,
			expectedError:  true,
		},
		{
			name:           "Download com nome vazio",
			filename:       "",
			expectedStatus: http.StatusBadRequest,
			expectedError:  true,
		},
		{
			name:           "Download com path traversal",
			filename:       "../../../etc/passwd",
			expectedStatus: http.StatusBadRequest,
			expectedError:  true,
		},
		{
			name:           "Download com caracteres especiais",
			filename:       "arquivo com espaços.pdf",
			expectedStatus: http.StatusBadRequest,
			expectedError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			controller := file.NewDownloadController()

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			urlPath := "/download/" + url.PathEscape(tt.filename)
			c.Request = httptest.NewRequest("GET", urlPath, nil)
			c.Params = gin.Params{{Key: "filename", Value: tt.filename}}

			controller.DownloadFile(c)

			assert.Equal(t, tt.expectedStatus, w.Code)

			if !tt.expectedError {
				assert.Contains(t, w.Header().Get("Content-Disposition"), "attachment")
				assert.Equal(t, "application/octet-stream", w.Header().Get("Content-Type"))
				assert.Equal(t, testContent, w.Body.String())
			}
		})
	}

	os.Remove(testFilePath)
}

func TestDownloadController_ConcurrentDownloads(t *testing.T) {
	gin.SetMode(gin.TestMode)
	controller := file.NewDownloadController()

	const numFiles = 5
	testFiles := make([]string, numFiles)

	uploadDir := "uploads"
	os.MkdirAll(uploadDir, 0755)

	for i := 0; i < numFiles; i++ {
		filename := fmt.Sprintf("test_%d.pdf", i)
		content := fmt.Sprintf("conteudo do arquivo %d", i)
		filepath := filepath.Join(uploadDir, filename)
		err := os.WriteFile(filepath, []byte(content), 0644)
		assert.NoError(t, err)
		testFiles[i] = filename
	}

	var wg sync.WaitGroup
	results := make(chan int, numFiles)

	downloadFile := func(filename string) {
		defer wg.Done()
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		urlPath := "/download/" + url.PathEscape(filename)
		c.Request = httptest.NewRequest("GET", urlPath, nil)
		c.Params = gin.Params{{Key: "filename", Value: filename}}
		controller.DownloadFile(c)
		results <- w.Code
	}

	for _, filename := range testFiles {
		wg.Add(1)
		go downloadFile(filename)
	}

	wg.Wait()
	close(results)

	successCount := 0
	for status := range results {
		if status == http.StatusOK {
			successCount++
		}
	}

	assert.Equal(t, numFiles, successCount, "Todos os downloads concorrentes devem ter sucesso")

	for _, filename := range testFiles {
		os.Remove(filepath.Join(uploadDir, filename))
	}
}

func TestDownloadController_SecurityTests(t *testing.T) {
	gin.SetMode(gin.TestMode)
	controller := file.NewDownloadController()

	maliciousFilenames := []string{
		"../../../etc/passwd",
		"..\\..\\..\\windows\\system32\\config\\sam",
		"file/with/slashes.pdf",
		"file\\with\\backslashes.pdf",
		"file%2e%2e%2f%2e%2e%2f%2e%2e%2fetc%2fpasswd", // URL encoded
	}

	for _, filename := range maliciousFilenames {
		t.Run("Security_"+filename, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			urlPath := "/download/" + url.PathEscape(filename)
			c.Request = httptest.NewRequest("GET", urlPath, nil)
			c.Params = gin.Params{{Key: "filename", Value: filename}}
			controller.DownloadFile(c)
			assert.Equal(t, http.StatusBadRequest, w.Code)
			assert.Contains(t, w.Body.String(), "Nome de arquivo inválido")
		})
	}
}

func TestDownloadController_Cleanup(t *testing.T) {
	uploadDir := "uploads"
	if _, err := os.Stat(uploadDir); err == nil {
		os.RemoveAll(uploadDir)
	}
} 