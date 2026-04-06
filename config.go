package main

import (
	"database/sql"
	"log"
	"os"
	"sync/atomic"

	"github.com/derrickyoo/go-chirpy/internal/database"
	"github.com/joho/godotenv"
)

type config struct {
	db             *database.Queries
	fileserverHits atomic.Int32
	filepathRoot   string
	platform       string
	port           string
}

func mustEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf(`Environment variable %s is not set`, key)
	}
	return value
}

func defineConfig() *config {
	godotenv.Load()

	dbURL := mustEnv("DB_URL")
	dbConn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Error opening database: %s", err)
	}
	dbQueries := database.New(dbConn)

	return &config{
		db:             dbQueries,
		fileserverHits: atomic.Int32{},
		filepathRoot:   mustEnv("FILEPATH_ROOT"),
		port:           mustEnv("PORT"),
		platform:       mustEnv("PLATFORM"),
	}
}
