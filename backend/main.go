package main

import (
	"fmt"
	"log"
	"net/http"

	// configs "cdn-project/Configs"

	routes "cdn-project/Routes"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	router := mux.NewRouter()
	apiRouter := router.PathPrefix("/api").Subrouter()

	// Enable CORS using rs/cors middleware
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Authorization", "Content-Type"},
	})

	// MongoDB Connection
	// configs.ConnectDB()

	routes.MemberRoutes(apiRouter)

	handler := c.Handler(router)

	fmt.Print("Server is running on port 8082 !!!")
	log.Fatal(http.ListenAndServe("0.0.0.0:8082", handler))
}
