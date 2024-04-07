package jmux

import (
	"context"
	"fmt"
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

// Handle registers a http.Handler to the route given by pattern.
// This is a proxy to the underlying http.ServeMux http.Handle method.
func (m *Mux) Handle(pattern string, handler http.Handler) {
	m.mux.Handle(pattern, handler)
}

// HandleFunc registers a http.HandlerFunc to the route given by pattern.
// This is a proxy to the underlying http.ServeMux HandleFunc method.
func (m *Mux) HandleFunc(pattern string, handler func(w http.ResponseWriter, r *http.Request)) {
	m.mux.HandleFunc(pattern, handler)
}

// Get registers a handler function for the pattern and only the GET method.
// Non-GET requests will be rejected with a 405 status code.
func (m *Mux) Get(pattern string, handler http.Handler) {
	p := fmt.Sprintf("GET %s", pattern)
	m.mux.Handle(p, handler)
}

// Post registers a handler function for the pattern and only the POST method.
// Non-POST requests will be rejected with a 405 status code.
func (m *Mux) Post(pattern string, handler http.Handler) {
	p := fmt.Sprintf("POST %s", pattern)
	m.mux.Handle(p, handler)
}

// Put registers a handler function for the pattern and only the PUT method.
// Non-PUT requests will be rejected with a 405 status code.
func (m *Mux) Put(pattern string, handler http.Handler) {
	p := fmt.Sprintf("PUT %s", pattern)
	m.mux.Handle(p, handler)
}

// Delete registers a handler function for the pattern and only the DELETE method.
// Non-DELETE requests will be rejected with a 405 status code.
func (m *Mux) Delete(pattern string, handler http.Handler) {
	p := fmt.Sprintf("DELETE %s", pattern)
	m.mux.Handle(p, handler)
}

// GetFunc registers a handler function for the pattern and only the GET method.
func (m *Mux) GetFunc(pattern string, handlerFunc func(w http.ResponseWriter, r *http.Request)) {
	m.Get(pattern, http.HandlerFunc(handlerFunc))
}

// PostFunc registers a handler function for the pattern and only the POST method.
func (m *Mux) PostFunc(pattern string, handlerFunc func(w http.ResponseWriter, r *http.Request)) {
	m.Post(pattern, http.HandlerFunc(handlerFunc))
}

// PutFunc registers a handler function for the pattern and only the PUT method.
func (m *Mux) PutFunc(pattern string, handlerFunc func(w http.ResponseWriter, r *http.Request)) {
	m.Put(pattern, http.HandlerFunc(handlerFunc))
}

// DeleteFunc registers a handler function for the pattern and only the DELETE method.
func (m *Mux) DeleteFunc(pattern string, handlerFunc func(w http.ResponseWriter, r *http.Request)) {
	m.Delete(pattern, http.HandlerFunc(handlerFunc))
}

// NotFoundHandler registers a handler for when no routes match the request.
// This is useful for returning a 404 page or similar.
// Alternative to using WithNotFoundHandler option on server creation.
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
	if p == "" && m.notFoundHandler != nil {
		m.notFoundHandler.ServeHTTP(w, req)
		return
	}

	h.ServeHTTP(w, req)
}
