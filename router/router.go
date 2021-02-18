package router

import (
	"context"
	"net/http"
	"net/url"
	"strings"
)

type Key string

type Handle func(http.ResponseWriter, *http.Request)

const (
	key = Key("params")
)

type Router struct {
	tree *node
}

func NewRouter() *Router {
	return &Router{
		tree: &node{component: "/", isNamedParam: false, methods: make(map[string]Handle)},
	}
}

func (rt *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	params := req.Form
	node, _ := rt.tree.traverse(strings.Split(req.URL.Path, "/")[1:], params)
	if handler := node.methods[req.Method]; handler != nil {
		ctx := context.WithValue(req.Context(), key, params)
		handler(w, req.WithContext(ctx))
		return
	}
	http.NotFound(w, req)
}

func GetParam(req *http.Request, field string) string {
	vals := req.Context().Value(key).(url.Values)
	if vals == nil {
		return ""
	}
	if len(vals) < 1 {
		return ""
	}
	return vals[field][0]
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
