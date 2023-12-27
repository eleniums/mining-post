package server

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
)

// AccessLog interceptor to log method calls.
func AccessLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		start := time.Now()

		// wrap the ResponseWriter so we can retrieve the response status
		ww := middleware.NewWrapResponseWriter(w, req.ProtoMajor)

		// execute method
		next.ServeHTTP(ww, req)

		// log completion and time elapsed
		slog.Info("Handled request",
			"verb", req.Method,
			"method", req.URL.Path,
			"elapsed", time.Since(start),
			"status", ww.Status(),
		)
	})
}
