package webserver

import (
	"net/http"
)

type Router struct {
	rules map[string]map[string]http.HandlerFunc
}

// ServeHTTP implements http.Handler
func (r *Router) ServeHTTP(wt http.ResponseWriter, req *http.Request) {
	handler, methodExist, exist := r.FindHandler(req.Method, req.URL.Path)

	if !methodExist {
		wt.WriteHeader(http.StatusBadRequest)
		return
	}

	if !exist {
		wt.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	handler(wt, req)
}

func (r *Router) FindHandler(method string, path string) (http.HandlerFunc, bool, bool) {
	methodPaths, methodExist := r.rules[method]
	handler, exist := methodPaths[path]
	return handler, methodExist, exist

}

func NewRouter() *Router {
	return &Router{
		rules: make(map[string]map[string]http.HandlerFunc),
	}
}
