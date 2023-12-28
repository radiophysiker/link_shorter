package repository

import "radiophysiker/link_shorter/internal/entity"

//go:generate mockery --name=URLFileRepository --output=../mocks --filename=fs.go
type URLFileRepository interface {
	Save(url entity.URL) error
	GetFullURL(shortURL string) (string, error)
}
