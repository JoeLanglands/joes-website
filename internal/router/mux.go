package router

import (
	"log/slog"
	"net/http"
)

// TODO Remove the logging after you've played around with it, there is no need here

var logger *slog.Logger

type Middleware func(http.Handler) http.Handler

type Mux struct {
	middlewares []Middleware
	mux         *http.ServeMux
}

func NewMux(l *slog.Logger) *Mux {
	logger = l
	return &Mux{
		mux: http.NewServeMux(),
	}
}

func (m *Mux) Handle(pattern string, handler http.Handler) {
	m.mux.Handle(pattern, handler)
}

func (m *Mux) Get(pattern string, handler http.HandlerFunc) {
	m.mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			logger.Info("request HTTP method not allowed", "method", r.Method, "path", r.URL.Path)
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		logger.Info("request", "method", r.Method, "path", r.URL.Path)
		handler(w, r)
	})
}

func (m *Mux) Post(pattern string, handler http.HandlerFunc) {
	m.mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			logger.Info("request HTTP method not allowed", "method", r.Method, "path", r.URL.Path)
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		logger.Info("request", "method", r.Method, "path", r.URL.Path)
		handler(w, r)
	})
}

func (m *Mux) Put(pattern string, handler http.HandlerFunc) {
	m.mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			logger.Info("request HTTP method not allowed", "method", r.Method, "path", r.URL.Path)
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		logger.Info("request", "method", r.Method, "path", r.URL.Path)
		handler(w, r)
	})
}

func (m *Mux) Delete(pattern string, handler http.HandlerFunc) {
	m.mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			logger.Info("request HTTP method not allowed", "method", r.Method, "path", r.URL.Path)
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		logger.Info("request", "method", r.Method, "path", r.URL.Path)
		handler(w, r)
	})
}

func (m *Mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// for _, middleware := range m.middlewares {

	// }
	m.mux.ServeHTTP(w, r)
}
