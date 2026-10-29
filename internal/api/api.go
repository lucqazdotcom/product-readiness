package api

import (
	"log"
	"net/http"
	"product-readiness/internal/db"
	i "product-readiness/internal/ingestion"
	"time"
)

func Api(database *db.Pool) {

	mux := http.NewServeMux()

	mux.HandleFunc("GET /ingestion", func(w http.ResponseWriter, r *http.Request) {
		ingestionFeedHandler(w, r, database)
	})
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
func ingestionFeedHandler(w http.ResponseWriter, r *http.Request, database *db.Pool) {
	_, err := i.Ingestion(r.Context(), "testdata", database)
	if err != nil {
		log.Printf("ingestion failed: %v", err)
		http.Error(w, "ingestion failed", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status": "Complete"}`))

}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status": "healthy"}`))
}
