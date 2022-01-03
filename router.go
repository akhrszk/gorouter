package gorouter

import (
	"net/http"
	"strings"
)

type Router struct {
	route *Node
}

type Params map[string]string

type Handler func(http.ResponseWriter, *http.Request, Params)

func New() *Router {
	return &Router{
		route: newRootNode(),
	}
}

func (rt *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p := strings.Split(r.URL.Path, "/")
	p = filter(p)
	node, params := rt.route.find(p, Params{})
	if node == nil {
		http.NotFound(w, r)
		return
	}
	fn := node.handlers[r.Method]
	if fn == nil {
		http.NotFound(w, r)
		return
	}
	fn(w, r, params)
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

func (rt *Router) Handle(method string, path string, fn Handler) {
	p := strings.Split(path, "/")
	p = filter(p)
	rt.route.add(p, method, fn)
}

func (rt *Router) Get(path string, fn Handler) {
	rt.Handle(http.MethodGet, path, fn)
}

func (rt *Router) Head(path string, fn Handler) {
	rt.Handle(http.MethodHead, path, fn)
}

func (rt *Router) Post(path string, fn Handler) {
	rt.Handle(http.MethodPost, path, fn)
}

func (rt *Router) Put(path string, fn Handler) {
	rt.Handle(http.MethodPut, path, fn)
}

func (rt *Router) Patch(path string, fn Handler) {
	rt.Handle(http.MethodPatch, path, fn)
}

func (rt *Router) Delete(path string, fn Handler) {
	rt.Handle(http.MethodDelete, path, fn)
}

func (rt *Router) Connect(path string, fn Handler) {
	rt.Handle(http.MethodConnect, path, fn)
}

func (rt *Router) Options(path string, fn Handler) {
	rt.Handle(http.MethodOptions, path, fn)
}

func (rt *Router) Trace(path string, fn Handler) {
	rt.Handle(http.MethodTrace, path, fn)
}
