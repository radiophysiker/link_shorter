package handlers

import (
	"io"
	"net/http"
)

func (h *URLHandler) CreateShortURL(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		h.log.Error("problem with read body: %s", err)
		return
	}
	url := string(body)
	if url == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("url is empty"))
		h.log.Error("url is empty")
		return
	}
	shortURL, err := h.URLUseCase.CreateShortURL(url)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		h.log.Error("problem with create short url: %s", err)
		return
	}
	w.WriteHeader(http.StatusCreated)
	baseURL := h.config.GetBaseURL()
	w.Write([]byte(baseURL + "/" + shortURL))
}
