package router

import (
	"log/slog"
	"net/http"
	"time"
)

type Middleware func(http.Handler) http.Handler

// Use adds a middleware to the Handler chain so a set of middleware can be
// applied to a single route.
func Use(mw Middleware, handler http.Handler) http.Handler {
	return mw(handler)
}

// UseStack creates a stack of middlewares. The last middleware in the slice is the first to be called.
func UseStack(mws ...Middleware) Middleware {
	return func(next http.Handler) http.Handler {
		for i := len(mws) - 1; i >= 0; i-- {
			x := mws[i]
			next = x(next)
		}
		return next
	}
}

func OnlyServeHTMX(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		htmxHead := r.Header.Get("HX-Request")
		if htmxHead == "true" {
			next.ServeHTTP(w, r)
			return
		}
		http.Error(w, "Method not allowed", http.StatusTeapot)
	})
}

func RequestLogging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger, ok := r.Context().Value(LoggerKey{}).(*slog.Logger)
		if !ok || logger == nil {
			logger = slog.Default()
		}
		start := time.Now()
		next.ServeHTTP(w, r)
		logger.Info("Request handled", "path", r.RequestURI, "handled in", time.Since(start).String())
	})
}
