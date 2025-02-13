package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	routes "cdn-project/Routes"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// Middleware pour logger chaque requête
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Création d'un wrapper pour capturer le code HTTP retourné
		recorder := &responseRecorder{ResponseWriter: w, statusCode: http.StatusOK}
		next.ServeHTTP(recorder, r)

		duration := time.Since(start)

		// Log de la requête
		log.Printf(
			"[%s] %s %s - %d %s - %s",
			r.Method,
			r.RemoteAddr,
			r.RequestURI,
			recorder.statusCode,
			http.StatusText(recorder.statusCode),
			duration,
		)
	})
}

// responseRecorder permet de capturer le code HTTP de la réponse
type responseRecorder struct {
	http.ResponseWriter
	statusCode int
}

func (rr *responseRecorder) WriteHeader(code int) {
	rr.statusCode = code
	rr.ResponseWriter.WriteHeader(code)
}

func main() {
	router := mux.NewRouter()
	apiRouter := router.PathPrefix("/api").Subrouter()

	// Middleware CORS
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Authorization", "Content-Type"},
	})

	// Ajouter les routes
	routes.MemberRoutes(apiRouter)
	routes.FileRoutes(apiRouter)

	// Appliquer le middleware de logging
	loggedRouter := LoggingMiddleware(router)

	// Activer CORS sur le routeur loggé
	handler := c.Handler(loggedRouter)

	// Définir le port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Défaut si non défini
	}

	fmt.Println("Server is running on port " + port + " !!!")
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
