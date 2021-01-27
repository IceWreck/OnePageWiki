package main

import (
	"net/http"
	"time"

	"github.com/IceWreck/OnePageWiki/config"
	"github.com/IceWreck/OnePageWiki/handlers"
	"github.com/IceWreck/OnePageWiki/logger"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	// View Markdown
	r.Get("/", handlers.MarkdownView)
	// Editor
	r.Route("/", func(r chi.Router) {
		r.Use(middleware.BasicAuth("Credentials:", config.Credentials))
		r.Get("/edit", handlers.EditView)
		r.Post("/edit", handlers.EditForm)
	})

	fileServer(r, "/static", http.Dir("./static"))

	logger.Info("Starting at " + config.Port)
	err := http.ListenAndServe(config.Port, r)
	if err != nil {
		logger.Error(err)
	}
}
