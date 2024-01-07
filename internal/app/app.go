package app

import (
	"log"
	"net/http"

	"radiophysiker/link_shorter/internal/config"
	v1 "radiophysiker/link_shorter/internal/controller/http/v1"
	"radiophysiker/link_shorter/internal/logger"
	"radiophysiker/link_shorter/internal/repository"
	"radiophysiker/link_shorter/internal/usecases"
)

func Run() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("cannot load config! %s", err)
	}
	l, err := logger.Init()
	if err != nil {
		log.Fatalf("cannot initialize logger! %s", err)
	}

	urlFileRepository, err := repository.NewURLFileRepository(cfg)
	if err != nil {
		l.Fatal("cannot initialize repository! %s", err)
	}
	useCasesURLShortener := usecases.NewURLShortener(urlFileRepository, cfg)
	router := v1.NewRouter(useCasesURLShortener, cfg, l)
	l.Info("starting server on port " + cfg.GetServerPort())
	l.Fatal("cannot initialize app!", http.ListenAndServe(cfg.GetServerPort(), router))
}
