package middleware

import (
	"net/http"
)

func Middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// TODO: Add middleware stuff here! Auth, validation
		// log.Info("Middleware", zap.Any("URL", r.URL), zap.Any("User", strings.ToLower(r.Header.Get("Uuid"))))
		h.ServeHTTP(w, r)
	})
}
