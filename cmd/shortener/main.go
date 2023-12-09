package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"radiophysiker/link_shorter/internal/config"
	"radiophysiker/link_shorter/internal/handlers"
)

func main() {
	cfg := config.New()
	cfgGetter := config.Getter(cfg)
	urlHandler := handlers.New(cfg)
	webApp := fiber.New()
	webApp.Post("/", urlHandler.CreateShortURL)
	webApp.Get("/:id", urlHandler.GetFullURL)
	logrus.Fatal(webApp.Listen(cfgGetter.GetServerPort()))
}
