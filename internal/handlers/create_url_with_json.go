package handlers

import (
	"encoding/json"
	"io"
	"net/http"
)

type CreateShortURLEntryRequest struct {
	FullURL string `json:"url"`
}

type CreateShortURLEntryResponse struct {
	ShortURL string `json:"result"`
}

func (h *URLHandler) CreateShortURLWithJSON(w http.ResponseWriter, r *http.Request) {
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
	shortURL, err := h.URLUseCase.CreateShortURL(url)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	baseURL := h.config.GetBaseURL()
	resp := CreateShortURLEntryResponse{ShortURL: baseURL + "/" + shortURL}

	jsonResp, err := json.Marshal(resp)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonResp)
}
