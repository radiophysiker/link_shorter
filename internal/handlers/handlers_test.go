package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

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
			app := fiber.New()
			handler := &URLHandler{
				storage: &storage.URLStorage{
					Urls: make(map[string]string),
				},
				config: &config.Config{
					BaseURL:    "localhost:8080",
					ServerPort: "localhost:8080",
				},
			}
			app.Post("/", handler.CreateShortURL)
			req := httptest.NewRequest(
				http.MethodPost,
				"/",
				strings.NewReader(tc.body),
			)
			resp, err := app.Test(req)
			require.NoError(t, err)
			defer resp.Body.Close()
			assert.Equal(t, tc.wantCode, resp.StatusCode)
			assert.NotEmpty(t, resp.Body)
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
			app := fiber.New()
			handler := &URLHandler{
				storage: &storage.URLStorage{
					Urls: map[string]string{"test": "test"},
				},
				config: &config.Config{
					BaseURL:    "localhost:8080",
					ServerPort: "localhost:8080",
				},
			}
			app.Get("/:id", handler.GetFullURL)
			req := httptest.NewRequest(
				http.MethodGet,
				tc.requestPath,
				nil,
			)
			resp, err := app.Test(req)
			require.NoError(t, err)
			defer resp.Body.Close()
			assert.Equal(t, tc.wantCode, resp.StatusCode)
		})
	}
}
