package middleware

import (
	"compress/gzip"
	"io"
	"net/http"
	"strings"

	"radiophysiker/link_shorter/internal/logger"
)

type gzipWriter struct {
	http.ResponseWriter
	Writer io.Writer
}

func (w gzipWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

var successCompressionContentType = []string{"application/json", "text/html"}

func CustomCompression(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		contentType := r.Header.Get("Content-Type")

		isContentTypeSupported := false
		for _, ct := range successCompressionContentType {
			if ct == contentType {
				isContentTypeSupported = true
				break
			}
		}
		if strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") &&
			isContentTypeSupported {
			logger.Infof("compressing response with gzip")
			gz, err := gzip.NewWriterLevel(w, gzip.BestSpeed)
			if err != nil {
				logger.Errorf("cannot create gzip writer! %s", err)
				io.WriteString(w, err.Error())
				return
			}
			defer gz.Close()

			w.Header().Set("Content-Encoding", "gzip")
			next.ServeHTTP(gzipWriter{ResponseWriter: w, Writer: gz}, r)
			return
		}

		if strings.Contains(r.Header.Get("Content-Encoding"), "gzip") {
			logger.Infof("compressing request with gzip")
			gz, err := gzip.NewReader(r.Body)
			if err != nil {
				logger.Errorf("cannot create gzip reader! %s", err)
				io.WriteString(w, err.Error())
				return
			}
			defer gz.Close()

			r.Body = gz
		}

		next.ServeHTTP(w, r)
	})
}
