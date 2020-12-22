package sariaf

import "net/http"

// Context is the most important part of sariaf. It allows us to pass variables between middleware,
// manage the flow, validate the JSON of a request and render a JSON response.
type Context struct {
	Request *http.Request
	Writer  http.ResponseWriter
}

// A Handler responds to an HTTP request.
type Handler interface {
	ServeHTTP(*Context)
}

// HandlerFunc defines the handler used by gin middleware as return value.
type HandlerFunc func(*Context)

// ServeHTTP calls f(w, r).
func (f HandlerFunc) ServeHTTP(c *Context) {
	f(c)
}

// HTTPHandlerFunc is an alias for http.HandlerFunc.
type HTTPHandlerFunc http.HandlerFunc

// ServeHTTP calls f(w, r).
func (f HTTPHandlerFunc) ServeHTTP(c *Context) {
	f(c.Writer, c.Request)
}
