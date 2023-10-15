package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func httpRequestHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Welcome"))
}

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Port is not found in the environment")
	}
	fmt.Printf("port is: %s", port)
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELTE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	router.Get("/", httpRequestHandler)
	srv := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}
	error := srv.ListenAndServe()
	if error != nil {
		log.Fatal(error)
	}
	http.ListenAndServe(port, router)
}
