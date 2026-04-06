package main

import (
	"log"
	"net/http"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	cfg := defineConfig()

	mux := http.NewServeMux()
	mux.Handle("GET /app/", cfg.middlewareMetricsInc(http.StripPrefix("/app", http.FileServer(http.Dir(cfg.filepathRoot)))))

	mux.HandleFunc("GET /api/healthz", handlerReadiness)

	mux.HandleFunc("GET /admin/metrics", cfg.handlerMetrics)
	mux.HandleFunc("POST /admin/reset", cfg.handlerReset)
	mux.HandleFunc("POST /api/users", cfg.handlerUsersCreate)
	mux.HandleFunc("POST /api/chirps", cfg.handlerChirpsCreate)
	mux.HandleFunc("GET /api/chirps", cfg.handlerChirpsRetrieve)
	mux.HandleFunc("GET /api/chirps/{chirpID}", cfg.handlerChirpsGet)

	server := &http.Server{
		Addr:         ":" + cfg.port,
		Handler:      mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	log.Printf("Server is running at http://localhost:%s\n", cfg.port)
	log.Fatal(server.ListenAndServe())
}
