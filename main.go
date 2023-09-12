package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	godotenv.Load(".env")
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not found in the environment")
	}

	v1Router := chi.NewRouter()
	v1Router.HandleFunc("/healthz", handlerReadiness)
	router := chi.NewRouter()

	router.Mount("/v1", v1Router)
	router.Get("/", handlerReadiness)
	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
