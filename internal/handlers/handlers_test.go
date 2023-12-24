package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"

	"radiophysiker/link_shorter/internal/config"
	"radiophysiker/link_shorter/internal/storage"
)

func TestUrlHandlerCreateShortUrlSimple(t *testing.T) {
	tests := []struct {
		name     string
		body     string
		wantCode int
	}{
		{
			name:     "simple",
			body:     "ya.ru",
			wantCode: http.StatusCreated,
		},
		{
			name:     "empty",
			body:     "",
			wantCode: http.StatusBadRequest,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			router := chi.NewRouter()
			handler := &URLHandler{
				storage: &storage.URLStorage{
					Urls: make(map[string]string),
				},
				config: &config.Config{
					BaseURL:    "localhost:8080",
					ServerPort: "localhost:8080",
				},
			}
			router.Post("/", handler.CreateShortURL)
			req := httptest.NewRequest(
				http.MethodPost,
				"/",
				strings.NewReader(tc.body),
			)
			rec := httptest.NewRecorder()
			router.ServeHTTP(rec, req)

			assert.Equal(t, tc.wantCode, rec.Code)
			assert.NotEmpty(t, rec.Body)
		})
	}
}

func TestUrlHandlerGetFullUrl(t *testing.T) {
	tests := []struct {
		name        string
		requestPath string
		wantCode    int
		wantBody    string
	}{
		{
			name:        "simple",
			requestPath: "/test",
			wantCode:    http.StatusTemporaryRedirect,
			wantBody:    "test",
		},
		{
			name:        "not found shortURL",
			requestPath: "/t",
			wantCode:    http.StatusNotFound,
			wantBody:    "test",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			handler := &URLHandler{
				storage: &storage.URLStorage{
					Urls: map[string]string{"test": "test"},
				},
				config: &config.Config{
					BaseURL:    "localhost:8080",
					ServerPort: "localhost:8080",
				},
			}
			router := chi.NewRouter()
			router.Get("/{id}", handler.GetFullURL)

			req := httptest.NewRequest(
				http.MethodGet,
				tc.requestPath,
				nil,
			)
			rec := httptest.NewRecorder()
			router.ServeHTTP(rec, req)

			assert.Equal(t, tc.wantCode, rec.Code)
		})
	}
}

func TestUrlHandlerCreateShortAPIURL(t *testing.T) {
	tests := []struct {
		name     string
		body     string
		wantCode int
	}{
		{
			name:     "simple",
			body:     `{"url": "https://yandex.ru"}`,
			wantCode: http.StatusCreated,
		},
		{
			name:     "empty",
			body:     "",
			wantCode: http.StatusBadRequest,
		},
	}

	router := chi.NewRouter()
	handler := &URLHandler{
		storage: &storage.URLStorage{
			Urls: make(map[string]string),
		},
		config: &config.Config{
			BaseURL:    "localhost:8080",
			ServerPort: "localhost:8080",
		},
	}
	router.Post("/", handler.CreateShortAPIURL)
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			req := httptest.NewRequest(
				http.MethodPost,
				"/",
				strings.NewReader(tc.body),
			)
			rec := httptest.NewRecorder()
			router.ServeHTTP(rec, req)

			assert.Equal(t, tc.wantCode, rec.Code)
			assert.NotEmpty(t, rec.Body)
		})
	}
}
