package gee

import(
	"fmt"
	"net/http"
)

const Get = "GET"
const Post = "POST"

// Request Handler
type HandlerFunc func(http.ResponseWriter, *http.Request)

// Route Table
type Engine struct{
	router map[string]HandlerFunc
}

// Constructor
func New() *Engine {

	return &Engine{router: make(map[string]HandlerFunc)}

}

func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {

	key := method + "-" + pattern
	engine.router[key] = handler

}

// add GET route
func (engine *Engine) GET(pattern string, handler HandlerFunc) {

	engine.addRoute(Get, pattern, handler)

}

// add POST route
func (engine *Engine) POST(pattern string, handler HandlerFunc) {

	engine.addRoute(Post, pattern, handler)

}

// Implementation of Handler interface
func (engine *Engine) ServeHTTP(rw http.ResponseWriter, req *http.Request) {

	key := req.Method + "-" + req.URL.Path
	if handler, ok := engine.router[key]; ok {
		handler(rw, req)
	} else {
		notFoundHandler(rw, req)
	}

}

// route that dose not exist
func notFoundHandler(rw http.ResponseWriter, req *http.Request) {
	rw.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(rw, "404: %q not found!\n", req.URL)
}

// add Run
func (engine *Engine) Run(port string) error {
	return http.ListenAndServe(port, engine)
}

