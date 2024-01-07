package repository

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/google/uuid"

	"radiophysiker/link_shorter/internal/config"
	"radiophysiker/link_shorter/internal/entity"
)

type urlRecord struct {
	UUID        string `json:"uuid"`
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
}

type URLFileRepository struct {
	config *config.Config
	urls   map[string]string
}

func NewURLFileRepository(cfg *config.Config) (*URLFileRepository, error) {
	repo := &URLFileRepository{
		config: cfg,
		urls:   make(map[string]string),
	}

	if err := repo.loadFromFile(); err != nil {
		return nil, err
	}

	return repo, nil
}

func (r *URLFileRepository) Save(url entity.URL) error {
	fullURL := url.FullURL
	if fullURL == "" {
		return errors.New("empty full URL")
	}
	r.urls[url.ShortURL] = fullURL
	return r.writeRecordToFile(url)
}

func (r *URLFileRepository) GetFullURL(shortURL string) (string, error) {
	if shortURL == "" {
		return "", errors.New("empty short URL")
	}
	fullURL, exists := r.urls[shortURL]
	if !exists {
		return "", errors.New("URL not found for " + shortURL)
	}
	return fullURL, nil
}

func (r *URLFileRepository) writeRecordToFile(url entity.URL) error {
	filePath := r.config.GetFileStoragePath()
	if filePath == "" {
		return nil
	}

	record := urlRecord{
		UUID:        uuid.New().String(),
		ShortURL:    url.ShortURL,
		OriginalURL: url.FullURL,
	}

	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer f.Close()

	if err := json.NewEncoder(f).Encode(&record); err != nil {
		return err
	}
	return nil
}

func (r *URLFileRepository) loadFromFile() error {
	filePath := r.config.GetFileStoragePath()
	if filePath == "" {
		return nil
	}
	f, err := os.OpenFile(filePath, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer f.Close()
	dec := json.NewDecoder(f)
	for dec.More() {
		var record urlRecord
		if err := dec.Decode(&record); err != nil {
			return err
		}

		r.urls[record.ShortURL] = record.OriginalURL
	}

	return nil
}
