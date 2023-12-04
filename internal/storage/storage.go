package storage

import (
	"errors"
	"radiophysiker/link_shorter/internal/utils"
)

type (
	ShortUrl = string
	FullUrl  = string
)

type UrlStorage struct {
	Urls map[ShortUrl]FullUrl
}

func (s UrlStorage) CreateShortUrl(fullUrl FullUrl) (ShortUrl, error) {
	shortUrl := utils.GetShortRandomString()
	s.Urls[shortUrl] = fullUrl
	return shortUrl, nil
}

func (s UrlStorage) GetFullUrl(shortUrl ShortUrl) (FullUrl, error) {
	fullUrl, ok := s.Urls[shortUrl]
	if !ok {
		return "", errors.New("%v not found")
	}
	return fullUrl, nil
}

type UrlCreatorGetter interface {
	CreateShortUrl(url FullUrl) (ShortUrl, error)
	GetFullUrl(url ShortUrl) (FullUrl, error)
}
