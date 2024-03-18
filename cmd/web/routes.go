package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/hculpan/go-web-base-project-1/cmd/web/handlers"
	"github.com/hculpan/go-web-base-project-1/internal/auth"
)

func routes(router *chi.Mux) *chi.Mux {
	if router != nil {
		router.Get("/", handlers.HomeHandler)
		router.Get("/info", handlers.InfoHandler)
		router.Get("/login", handlers.LoginHandler)
		router.Post("/login", handlers.LoginPostHandler)
		router.Get("/logout", handlers.LogoutHandler)

		router.Get("/home", auth.AuthMiddleware(handlers.HomeHandler))

		// Serve static files
		fileServer := http.FileServer(http.Dir("./static"))
		router.Handle("/static/*", http.StripPrefix("/static", fileServer))
	}

	return router
}
