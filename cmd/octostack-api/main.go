package main

import (
	"log"
	"net"
	"net/http"
	"os"

	chi "github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/octostack/api/endpoint"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})

	r.Post("/auth/login", endpoint.Login)
	r.Post("/auth/signup", endpoint.Register)
	r.Post("/auth/token", endpoint.Token)

	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8080"
	}

	srv := &http.Server{
		Addr:    net.JoinHostPort("", port),
		Handler: r,
	}

	log.Printf("http server listening at %s\n", srv.Addr)
	if err := srv.ListenAndServe(); err != nil {
		log.Printf("http server closed unexpectedly: %v", err)
	}
}
