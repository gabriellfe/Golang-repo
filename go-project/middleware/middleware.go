package middleware

import (
	"encoding/json"
	"fmt"
	"gabriellfe/dto"
	"log/slog"
	"net/http"
	"os"

	"github.com/google/uuid"
)

func SchemaValidatorMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

func ApplicationJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	}
}

func LoggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}).WithAttrs([]slog.Attr{
			slog.String(
				"traceid", uuid.NewString(),
			),
		}))
		slog.SetDefault(logger)
		slog.Info(fmt.Sprintf("Request URI: %s, Method: %s", r.URL.Path, r.Method))
		next.ServeHTTP(w, r)
	}
}

func AuthenticationMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		jwt := r.Header.Get("Authorization")
		if jwt != "123" {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(dto.ErrorDto{StatusCode: http.StatusForbidden, Message: "Authorization Invalid"})
			return
		}
		next.ServeHTTP(w, r)
	}
}
