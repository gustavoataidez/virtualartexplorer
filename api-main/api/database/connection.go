package database

import (
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitDB initializes the database connection using the DB_URL environment variable
func InitDB() {
	// Obter a URL do banco de dados da variável de ambiente
	dbURL := viper.GetString("DB_URL")
	if dbURL == "" {
		log.Fatalf("DB_URL is not set. Please configure the database connection string.")
	}

	// Conectar ao banco de dados usando a string de conexão
	var err error
	DB, err = gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	log.Println("Successfully connected to the database.")
}

// CreateDatabase ensures that the database exists (no-op when using a single DB_URL)
func CreateDatabase() {
	// Verificar se a conexão com o banco está inicializada
	if DB == nil {
		log.Fatalf("Database connection is nil")
	}

	log.Println("Database connection is active. Skipping explicit database creation since DB_URL is used.")
}
