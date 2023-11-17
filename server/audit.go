package server

import (
	"log/slog"
	"net/http"
	"time"
)

// Audit interceptor to log method calls.
func Audit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// execute method
		start := time.Now()
		next.ServeHTTP(w, req)

		// log completion and time elapsed
		slog.Info("Handled request",
			"verb", req.Method,
			"method", req.URL.Path,
			"elapsed", time.Since(start))
	})
}
