package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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
	router.Get("/", httpRequestHandler)
	http.ListenAndServe(port, router)
}
