package logger

import (
	"fmt"
	"net/http"
	"time"

	"go.uber.org/zap"
)

type Log struct {
	*zap.SugaredLogger
}

type Logger interface {
	Fatal(format string, args ...any)
	Error(format string, args ...any)
	Info(format string, args ...any)
	CustomMiddlewareLogger(next http.Handler) http.Handler
}

func Init() (*Log, error) {
	z, err := zap.NewDevelopment()
	if err != nil {
		return nil, fmt.Errorf("logger don't Run! %s", err)
	}
	logger := z.Sugar()
	defer logger.Sync()
	return &Log{logger}, nil
}

func (l *Log) Fatal(format string, args ...any) {
	l.Fatalf(format, args...)
}

func (l *Log) Error(format string, args ...any) {
	l.Errorf(format, args...)
}

func (l *Log) Info(format string, args ...any) {
	l.Infof(format, args...)
}

func (l *Log) CustomMiddlewareLogger(next http.Handler) http.Handler {
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
