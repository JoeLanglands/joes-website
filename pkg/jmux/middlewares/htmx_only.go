package middlewares

import "net/http"

// OnlyServeHTMX is a middleware that only allows requests with the HX-Request header set to true,
// therefore blocking requests not originating from HTMX. Disallowed requests are served with
// a 400 status code and an error message.
func OnlyServeHTMX(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		htmxHeader := r.Header.Get("HX-Request")
		if htmxHeader == "true" {
			next.ServeHTTP(w, r)
			return
		}
		http.Error(w, "Route should only be served via HTMX", http.StatusBadRequest)
	})
}

// OnlyServerHTMXHandler is a middleware that only allows requests with the HX-Request header set to true.
// If the request does not have the header set, it will serve the request with the provided handler: `htmxOnlyHandler`.
func OnlyServeHTMXHandler(htmxOnlyHandler http.Handler) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			htmxHeader := r.Header.Get("HX-Request")
			if htmxHeader == "true" {
				next.ServeHTTP(w, r)
				return
			}
			htmxOnlyHandler.ServeHTTP(w, r)
		})
	}
}
