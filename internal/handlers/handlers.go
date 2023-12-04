package handlers

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"radiophysiker/link_shorter/internal/storage"
)

type URLHandler struct {
	Storage storage.URLCreatorGetter
}

func (h *URLHandler) CreateShortURL(c *fiber.Ctx) error {
	var url = string(c.BodyRaw())
	if url == "" {
		return c.Status(http.StatusBadRequest).SendString("url is empty")
	}
	shortURL, err := h.Storage.CreateShortURL(url)
	if err != nil {
		return c.SendStatus(http.StatusBadRequest)
	}
	return c.Status(http.StatusCreated).SendString("http://localhost:8080/" + shortURL)
}

func (h *URLHandler) GetFullURL(c *fiber.Ctx) error {
	shortURL := c.Params("id")
	fullURL, err := h.Storage.GetFullURL(shortURL)
	if err != nil {
		return c.Status(http.StatusNotFound).SendString("url is not found for " + shortURL)
	}
	return c.Redirect(fullURL, http.StatusTemporaryRedirect)
}
