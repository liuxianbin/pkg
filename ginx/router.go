package ginx

import (
	"github.com/gin-gonic/gin"
	"strings"
)

type Router struct {
	*gin.Engine
	excludePaths map[string][]string
}

func (r *Router) Exists(method string, path string) bool {
	if paths, ok := r.excludePaths[method]; ok {
		for _, p := range paths {
			if p == path {
				return true
			}
			if index := strings.Index(p, "/:"); index != -1 {
				if p[:index] == path[:index] {
					return true
				}
			}
		}
	}
	return false
}

func (r *Router) Exclude(method string, path string) {
	r.excludePaths[method] = append(r.excludePaths[method], path)
}

func (r *Router) POST(relativePath string, handlers ...gin.HandlerFunc) (string, string) {
	r.Engine.POST(relativePath, handlers...)
	return "POST", relativePath
}

func (r *Router) GET(relativePath string, handlers ...gin.HandlerFunc) (string, string) {
	r.Engine.GET(relativePath, handlers...)
	return "GET", relativePath
}

func (r *Router) DELETE(relativePath string, handlers ...gin.HandlerFunc) (string, string) {
	r.Engine.DELETE(relativePath, handlers...)
	return "DELETE", relativePath
}

func (r *Router) PUT(relativePath string, handlers ...gin.HandlerFunc) (string, string) {
	r.Engine.PUT(relativePath, handlers...)
	return "PUT", relativePath
}

func NewRouter() *Router {
	e := gin.New()
	r := &Router{
		Engine:       e,
		excludePaths: make(map[string][]string),
	}
	return r
}
