package database

import (
	"log"
	"os"
	"path/filepath"
)

func RunMigrations() {
	if DB == nil {
		log.Fatalf("Database connection is nil")
	}
	migrationFiles := []string{
		"database/migrations/0001_create_managers_table.sql",
		"database/migrations/0002_create_museums_table.sql",
		"database/migrations/0003_create_artworks_table.sql",
	}

	for _, file := range migrationFiles {
		if err := applyMigration(file); err != nil {
			log.Fatalf("Failed to execute migration file %s: %v", file, err)
		}
	}
}

func applyMigration(file string) error {
	// Obter o caminho absoluto do arquivo de migração
	basePath, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get working directory: %v", err)
	}

	fullPath := filepath.Join(basePath, file)
	log.Printf("Applying migration from file: %s", fullPath)

	// Ler o conteúdo do arquivo
	migration, err := os.ReadFile(fullPath)
	if err != nil {
		return err
	}

	// Executar a migração no banco de dados
	return DB.Exec(string(migration)).Error
}