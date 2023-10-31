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
	"github.com/lib/pq"
	"github.com/prabal01pathak/scratch/internal/database"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
	d := pq.Driver{}
	fmt.Println(d)
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
	fmt.Printf("db url is: %s", dbURL)
	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Error while connecting with the database: ", err)
	}
	apiCfg := apiConfig{
		DB: database.New(conn),
	}
	v1Router.Post("/create", apiCfg.handlerUser)
	v1Router.Get("/user", apiCfg.middlewareAuth(apiCfg.handlerGetUser))
	v1Router.Post("/feeds", apiCfg.middlewareAuth(apiCfg.handlerFeed))
	v1Router.Get("/feeds", apiCfg.handleGetAllFeed)
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
