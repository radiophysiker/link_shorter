package repository

import "radiophysiker/link_shorter/internal/entity"

type URLFileRepository interface {
	Save(url entity.URL) error
	GetFullURL(shortURL string) (string, error)
}
