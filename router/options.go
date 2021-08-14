package router

import (
	"github.com/gin-gonic/gin"
)

type Option func(router *Router)

func API(api IAPI) Option {
	return func(router *Router) {
		router.API = api
		router.Handlers.PushBack(BindModel(api))
	}
}
func Middlewares(handlers ...gin.HandlerFunc) Option {
	return func(router *Router) {
		for _, handler := range handlers {
			router.Handlers.PushBack(handler)
		}
	}
}
func Tags(tags ...string) Option {
	return func(router *Router) {
		router.Tags = tags
	}
}
func Summary(summary string) Option {
	return func(router *Router) {
		router.Summary = summary
	}
}
func Description(description string) Option {
	return func(router *Router) {
		router.Description = description
	}
}
func Deprecated() Option {
	return func(router *Router) {
		router.Deprecated = true
	}
}
func OperationID(ID string) Option {
	return func(router *Router) {
		router.OperationID = ID
	}
}

func Exclude() Option {
	return func(router *Router) {
		router.Exclude = true
	}
}
