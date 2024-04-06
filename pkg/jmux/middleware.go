package jmux

import (
	"net/http"
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
