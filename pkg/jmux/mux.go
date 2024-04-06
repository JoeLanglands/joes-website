package jmux

import (
	"context"
	"log/slog"
	"net/http"
)

type Mux struct {
	mux             *http.ServeMux
	logger          *slog.Logger
	notFoundHandler http.Handler
}

type LoggerKey struct{}

// GetLogger returns the slog.Logger from the request context.
// This can be used to access the logger instance injected into the request
// context by the WithLogger option. If
func GetLogger(r *http.Request) *slog.Logger {
	if v := r.Context().Value(LoggerKey{}); v != nil {
		return v.(*slog.Logger)
	}
	return slog.Default()
}

func NewMux(opts ...MuxOption) *Mux {
	base := &Mux{
		mux: http.NewServeMux(),
	}

	for _, opt := range opts {
		opt(base)
	}

	return base
}

func (m *Mux) Handle(pattern string, handler http.Handler) {
	m.mux.Handle(pattern, handler)
}

// Get registers a handler function for the pattern and only the GET method.
// Non-GET requests will be rejected with a 405 status code.
func (m *Mux) Get(pattern string, handler http.Handler) {
	m.mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		handler.ServeHTTP(w, r)
	})
}

// Post registers a handler function for the pattern and only the POST method.
// Non-POST requests will be rejected with a 405 status code.
func (m *Mux) Post(pattern string, handler http.Handler) {
	m.mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		handler.ServeHTTP(w, r)
	})
}

// Put registers a handler function for the pattern and only the PUT method.
// Non-PUT requests will be rejected with a 405 status code.
func (m *Mux) Put(pattern string, handler http.Handler) {
	m.mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		handler.ServeHTTP(w, r)
	})
}

// Delete registers a handler function for the pattern and only the DELETE method.
// Non-DELETE requests will be rejected with a 405 status code.
func (m *Mux) Delete(pattern string, handler http.Handler) {
	m.mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		handler.ServeHTTP(w, r)
	})
}

// GetFunc registers a handler function for the pattern and only the GET method.
func (m *Mux) GetFunc(pattern string, handlerFunc http.HandlerFunc) {
	m.Get(pattern, handlerFunc)
}

// PostFunc registers a handler function for the pattern and only the POST method.
func (m *Mux) PostFunc(pattern string, handlerFunc http.HandlerFunc) {
	m.Post(pattern, handlerFunc)
}

// PutFunc registers a handler function for the pattern and only the PUT method.
func (m *Mux) PutFunc(pattern string, handlerFunc http.HandlerFunc) {
	m.Put(pattern, handlerFunc)
}

// DeleteFunc registers a handler function for the pattern and only the DELETE method.
func (m *Mux) DeleteFunc(pattern string, handlerFunc http.HandlerFunc) {
	m.Delete(pattern, handlerFunc)
}

// NotFoundHandler registers a handler for when no routes match the request.
// This is useful for returning a 404 page or similar.
// Alternative to using WithNotFoundHandler option.
func (m *Mux) NotFoundHandler(h http.Handler) {
	m.notFoundHandler = h
}

func (m *Mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	req := r
	if m.logger != nil {
		ctx := context.WithValue(r.Context(), LoggerKey{}, m.logger)
		req = r.WithContext(ctx)
	}
	h, p := m.mux.Handler(req)

	m.logger.Info("Request", "handler", h, "pattern", p)

	if h == http.NotFoundHandler() {
		m.logger.Info("Not found handler", "code", http.StatusNotFound)
	}

	h.ServeHTTP(w, req)
	//m.mux.ServeHTTP(w, req)
}
