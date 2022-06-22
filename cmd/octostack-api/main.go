package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	chi "github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/octostack/api/endpoint"
)

func main() {
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hello world")
	})

	authEndpoint := endpoint.NewAuthEndpoint()

	r.Post("/auth/login", authEndpoint.Login)
	r.Post("/auth/signup", authEndpoint.Register)
	r.Post("/auth/token", authEndpoint.Token)

	userEndpoint := endpoint.NewUserEndpoint()

	r.Get("/users", userEndpoint.List)
	r.Get("/users/{username}", userEndpoint.Get)
	r.Put("/users/{username}", userEndpoint.Update)
	r.Delete("/users/{username}", userEndpoint.Delete)
	r.Get("/users/{username}/hovercard", userEndpoint.GetHovercard)

	orgEndpoint := endpoint.NewOrgEndpoint()

	r.Get("/orgs", orgEndpoint.List)
	r.Get("/orgs/{org}", orgEndpoint.Get)

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
		if err != http.ErrServerClosed {
			log.Printf("http server closed unexpectedly: %v", err)
		}
	}
}
