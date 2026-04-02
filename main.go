package main

import (
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	cfg := defineConfig()

	mux := http.NewServeMux()
	mux.Handle("GET /app/", cfg.middlewareMetricsInc(http.StripPrefix("/app", http.FileServer(http.Dir(cfg.filepathRoot)))))

	mux.HandleFunc("GET /api/healthz", handlerReadiness)
	mux.HandleFunc("POST /api/validate_chirp", handlerChirpsValidate)

	mux.HandleFunc("GET /admin/metrics", cfg.handlerMetrics)
	mux.HandleFunc("POST /admin/reset", cfg.handlerReset)

	server := &http.Server{
		Addr:    ":" + cfg.port,
		Handler: mux,
	}

	log.Printf("Server is running at http://localhost:%s\n", cfg.port)
	log.Fatal(server.ListenAndServe())
}
