package router

import (
	"context"
	"log/slog"
	"net/http"
)

type Mux struct {
	middlewares []Middleware
	mux         *http.ServeMux
	logger      *slog.Logger
}

type MuxOption func(m *Mux)

type LoggerKey struct{}

func WithLogger(l *slog.Logger) func(*Mux) {
	return func(m *Mux) {
		m.logger = l
	}
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
	m.Get(pattern, http.HandlerFunc(handlerFunc))
}

// PostFunc registers a handler function for the pattern and only the POST method.
func (m *Mux) PostFunc(pattern string, handlerFunc http.HandlerFunc) {
	m.Post(pattern, http.HandlerFunc(handlerFunc))
}

// PutFunc registers a handler function for the pattern and only the PUT method.
func (m *Mux) PutFunc(pattern string, handlerFunc http.HandlerFunc) {
	m.Put(pattern, http.HandlerFunc(handlerFunc))
}

// DeleteFunc registers a handler function for the pattern and only the DELETE method.
func (m *Mux) DeleteFunc(pattern string, handlerFunc http.HandlerFunc) {
	m.Delete(pattern, http.HandlerFunc(handlerFunc))
}

func (m *Mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), LoggerKey{}, m.logger)
	req := r.WithContext(ctx)
	m.mux.ServeHTTP(w, req)
}
