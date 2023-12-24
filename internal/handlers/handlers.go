package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/go-chi/chi"

	"radiophysiker/link_shorter/internal/config"
	"radiophysiker/link_shorter/internal/storage"
)

type URLHandler struct {
	storage storage.URLCreatorGetter
	config  config.Getter
}

type CreateShortURLEntryRequest struct {
	FullURL string `json:"url"`
}

type CreateShortURLEntryResponse struct {
	ShortURL string `json:"result"`
}

func New(cfg *config.Config) *URLHandler {
	return &URLHandler{
		storage: &storage.URLStorage{
			Urls: make(map[storage.ShortURL]storage.FullURL),
		},
		config: cfg,
	}
}

func (h *URLHandler) CreateShortURL(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	url := string(body)
	if url == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("url is empty"))
		return
	}
	shortURL := h.storage.CreateShortURL(url)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(h.config.GetBaseURL() + "/" + shortURL))
}

func (h *URLHandler) CreateShortAPIURL(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var request CreateShortURLEntryRequest
	json.Unmarshal(body, &request)
	var url = request.FullURL
	if url == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("url is empty"))
		return
	}
	shortURL := h.storage.CreateShortURL(url)
	resp := CreateShortURLEntryResponse{ShortURL: h.config.GetBaseURL() + "/" + shortURL}

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonResp)
}

func (h *URLHandler) GetFullURL(w http.ResponseWriter, r *http.Request) {
	shortURL := chi.URLParam(r, "id")
	fullURL, err := h.storage.GetFullURL(shortURL)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("url is not found for " + shortURL))
		return
	}
	w.Header().Set("Location", fullURL)
	w.WriteHeader(http.StatusTemporaryRedirect)
}
