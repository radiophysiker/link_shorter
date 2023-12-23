package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gofiber/fiber/v2"

	"radiophysiker/link_shorter/internal/config"
	"radiophysiker/link_shorter/internal/storage"
)

type URLHandler struct {
	storage storage.URLCreatorGetter
	config  config.Getter
}

type CreateShortURLEntryRequest struct {
	FullURL string `json:"url"`
}

type CreateShortURLEntryResponse struct {
	ShortURL string `json:"result"`
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

func (h *URLHandler) CreateShortAPIURL(c *fiber.Ctx) error {
	var request CreateShortURLEntryRequest
	json.Unmarshal(c.BodyRaw(), &request)
	var url = request.FullURL
	if url == "" {
		return c.Status(http.StatusBadRequest).SendString("url is empty")
	}
	shortURL := h.storage.CreateShortURL(url)
	return c.Status(http.StatusCreated).JSON(CreateShortURLEntryResponse{ShortURL: h.config.GetBaseURL() + "/" + shortURL})
}

func (h *URLHandler) GetFullURL(c *fiber.Ctx) error {
	shortURL := c.Params("id")
	fullURL, err := h.storage.GetFullURL(shortURL)
	if err != nil {
		return c.Status(http.StatusNotFound).SendString("url is not found for " + shortURL)
	}
	return c.Redirect(fullURL, http.StatusTemporaryRedirect)
}
