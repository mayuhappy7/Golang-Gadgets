package gee

import (
	"fmt"
	"net/http"
)

// HandlerFunc is a type that defines a function to handle HTTP requests
type HandlerFunc func(w http.ResponseWriter, req *http.Request)

// Engine is the core structure of the Gee framework, which implements http.Handler interface
type Engine struct {
	router map[string]HandlerFunc // key: HTTP method + "-" + path
}

// New creates and initializes a new Engine instance
func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

// addRoute registers a new route with method, pattern and handler function
func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	engine.router[key] = handler
}

// GET registers a new GET route for a path pattern with handler function
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

// POST registers a new POST route for a path pattern with handler function
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

// Run starts the HTTP server at specified address
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

// ServeHTTP handles all HTTP requests and matches routes based on method and path
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	if handler, ok := engine.router[key]; ok {
		handler(w, req)
	} else {
		fmt.Fprintf(w, "404 NOT FOUND: %q\n", req.URL)
	}
}
