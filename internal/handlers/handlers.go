package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"

	"radiophysiker/link_shorter/internal/config"
	"radiophysiker/link_shorter/internal/storage"
)

type URLHandler struct {
	storage storage.URLCreatorGetter
	config  config.Getter
}

func New(cfg *config.Config) *URLHandler {
	return &URLHandler{
		storage: &storage.URLStorage{
			Urls: make(map[storage.ShortURL]storage.FullURL),
		},
		config: cfg,
	}
}

func (h *URLHandler) CreateShortURL(c *fiber.Ctx) error {
	var url = string(c.BodyRaw())
	if url == "" {
		return c.Status(http.StatusBadRequest).SendString("url is empty")
	}
	shortURL := h.storage.CreateShortURL(url)
	return c.Status(http.StatusCreated).SendString(h.config.GetBaseURL() + "/" + shortURL)
}

func (h *URLHandler) GetFullURL(c *fiber.Ctx) error {
	shortURL := c.Params("id")
	fullURL, err := h.storage.GetFullURL(shortURL)
	if err != nil {
		return c.Status(http.StatusNotFound).SendString("url is not found for " + shortURL)
	}
	return c.Redirect(fullURL, http.StatusTemporaryRedirect)
}
