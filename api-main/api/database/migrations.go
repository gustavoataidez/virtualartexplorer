package database

import (
	"log"
)

func RunMigrations() {
	if DB == nil {
		log.Fatalf("Database connection is nil")
	}

	migrations := []string{
		`
		CREATE TABLE IF NOT EXISTS managers (
			id SERIAL PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			email VARCHAR(255) UNIQUE NOT NULL,
			password VARCHAR(255) NOT NULL
		);
		`,
		`
		CREATE TABLE IF NOT EXISTS museums (
			id SERIAL PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			location VARCHAR(255) NOT NULL,
			category VARCHAR(255)
		);
		`,
		`
		CREATE TABLE IF NOT EXISTS artworks (
			id SERIAL PRIMARY KEY,
			title VARCHAR(255) NOT NULL,
			artist VARCHAR(255) NOT NULL,
			year INT,
			museum_id INT REFERENCES museums(id)
		);
		`,
	}

	for i, migration := range migrations {
		if err := DB.Exec(migration).Error; err != nil {
			log.Fatalf("Failed to execute migration #%d: %v", i+1, err)
		}
		log.Printf("Migration #%d executed successfully.", i+1)
	}
}