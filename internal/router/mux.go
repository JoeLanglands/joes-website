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

// AddMiddleware adds a middleware to the mux. These middleware will be
// executed in the order they are added and are called for every request served.
// Global middleware are be called before route specific middleware.
func (m *Mux) AddGlobalMiddleware(middleware Middleware) {
	m.middlewares = append(m.middlewares, middleware)
}

// Use adds a middleware to the Handler chain so a set of middleware can be
// applied to a single route.
func Use(mw Middleware, handler http.Handler) http.Handler {
	return mw(handler)
}

// Get registers a handler function for the pattern and only the GET method.
// Non-GET requests will be rejected with a 405 status code.
func (m *Mux) Get(pattern string, handler http.Handler) {
	m.mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		m.applyGlobalMiddlewares(handler).ServeHTTP(w, r)
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
		m.applyGlobalMiddlewares(handler).ServeHTTP(w, r)
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
		m.applyGlobalMiddlewares(handler).ServeHTTP(w, r)
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
		m.applyGlobalMiddlewares(handler).ServeHTTP(w, r)
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

func (m *Mux) applyGlobalMiddlewares(handler http.Handler) http.Handler {
	for _, m := range m.middlewares {
		handler = m(handler)
	}
	return handler
}

func (m *Mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m.mux.ServeHTTP(w, r)
}
