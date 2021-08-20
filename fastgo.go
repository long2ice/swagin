package fastgo

import (
	"embed"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/long2ice/fastgo/router"
	"github.com/long2ice/fastgo/swagger"
	"html/template"
	"net/http"
)

//go:embed templates/*
var templates embed.FS

type FastGo struct {
	*gin.Engine
	Swagger *swagger.Swagger
	Routers map[string]map[string]*router.Router
}

func New(swagger *swagger.Swagger) *FastGo {
	f := &FastGo{Engine: gin.Default(), Swagger: swagger, Routers: make(map[string]map[string]*router.Router)}
	f.SetHTMLTemplate(template.Must(template.ParseFS(templates, "templates/*.html")))
	swagger.Routers = f.Routers
	return f
}

func (g *FastGo) Group(path string, options ...Option) *Group {
	group := &Group{
		FastGo: g,
		Path:   path,
	}
	for _, option := range options {
		option(group)
	}
	return group
}

func (g *FastGo) Handle(path string, method string, r *router.Router) gin.IRoutes {
	r.Method = method
	r.Path = path
	if g.Routers[path] == nil {
		g.Routers[path] = make(map[string]*router.Router)
	}
	g.Routers[path][method] = r
	handlers := r.GetHandlers()
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
	return g.Handle(path, http.MethodGet, router)
}

func (g *FastGo) POST(path string, router *router.Router) gin.IRoutes {
	return g.Handle(path, http.MethodPost, router)
}

func (g *FastGo) HEAD(path string, router *router.Router) gin.IRoutes {
	return g.Handle(path, http.MethodHead, router)
}

func (g *FastGo) PATCH(path string, router *router.Router) gin.IRoutes {
	return g.Handle(path, http.MethodPatch, router)
}

func (g *FastGo) DELETE(path string, router *router.Router) gin.IRoutes {
	return g.Handle(path, http.MethodDelete, router)
}

func (g *FastGo) PUT(path string, router *router.Router) gin.IRoutes {
	return g.Handle(path, http.MethodPut, router)
}

func (g *FastGo) OPTIONS(path string, router *router.Router) gin.IRoutes {
	return g.Handle(path, http.MethodOptions, router)
}

func (g *FastGo) init() {
	g.Engine.GET(g.Swagger.OpenAPIUrl, func(c *gin.Context) {
		c.JSON(http.StatusOK, g.Swagger)
	})
	g.Engine.GET(g.Swagger.DocsUrl, func(c *gin.Context) {
		options := `{}`
		if g.Swagger.SwaggerOptions != nil {
			data, err := json.Marshal(g.Swagger.SwaggerOptions)
			if err != nil {
				panic(err)
			}
			options = string(data)
		}
		c.HTML(http.StatusOK, "swagger.html", gin.H{
			"openapi_url":     g.Swagger.OpenAPIUrl,
			"title":           g.Swagger.Title,
			"swagger_options": options,
		})
	})
	g.Engine.GET(g.Swagger.RedocUrl, func(c *gin.Context) {
		options := `{}`
		if g.Swagger.RedocOptions != nil {
			data, err := json.Marshal(g.Swagger.RedocOptions)
			if err != nil {
				panic(err)
			}
			options = string(data)
		}
		c.HTML(http.StatusOK, "redoc.html", gin.H{
			"openapi_url":   g.Swagger.OpenAPIUrl,
			"title":         g.Swagger.Title,
			"redoc_options": options,
		})
	})
	g.Swagger.BuildOpenAPI()
}

func (g *FastGo) Run(addr ...string) error {
	g.init()
	return g.Engine.Run(addr...)
}
