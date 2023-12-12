package ginx

import (
	"github.com/gin-gonic/gin"
	"strings"
)

type Router struct {
	*gin.Engine
	excludeToken      map[string][]string // 无需token登录
	excludePermission map[string][]string // 无需权限
}

const (
	ANONYMOUS = iota
	ALLOW
)

func (r *Router) Exists(retype int, method string, path string) bool {
	var excludes map[string][]string
	switch retype {
	case ALLOW:
		excludes = r.excludePermission
	case ANONYMOUS:
		excludes = r.excludeToken
	}
	if paths, ok := excludes[method]; ok {
		for _, p := range paths {
			if p == path {
				return true
			}
			if strings.Count(p, "/") == strings.Count(path, "/") {
				if index := strings.Index(p, "/:"); index != -1 {
					if strings.HasPrefix(path, p[:index]) {
						return true
					}
				}
			}
		}
	}
	return false
}

func (r *Router) Anonymous(method string, path string) {
	r.excludeToken[method] = append(r.excludeToken[method], path)
}

func (r *Router) Allow(method string, path string) {
	r.excludePermission[method] = append(r.excludePermission[method], path)
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
		Engine:            e,
		excludeToken:      make(map[string][]string),
		excludePermission: make(map[string][]string),
	}
	return r
}
