package gorouter

import "net/http"

type Method string

const (
	Get    Method = "GET"
	Post   Method = "POST"
	Put    Method = "PUT"
	Delete Method = "DELETE"
)

type Route struct {
	method  Method
	path    string
	handler http.HandlerFunc
}

type Router struct {
	routes []*Route
}

func NewRouter() *Router {
	return new(Router)
}

func (h *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fn := h.findHandler(r)
	if fn == nil {
		http.NotFound(w, r)
		return
	}
	fn(w, r)
}

func (h *Router) findHandler(r *http.Request) http.HandlerFunc {
	for _, v := range h.routes {
		if v.path == r.URL.Path && string(v.method) == r.Method {
			return v.handler
		}
	}
	return nil
}

func (h *Router) Add(method Method, path string, fn http.HandlerFunc) {
	route := new(Route)
	route.path = path
	route.method = method
	route.handler = fn
	h.routes = append(h.routes, route)
}

func (h *Router) Get(path string, fn http.HandlerFunc) {
	h.Add(Get, path, fn)
}

func (h *Router) Post(path string, fn http.HandlerFunc) {
	h.Add(Post, path, fn)
}

func (h *Router) Put(path string, fn http.HandlerFunc) {
	h.Add(Put, path, fn)
}

func (h *Router) Delete(path string, fn http.HandlerFunc) {
	h.Add(Delete, path, fn)
}
