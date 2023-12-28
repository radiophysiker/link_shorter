package handlers

import (
	"github.com/go-chi/chi"

	"radiophysiker/link_shorter/internal/config"
	"radiophysiker/link_shorter/internal/logger"
	"radiophysiker/link_shorter/internal/usecases"
)

type URLHandler struct {
	URLUseCase usecases.URL
	config     *config.Config
	log        logger.Logger
}

func NewURLHandler(u usecases.URL, cfg *config.Config, log logger.Logger) *URLHandler {
	return &URLHandler{
		URLUseCase: u,
		config:     cfg,
		log:        log,
	}
}

func (h *URLHandler) RegisterRoutes(r chi.Router) {
	r.Get("/{id}", h.GetFullURL)
	r.Post("/api/shorten", h.CreateShortURLWithJSON)
	r.Post("/", h.CreateShortURL)
}
