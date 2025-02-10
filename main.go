package main

import (
<<<<<<< HEAD
=======
	"cdn-project/database"
>>>>>>> 962533d (feat: add mongodb)
	"fmt"
	"log"
	"net/http"
	"time"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %s", r.Method, r.RequestURI, time.Since(start))
<<<<<<< HEAD
=======
		fmt.Println("Yoyo")
>>>>>>> 962533d (feat: add mongodb)
	})
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "OK")
}

func main() {
	db := database.ConnectDB()
	fmt.Println("Base de données prête :", db.Name())

	mux := http.NewServeMux()
	mux.HandleFunc("/health", HealthCheckHandler)

	loggedMux := loggingMiddleware(mux)

	fmt.Println("Serveur démarré sur :8080")
	log.Fatal(http.ListenAndServe(":8080", loggedMux))
}
