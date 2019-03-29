package middleware

import (
	"log"
	"net/http"
)

// LoggingMiddleware is middleware that logs the request URI
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
