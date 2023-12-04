package handlers

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"radiophysiker/link_shorter/internal/storage"
)

type UrlHandler struct {
	Storage storage.UrlCreatorGetter
}

func (h *UrlHandler) CreateShortUrl(c *fiber.Ctx) error {
	var url = c.BodyRaw()
	shortUrl, err := h.Storage.CreateShortUrl(string(url))
	if err != nil {
		return c.SendStatus(http.StatusBadRequest)
	}
	return c.Status(http.StatusCreated).SendString("http://localhost:8080/" + shortUrl)
}

func (h *UrlHandler) GetFullUrl(c *fiber.Ctx) error {
	shortUrl := c.Params("id")
	fullUrl, err := h.Storage.GetFullUrl(shortUrl)
	if err != nil {
		return c.SendStatus(http.StatusBadRequest)
	}
	return c.Redirect(fullUrl, http.StatusTemporaryRedirect)
}
