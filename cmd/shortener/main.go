package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	. "radiophysiker/link_shorter/internal/handlers"
	. "radiophysiker/link_shorter/internal/storage"
)

func main() {
	urlHandler := &UrlHandler{
		Storage: &UrlStorage{
			Urls: make(map[ShortUrl]FullUrl),
		},
	}
	webApp := fiber.New()
	webApp.Post("/", urlHandler.CreateShortUrl)
	webApp.Get("/:id", urlHandler.GetFullUrl)
	logrus.Fatal(webApp.Listen(":8080"))
}
