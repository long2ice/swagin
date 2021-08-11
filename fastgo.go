package fastgo

import (
	"github.com/gin-gonic/gin"
	"github.com/long2ice/fastgo/options"
	"github.com/long2ice/fastgo/swagger"
	"net/http"
)

type FastGo struct {
	*gin.Engine
	Swagger  *swagger.Swagger
	Handlers map[string]map[string][]gin.HandlerFunc
}

func Default(swagger *swagger.Swagger) *FastGo {
	return &FastGo{Engine: gin.Default(), Swagger: swagger, Handlers: make(map[string]map[string][]gin.HandlerFunc)}
}

func (g *FastGo) Group(path string) {

}
func (g *FastGo) handle(path string, method string, options ...options.Option) gin.IRoutes {
	key := path + method
	for _, option := range options {
		option(key, method, g)
	}
	if method == http.MethodGet {
		return g.Engine.GET(path, g.Handlers[key][method]...)
	} else if method == http.MethodPost {
		return g.Engine.POST(path, g.Handlers[key][method]...)
	} else if method == http.MethodHead {
		return g.Engine.HEAD(path, g.Handlers[key][method]...)
	} else if method == http.MethodPatch {
		return g.Engine.PATCH(path, g.Handlers[key][method]...)
	} else if method == http.MethodDelete {
		return g.Engine.DELETE(path, g.Handlers[key][method]...)
	} else if method == http.MethodPut {
		return g.Engine.PUT(path, g.Handlers[key][method]...)
	} else if method == http.MethodOptions {
		return g.Engine.OPTIONS(path, g.Handlers[key][method]...)
	} else {
		return g.Engine.Any(path, g.Handlers[key][method]...)
	}
}

func (g *FastGo) GET(path string, options ...options.Option) gin.IRoutes {
	return g.handle(path, http.MethodGet, options...)
}

func (g *FastGo) POST(path string, options ...options.Option) gin.IRoutes {
	return g.handle(path, http.MethodPost, options...)
}

func (g *FastGo) Run(addr ...string) error {
	return g.Engine.Run(addr...)
}
