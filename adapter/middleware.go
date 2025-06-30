package adapter

import (
	"log"
	"net/http"
	"time"
)

func LoggingMiddleware(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		log.Printf("Request started: %s %s",
			r.Method,
			r.URL.Path,
		)

		handler.ServeHTTP(w, r)

		duration := time.Since(start)
		log.Printf("Request completed: %s %s - Duration: %v",
			r.Method,
			r.URL.Path,
			duration,
		)
	}
}
