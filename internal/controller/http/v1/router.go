package v1

import (
	"github.com/go-chi/chi"

	"radiophysiker/link_shorter/internal/config"
	"radiophysiker/link_shorter/internal/handlers"
	"radiophysiker/link_shorter/internal/logger"
	"radiophysiker/link_shorter/internal/middleware"
	"radiophysiker/link_shorter/internal/usecases"
)

func NewRouter(u usecases.URL, cfg *config.Config, log logger.Logger) *chi.Mux {
	r := chi.NewRouter()

	r.Use(log.CustomMiddlewareLogger)
	r.Use(middleware.CustomCompression)
	urlHandler := handlers.NewURLHandler(u, cfg, log)
	urlHandler.RegisterRoutes(r)
	return r
}
