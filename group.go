package fastgo

import (
	"github.com/gin-gonic/gin"
	"github.com/long2ice/fastgo/router"
	"github.com/long2ice/fastgo/security"
	"net/http"
)

type Group struct {
	*FastGo
	Path       string
	Tags       []string
	Handlers   []gin.HandlerFunc
	Securities []security.ISecurity
}
type Option func(*Group)

func Handlers(handlers ...gin.HandlerFunc) Option {
	return func(g *Group) {
		for _, handler := range handlers {
			g.Handlers = append(g.Handlers, handler)
		}
	}
}

func Tags(tags ...string) Option {
	return func(g *Group) {
		g.Tags = tags
	}
}

func Security(securities ...security.ISecurity) Option {
	return func(g *Group) {
		for _, s := range securities {
			g.Securities = append(g.Securities, s)
		}
	}
}

func (g *Group) Handle(path string, method string, r *router.Router) gin.IRoutes {
	router.Handlers(g.Handlers...)(r)
	router.Tags(g.Tags...)(r)
	router.Security(g.Securities...)(r)
	return g.FastGo.Handle(g.Path+path, method, r)
}
func (g *Group) GET(path string, router *router.Router) gin.IRoutes {
	return g.Handle(path, http.MethodGet, router)
}

func (g *Group) POST(path string, router *router.Router) gin.IRoutes {
	return g.Handle(path, http.MethodPost, router)
}

func (g *Group) HEAD(path string, router *router.Router) gin.IRoutes {
	return g.Handle(path, http.MethodHead, router)
}

func (g *Group) PATCH(path string, router *router.Router) gin.IRoutes {
	return g.Handle(path, http.MethodPatch, router)
}

func (g *Group) DELETE(path string, router *router.Router) gin.IRoutes {
	return g.Handle(path, http.MethodDelete, router)
}

func (g *Group) PUT(path string, router *router.Router) gin.IRoutes {
	return g.Handle(path, http.MethodPut, router)
}

func (g *Group) OPTIONS(path string, router *router.Router) gin.IRoutes {
	return g.Handle(path, http.MethodOptions, router)
}

func (g *Group) Group(path string, options ...Option) *Group {
	group := &Group{
		FastGo:     g.FastGo,
		Path:       g.Path + path,
		Tags:       g.Tags,
		Handlers:   g.Handlers,
		Securities: g.Securities,
	}
	for _, option := range options {
		option(group)
	}
	return group
}