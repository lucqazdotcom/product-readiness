package api

import (
	"log"
	"net/http"
	i "product-readiness/internal/ingestion"
	"time"
)

func Api() {

	mux := http.NewServeMux()

	mux.HandleFunc("GET /ingestion", ingestionFeedHandler)
	mux.HandleFunc("GET /health", healthHandler)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	log.Println("Server listening on http://localhost:8080")
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server failed: %s", err)
	}

}

// NOTE: Endpoint for ingestion polling after intial startup call to
// simulate an existing pipeline
func ingestionFeedHandler(w http.ResponseWriter, r *http.Request) {
	i.Ingestion("testdata")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status": "Complete"}`))

}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status": "healthy"}`))
}
