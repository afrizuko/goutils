package router

import (
	"net/http"
	"net/url"
	"strings"
)

type key string

type Handle func(http.ResponseWriter, *http.Request, url.Values)

type Router struct {
	tree *node
	key  key
}

func NewRouter() *Router {
	return &Router{key: "params",
		tree: &node{component: "/", isNamedParam: false, methods: make(map[string]Handle)},
	}
}

func (rt *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	params := req.Form
	node, _ := rt.tree.traverse(strings.Split(req.URL.Path, "/")[1:], params)
	if handler := node.methods[req.Method]; handler != nil {
		// ctx := context.WithValue(req.Context(), rt.key, params)
		// handler(w, req.WithContext(ctx), params)
		handler(w, req, params)
		return
	}
	http.NotFound(w, req)
}

func GetParam(req *http.Request, field string) string {
	return req.Context().Value(field).(string)
}

func (rt *Router) Handle(method, path string, handler Handle) {
	if path[0] != '/' {
		panic("Path should to start with a /.")
	}
	rt.tree.addNode(method, path, handler)
}

func (rt *Router) GET(path string, h Handle) {
	rt.Handle("GET", path, h)
}

func (rt *Router) PUT(path string, h Handle) {
	rt.Handle("PUT", path, h)
}

func (rt *Router) POST(path string, h Handle) {
	rt.Handle("POST", path, h)
}

func (rt *Router) DELETE(path string, h Handle) {
	rt.Handle("DELETE", path, h)
}

func (rt *Router) OPTIONS(path string, h Handle) {
	rt.Handle("OPTIONS", path, h)
}
