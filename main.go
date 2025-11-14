package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Hello, World!")
	godotenv.Load(".env")
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT environment variable not set")
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/health", handlerReadiness)
	v1Router.Get("/err",handleErr)


	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
		// WriteTimeout: 15 * time.Second,
		// ReadTimeout:  15 * time.Second,
	}
	log.Printf("Starting server on port %s", portString)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
