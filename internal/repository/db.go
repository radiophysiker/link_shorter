package repository

import "radiophysiker/link_shorter/internal/config"

type URLDBRepository struct {
	config *config.Config
	urls   map[string]string
}

func NewURLDBRepository(cfg *config.Config) (*URLDBRepository, error) {
	repo := &URLDBRepository{
		config: cfg,
		urls:   make(map[string]string),
	}
	return repo, nil
}
