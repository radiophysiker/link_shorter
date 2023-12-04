package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"radiophysiker/link_shorter/internal/handlers"
	"radiophysiker/link_shorter/internal/storage"
)

func main() {
	urlHandler := &handlers.URLHandler{
		Storage: &storage.URLStorage{
			Urls: make(map[storage.ShortURL]storage.FullURL),
		},
	}
	webApp := fiber.New()
	webApp.Post("/", urlHandler.CreateShortURL)
	webApp.Get("/:id", urlHandler.GetFullURL)
	logrus.Fatal(webApp.Listen("localhost:8080"))
}
