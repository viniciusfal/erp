package services

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

type FileCleanupService struct {
	db          *sql.DB
	uploadDir   string
	stopChan    chan bool
	isRunning   bool
	mutex       sync.RWMutex
}

type FileInfo struct {
	Path         string
	Size         int64
	LastModified time.Time
}

func NewFileCleanupService(db *sql.DB, uploadDir string) *FileCleanupService {
	return &FileCleanupService{
		db:        db,
		uploadDir: uploadDir,
		stopChan:  make(chan bool),
	}
}

// Start inicia o serviço de limpeza em background
func (fcs *FileCleanupService) Start() {
	fcs.mutex.Lock()
	if fcs.isRunning {
		fcs.mutex.Unlock()
		return
	}
	fcs.isRunning = true
	fcs.mutex.Unlock()

	go fcs.runCleanupLoop()
}

// Stop para o serviço de limpeza
func (fcs *FileCleanupService) Stop() {
	fcs.mutex.Lock()
	defer fcs.mutex.Unlock()
	
	if fcs.isRunning {
		fcs.stopChan <- true
		fcs.isRunning = false
	}
}

// runCleanupLoop executa a limpeza periodicamente
func (fcs *FileCleanupService) runCleanupLoop() {
	ticker := time.NewTicker(24 * time.Hour) // Executa a cada 24 horas
	defer ticker.Stop()

	// Executar limpeza imediatamente na primeira vez
	go fcs.performCleanup()

	for {
		select {
		case <-ticker.C:
			go fcs.performCleanup()
		case <-fcs.stopChan:
			return
		}
	}
}

// performCleanup executa a limpeza de arquivos órfãos
func (fcs *FileCleanupService) performCleanup() {
	fmt.Println("Iniciando limpeza de arquivos órfãos...")

	// Obter todos os arquivos no diretório de upload
	files, err := fcs.scanUploadDirectory()
	if err != nil {
		fmt.Printf("Erro ao escanear diretório de upload: %v\n", err)
		return
	}

	// Obter referências válidas do banco de dados
	validReferences, err := fcs.getValidFileReferences()
	if err != nil {
		fmt.Printf("Erro ao obter referências do banco: %v\n", err)
		return
	}

	// Identificar arquivos órfãos
	orphanedFiles := fcs.findOrphanedFiles(files, validReferences)

	// Remover arquivos órfãos em paralelo
	if len(orphanedFiles) > 0 {
		fcs.removeOrphanedFiles(orphanedFiles)
	}

	fmt.Printf("Limpeza concluída. %d arquivos órfãos removidos.\n", len(orphanedFiles))
}

// scanUploadDirectory escaneia o diretório de upload
func (fcs *FileCleanupService) scanUploadDirectory() ([]FileInfo, error) {
	var files []FileInfo

	err := filepath.Walk(fcs.uploadDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Ignorar diretórios
		if info.IsDir() {
			return nil
		}

		// Ignorar arquivos ocultos
		if strings.HasPrefix(filepath.Base(path), ".") {
			return nil
		}

		// Calcular caminho relativo
		relPath, err := filepath.Rel(fcs.uploadDir, path)
		if err != nil {
			return err
		}

		files = append(files, FileInfo{
			Path:         relPath,
			Size:         info.Size(),
			LastModified: info.ModTime(),
		})

		return nil
	})

	return files, err
}

// getValidFileReferences obtém todas as referências válidas do banco
func (fcs *FileCleanupService) getValidFileReferences() (map[string]bool, error) {
	references := make(map[string]bool)

	// Query para accountability
	rows, err := fcs.db.Query("SELECT annex FROM accountability WHERE annex IS NOT NULL")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var annex []string
		if err := rows.Scan(&annex); err != nil {
			continue
		}
		for _, filePath := range annex {
			if filePath != "" {
				references[filePath] = true
			}
		}
	}

	// Query para accountability_change_request
	rows2, err := fcs.db.Query("SELECT new_annex FROM accountability_change_request WHERE new_annex IS NOT NULL")
	if err != nil {
		return nil, err
	}
	defer rows2.Close()

	for rows2.Next() {
		var annex []string
		if err := rows2.Scan(&annex); err != nil {
			continue
		}
		for _, filePath := range annex {
			if filePath != "" {
				references[filePath] = true
			}
		}
	}

	return references, nil
}

// findOrphanedFiles identifica arquivos órfãos
func (fcs *FileCleanupService) findOrphanedFiles(files []FileInfo, validReferences map[string]bool) []FileInfo {
	var orphaned []FileInfo

	for _, file := range files {
		// Verificar se o arquivo está referenciado no banco
		if !validReferences[file.Path] {
			// Verificar se o arquivo é muito antigo (mais de 30 dias)
			if time.Since(file.LastModified) > 30*24*time.Hour {
				orphaned = append(orphaned, file)
			}
		}
	}

	return orphaned
}

// removeOrphanedFiles remove arquivos órfãos em paralelo
func (fcs *FileCleanupService) removeOrphanedFiles(files []FileInfo) {
	var wg sync.WaitGroup
	results := make(chan string, len(files))

	// Função para remover arquivo individual
	removeFile := func(file FileInfo) {
		defer wg.Done()

		filePath := filepath.Join(fcs.uploadDir, file.Path)
		if err := os.Remove(filePath); err != nil {
			results <- fmt.Sprintf("Erro ao remover %s: %v", file.Path, err)
		} else {
			results <- fmt.Sprintf("Removido: %s", file.Path)
		}
	}

	// Executar remoções em paralelo
	for _, file := range files {
		wg.Add(1)
		go removeFile(file)
	}

	// Aguardar conclusão
	wg.Wait()
	close(results)

	// Log dos resultados
	for result := range results {
		fmt.Println(result)
	}
}

// ForceCleanup força uma limpeza imediata
func (fcs *FileCleanupService) ForceCleanup() {
	go fcs.performCleanup()
}

// IsRunning verifica se o serviço está rodando
func (fcs *FileCleanupService) IsRunning() bool {
	fcs.mutex.RLock()
	defer fcs.mutex.RUnlock()
	return fcs.isRunning
} 