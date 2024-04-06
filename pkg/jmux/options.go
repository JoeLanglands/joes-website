package jmux

import (
	"log/slog"
	"net/http"
)

type MuxOption func(m *Mux)

// WithLogger injects a slog.Logger into the request context for each request.
// This can be used by middleware or in handlers to log information.
func WithLogger(l *slog.Logger) func(*Mux) {
	return func(m *Mux) {
		m.logger = l
	}
}

// WithNotFoundHandler sets the handler to be used when no routes match the request.
func WithNotFoundHandler(h http.Handler) func(*Mux) {
	return func(m *Mux) {
		m.notFoundHandler = h
	}
}
