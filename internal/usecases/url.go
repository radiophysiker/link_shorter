package usecases

import (
	"radiophysiker/link_shorter/internal/config"
	"radiophysiker/link_shorter/internal/entity"
	"radiophysiker/link_shorter/internal/repository"
	"radiophysiker/link_shorter/internal/utils"
)

type URLUseCase struct {
	urlRepository repository.URLFileRepository
	config        *config.Config
}

func NewURLShortener(repo repository.URLFileRepository, config *config.Config) *URLUseCase {
	return &URLUseCase{
		urlRepository: repo,
		config:        config,
	}
}

func (us URLUseCase) CreateShortURL(fullURL string) (string, error) {
	shortURL := utils.GetShortRandomString()
	url := entity.URL{
		ShortURL: shortURL,
		FullURL:  fullURL,
	}
	err := us.urlRepository.Save(url)
	if err != nil {
		return "", err
	}
	return shortURL, nil
}

func (us URLUseCase) GetFullURL(shortURL string) (string, error) {
	return us.urlRepository.GetFullURL(shortURL)
}