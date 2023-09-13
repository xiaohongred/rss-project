package main

import (
	"database/sql"
	"github.com/go-chi/chi/v5"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/xiaohongred/rss-project/internal/database"
	"log"
	"net/http"
	"os"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load(".env")
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not found in the environment")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL is not found in the environment")
	}

	conn, err := sql.Open("mysql", dbURL)
	if err != nil {
		log.Fatal("sql.Open err", err)
	}
	queries := database.New(conn)
	apiCfg := apiConfig{
		DB: queries,
	}

	v1Router := chi.NewRouter()
	v1Router.HandleFunc("/healthz", handlerReadiness)
	v1Router.HandleFunc("/err", handlerErr)
	v1Router.Post("/users", apiCfg.handlerCreateUser)
	router := chi.NewRouter()
	router.Mount("/v1", v1Router)
	router.Get("/", handlerReadiness)
	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}
	log.Printf("listening on %s", ":"+portString)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
