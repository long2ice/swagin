package router

import (
	"github.com/gin-gonic/gin"
	"github.com/long2ice/fastgo/middlewares"
)

type Option func(router *Router)

func Handlers(handlers ...gin.HandlerFunc) Option {
	return func(router *Router) {
		for _, handler := range handlers {
			router.Handlers.PushBack(handler)
		}
	}
}

func Summary(summary string) Option {
	return func(router *Router) {
		router.Summary = summary
	}
}
func Deprecated() Option {
	return func(router *Router) {
		router.Deprecated = true
	}
}
func Model(model interface{}) Option {
	return func(router *Router) {
		router.Handlers.PushFront(middlewares.Validate(model))
	}
}
