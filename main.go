package main

import (
<<<<<<< HEAD
<<<<<<< HEAD
=======
	"cdn-project/database"
>>>>>>> 962533d (feat: add mongodb)
=======
>>>>>>> bab9c8b (feat: add mongodb)
	"fmt"
	"log"
	"net/http"

	routes "github.com/hugolhld/cdn-project/Routes"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

<<<<<<< HEAD
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

=======
>>>>>>> bab9c8b (feat: add mongodb)
func main() {
	router := mux.NewRouter()

	// Enable CORS using rs/cors middleware
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Authorization", "Content-Type"},
	})

	// MongoDB Connection
	// configs.ConnectDB()

	routes.MemberRoutes(router)

	handler := c.Handler(router)

	fmt.Print("Server is running on port 8000 !!!")
	log.Fatal(http.ListenAndServe(":8000", handler))
}
