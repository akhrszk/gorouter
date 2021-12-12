package gorouter

import (
	"net/http"
	"strings"
)

type Router struct {
	route *Node
}

func NewRouter() *Router {
	return &Router{
		route: newRootNode(),
	}
}

func (rt *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p := strings.Split(r.URL.Path, "/")
	p = filter(p)
	node := rt.route.find(p)
	if node == nil {
		http.NotFound(w, r)
		return
	}
	fn := node.handlers[r.Method]
	if fn == nil {
		http.NotFound(w, r)
		return
	}
	fn(w, r)
}

// remove empty string
func filter(arr []string) []string {
	b := arr[:0]
	for _, x := range arr {
		if len(x) > 0 {
			b = append(b, x)
		}
	}
	return b
}

func (rt *Router) Add(method string, path string, fn http.HandlerFunc) {
	p := strings.Split(path, "/")
	p = filter(p)
	rt.route.add(p, method, fn)
}

func (rt *Router) Get(path string, fn http.HandlerFunc) {
	rt.Add(http.MethodGet, path, fn)
}

func (rt *Router) Post(path string, fn http.HandlerFunc) {
	rt.Add(http.MethodPost, path, fn)
}

func (rt *Router) Put(path string, fn http.HandlerFunc) {
	rt.Add(http.MethodPut, path, fn)
}

func (rt *Router) Patch(path string, fn http.HandlerFunc) {
	rt.Add(http.MethodPatch, path, fn)
}

func (rt *Router) Delete(path string, fn http.HandlerFunc) {
	rt.Add(http.MethodDelete, path, fn)
}
