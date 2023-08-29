package main

import (
	"net/http"

	"github.com/Kevonosdiaz/bnb-web/pkg/config"
	"github.com/Kevonosdiaz/bnb-web/pkg/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// routes uses the chi external package to process routing
func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	// Use middleware from chi pkg and from self-written middleware.go
	mux.Use(middleware.Recoverer)
	mux.Use(WriteToConsole)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	return mux
}
