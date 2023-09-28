package router

import (
	"net/http"
)

type Middleware func(http.Handler) http.Handler

type Mux struct {
	middlewares []Middleware
	mux         *http.ServeMux
}

func NewMux() *Mux {
	return &Mux{
		mux: http.NewServeMux(),
	}
}

func (m *Mux) Handle(pattern string, handler http.Handler) {
	m.mux.Handle(pattern, handler)
}

func (m *Mux) AddMiddleware(middleware Middleware) {
	m.middlewares = append(m.middlewares, middleware)
}

// Get registers a handler function for the pattern and only the GET method.
// Non-GET requests will be rejected with a 405 status code.
func (m *Mux) Get(pattern string, handler http.HandlerFunc) {
	m.mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		handler(w, r)
	})
}

// Post registers a handler function for the pattern and only the POST method.
// Non-POST requests will be rejected with a 405 status code.
func (m *Mux) Post(pattern string, handler http.HandlerFunc) {
	m.mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		handler(w, r)
	})
}

// Put registers a handler function for the pattern and only the PUT method.
// Non-PUT requests will be rejected with a 405 status code.
func (m *Mux) Put(pattern string, handler http.HandlerFunc) {
	m.mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		handler(w, r)
	})
}

// Delete registers a handler function for the pattern and only the DELETE method.
// Non-DELETE requests will be rejected with a 405 status code.
func (m *Mux) Delete(pattern string, handler http.HandlerFunc) {
	m.mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		handler(w, r)
	})
}

func (m *Mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// for _, middleware := range m.middlewares {

	// }
	m.mux.ServeHTTP(w, r)
}
