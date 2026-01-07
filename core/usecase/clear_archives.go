package usecase

import (
	"faceclean/core/domain/paths"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func isProtected(path string) bool {
	lower := strings.ToLower(path)
	for _, p := range paths.ProtectedPaths {
		if strings.Contains(lower, strings.ToLower(p)) {
			return true
		}
	}
	return false
}

func ClearArchives() {
	minSized := 200
	var totalCleanedMB int64
	var filesDeleted int
	root := "C:\\"
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil || info == nil {
			return nil
		}
		if info.IsDir() {
			if isProtected(path) {
				return filepath.SkipDir
			}
			return nil
		}

		sizeMB := info.Size() / (1024 * 1024)

		if sizeMB >= int64(minSized) && !isProtected(path) {
			fmt.Printf("Excluindo: %s (%d MB)\n", path, sizeMB)
			if err := os.Remove(path); err == nil {
				totalCleanedMB += sizeMB
				filesDeleted++
			} else {
				fmt.Println("Erro ao excluir:", err)
			}
		}
		return nil
	})
	if err != nil {
		fmt.Println("Erro:", err)
	}
	fmt.Println("\nLimpeza concluída!")
	fmt.Printf("Arquivos removidos: %d\n", filesDeleted)
	fmt.Printf("Espaço liberado: %d MB (%.2f GB)\n",
		totalCleanedMB,
		float64(totalCleanedMB)/1024,
	)
}
