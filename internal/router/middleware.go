package router

import (
	"net/http"
)

func OnlyServeHTMX(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		htmxHead := r.Header.Get("HX-Request")
		if htmxHead == "true" {
			next.ServeHTTP(w, r)
			return
		}
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	})
}
