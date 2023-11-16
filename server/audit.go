package server

import (
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

// Audit interceptor to log method calls.
func Audit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// execute method
		start := time.Now()
		next.ServeHTTP(w, req)
		elapsed := time.Since(start)

		// log completion and time elapsed
		log.WithFields(log.Fields{
			"verb":    req.Method,
			"method":  req.URL.Path,
			"elapsed": elapsed.Nanoseconds(),
		}).Info("Handled request")
	})
}
