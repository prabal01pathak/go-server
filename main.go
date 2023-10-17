package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/prabal01pathak/scratch/internal/database"
)

type apiConfig struct {
	DB *database.Queries
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
	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err", handlerErr)
	router.Mount("/v1", v1Router)
	// router.Get("/", handlerReadiness)
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("Database url is misssing please add that in the env")
	}
	conn, err := sql.Open("postgresql", dbURL)
	if err != nil {
		log.Fatal("Error while connecting with the database: ", err)
	}
	apiCfg := apiConfig{
		DB: database.New(conn),
	}
	v1Router.Post("/create", apiCfg.handlerUser)
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
