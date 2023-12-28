package logger

import (
	"fmt"
	"net/http"
	"time"

	"go.uber.org/zap"
)

type Logger struct {
	*zap.SugaredLogger
}

type Interface interface {
	Fatal(format string, args ...any)
	Error(format string, args ...any)
	Info(format string, args ...any)
	CustomMiddlewareLogger(next http.Handler) http.Handler
}

func Init() (*Logger, error) {
	z, err := zap.NewDevelopment()
	if err != nil {
		return nil, fmt.Errorf("logger don't Run! %s", err)
	}
	logger := z.Sugar()
	defer logger.Sync()
	return &Logger{logger}, nil
}

func (l *Logger) Fatal(format string, args ...any) {
	l.Fatalf(format, args...)
}

func (l *Logger) Error(format string, args ...any) {
	l.Errorf(format, args...)
}

func (l *Logger) Info(format string, args ...any) {
	l.Infof(format, args...)
}

func (l *Logger) CustomMiddlewareLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)

		l.Infof(
			"URI:", r.RequestURI,
			"Method:", r.Method,
			"Duration:", time.Since(start),
		)
	})
}

type Getter interface {
}
