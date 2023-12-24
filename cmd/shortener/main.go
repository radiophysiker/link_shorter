package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"

	"radiophysiker/link_shorter/internal/config"
	"radiophysiker/link_shorter/internal/handlers"
	"radiophysiker/link_shorter/internal/logger"
	"radiophysiker/link_shorter/internal/middleware"
)

func main() {
	cfg := config.New()
	cfgGetter := config.Getter(cfg)
	urlHandler := handlers.New(cfg)
	webApp := fiber.New()
	if err := logger.Init(); err != nil {
		log.Fatalf("cannot initialize logger! %s", err)
	}

	webApp.Use(adaptor.HTTPMiddleware(logger.CustomMiddlewareLogger))
	webApp.Use(adaptor.HTTPMiddleware(middleware.CustomCompression))
	webApp.Post("/", urlHandler.CreateShortURL)
	webApp.Get("/:id", urlHandler.GetFullURL)
	webApp.Post("/api/shorten", urlHandler.CreateShortAPIURL)

	logger.Fatalf("cannot initialize app!", webApp.Listen(cfgGetter.GetServerPort()))
}
