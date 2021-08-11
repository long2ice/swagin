package options

import (
	"github.com/gin-gonic/gin"
	"github.com/long2ice/fastgo"
)

type Option func(string, string, *fastgo.FastGo)

func WithHandlers(handlers ...gin.HandlerFunc) Option {
	return func(path string, method string, g *fastgo.FastGo) {
		g.Handlers[path][method] = handlers
	}
}

func WithSummary(summary string) Option {
	return func(path string, method string, g *fastgo.FastGo) {

	}
}
