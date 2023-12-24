package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"

	"radiophysiker/link_shorter/internal/config"
	"radiophysiker/link_shorter/internal/handlers"
	"radiophysiker/link_shorter/internal/logger"
	"radiophysiker/link_shorter/internal/middleware"
)

func main() {
	cfg := config.New()
	cfgGetter := config.Getter(cfg)
	urlHandler := handlers.New(cfg)
	webApp := chi.NewRouter()
	if err := logger.Init(); err != nil {
		log.Fatalf("cannot initialize logger! %s", err)
	}

	webApp.Use(logger.CustomMiddlewareLogger)
	webApp.Use(middleware.CustomCompression)
	webApp.Post("/", urlHandler.CreateShortURL)
	webApp.Get("/{id}", urlHandler.GetFullURL)
	webApp.Post("/api/shorten", urlHandler.CreateShortAPIURL)

	logger.Fatalf("cannot initialize app!", http.ListenAndServe(cfgGetter.GetServerPort(), webApp))
}
