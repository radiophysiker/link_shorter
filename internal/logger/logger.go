package logger

import (
	"fmt"
	"net/http"
	"time"

	"go.uber.org/zap"
)

var logger *zap.SugaredLogger

func Init() error {
	var err error
	z, err := zap.NewDevelopment()
	if err != nil {
		return fmt.Errorf("logger don't Run! %s", err)
	}
	logger = z.Sugar()
	defer logger.Sync()
	return nil
}

func Fatalf(format string, args ...any) {
	logger.Fatalf(format, args...)
}

func CustomMiddlewareLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)

		logger.Infoln(
			"URI:", r.RequestURI,
			"Method:", r.Method,
			"Duration:", time.Since(start))
	})
}
