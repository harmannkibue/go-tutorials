package main

import (
	"context"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.Use(guidMiddleware)
	router.HandleFunc("/ishealthy", handleIsHealthy).Methods(http.MethodGet)
	http.ListenAndServe(":8080", router)
}

func handleIsHealthy(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	uuid := r.Context().Value("uuid")
	log.Printf("[%v] Returning 200 - Healthy", uuid)
	w.Write([]byte("Healthy"))
}

func guidMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		uuid := uuid.New()
		r = r.WithContext(context.WithValue(r.Context(), "uuid", uuid))
		next.ServeHTTP(w, r)
	})
}
