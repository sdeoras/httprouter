package httprouter

import (
	"net/http"
	"path/filepath"
	"strings"
	"sync"
)

// Router routes traffic to a registered http handle func.
type Router interface {
	// Route routes traffic to a registered http handle func.
	Route(w http.ResponseWriter, r *http.Request)
	// Register registers a http handle func against a path string.
	Register(path string, f func(w http.ResponseWriter, r *http.Request))
	// IsRegistered checks if a path is registered
	IsRegistered(path string) bool
}

// router implements Router
type router struct {
	mu       sync.Mutex
	registry map[string]func(w http.ResponseWriter, r *http.Request)
}

// NewRouter provides a new instance of Router
func NewRouter() Router {
	r := new(router)
	r.registry = make(map[string]func(w http.ResponseWriter, r *http.Request))
	return r
}

// Route routes traffic to a registered http handle func.
func (g *router) Route(w http.ResponseWriter, r *http.Request) {
	if f, ok := g.registry[filepath.Join("/", strings.ToLower(r.URL.Path))]; ok {
		f(w, r)
		return
	}

	http.Error(w, "404 page not found", http.StatusNotFound)
}

// Register registers a http handle func against a path string.
func (g *router) Register(path string, f func(w http.ResponseWriter, r *http.Request)) {
	g.mu.Lock()
	defer g.mu.Unlock()

	g.registry[filepath.Join("/", strings.ToLower(path))] = f
}

func (g *router) IsRegistered(path string) bool {
	_, ok := g.registry[filepath.Join("/", strings.ToLower(path))]
	return ok
}
