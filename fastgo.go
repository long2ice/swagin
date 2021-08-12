package fastgo

import (
	"github.com/gin-gonic/gin"
	"github.com/long2ice/fastgo/router"
	"github.com/long2ice/fastgo/swagger"
	"net/http"
)

type FastGo struct {
	*gin.Engine
	Swagger *swagger.Swagger
	Routers map[string]*router.Router
}

func Default(swagger *swagger.Swagger) *FastGo {
	return &FastGo{Engine: gin.Default(), Swagger: swagger, Routers: make(map[string]*router.Router)}
}

func (g *FastGo) Group(path string, handlers ...gin.HandlerFunc) *gin.RouterGroup {
	return g.Engine.Group(path, handlers...)
}
func (g *FastGo) handle(path string, method string, router *router.Router) gin.IRoutes {
	key := path + method
	g.Routers[key] = router
	handlers := router.GetHandlers()
	if method == http.MethodGet {
		return g.Engine.GET(path, handlers...)
	} else if method == http.MethodPost {
		return g.Engine.POST(path, handlers...)
	} else if method == http.MethodHead {
		return g.Engine.HEAD(path, handlers...)
	} else if method == http.MethodPatch {
		return g.Engine.PATCH(path, handlers...)
	} else if method == http.MethodDelete {
		return g.Engine.DELETE(path, handlers...)
	} else if method == http.MethodPut {
		return g.Engine.PUT(path, handlers...)
	} else if method == http.MethodOptions {
		return g.Engine.OPTIONS(path, handlers...)
	} else {
		return g.Engine.Any(path, handlers...)
	}
}

func (g *FastGo) GET(path string, router *router.Router) gin.IRoutes {
	return g.handle(path, http.MethodGet, router)
}

func (g *FastGo) POST(path string, router *router.Router) gin.IRoutes {
	return g.handle(path, http.MethodPost, router)
}
func (g *FastGo) HEAD(path string, router *router.Router) gin.IRoutes {
	return g.handle(path, http.MethodHead, router)
}
func (g *FastGo) PATCH(path string, router *router.Router) gin.IRoutes {
	return g.handle(path, http.MethodPatch, router)
}
func (g *FastGo) DELETE(path string, router *router.Router) gin.IRoutes {
	return g.handle(path, http.MethodDelete, router)
}
func (g *FastGo) PUT(path string, router *router.Router) gin.IRoutes {
	return g.handle(path, http.MethodPut, router)
}
func (g *FastGo) OPTIONS(path string, router *router.Router) gin.IRoutes {
	return g.handle(path, http.MethodOptions, router)
}

func (g *FastGo) Run(addr ...string) error {
	return g.Engine.Run(addr...)
}
