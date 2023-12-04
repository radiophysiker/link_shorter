package storage

import (
	"errors"
	"radiophysiker/link_shorter/internal/utils"
)

type (
	ShortURL = string
	FullURL  = string
)

type URLStorage struct {
	Urls map[ShortURL]FullURL
}

func (s URLStorage) CreateShortURL(fullURL FullURL) (ShortURL, error) {
	shortURL := utils.GetShortRandomString()
	s.Urls[shortURL] = fullURL
	return shortURL, nil
}

func (s URLStorage) GetFullURL(shortURL ShortURL) (FullURL, error) {
	fullURL, ok := s.Urls[shortURL]
	if !ok {
		return "", errors.New("url=" + shortURL + " not found")
	}
	return fullURL, nil
}

type URLCreatorGetter interface {
	CreateShortURL(fullURL FullURL) (ShortURL, error)
	GetFullURL(shortURL ShortURL) (FullURL, error)
}
