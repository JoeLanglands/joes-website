package middlewares

import (
	"github.com/JoeLanglands/joes-website/pkg/jmux"
	"log/slog"
	"net/http"
	"time"
)

// RequestLogging is a middleware that logs the request path, method and time taken to handle the request.
// The logger used is the one provided in the context, if it exists. If not, the default logger (from slog) is used.
func RequestLogging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger, ok := r.Context().Value(jmux.LoggerKey{}).(*slog.Logger)
		if !ok || logger == nil {
			logger = slog.Default()
		}
		start := time.Now()
		next.ServeHTTP(w, r)
		logger.Info("Request handled", "path", r.RequestURI, "method", r.Method, "latency", time.Since(start).String())
	})
}
