package router

import (
	"container/list"
	"github.com/gin-gonic/gin"
)

type Router struct {
	Handlers    *list.List
	Path        string
	Method      string
	Summary     string
	Description string
	Deprecated  bool
	Tags        []string
	Model       interface{}
}

func (router *Router) GetHandlers() []gin.HandlerFunc {
	var handlers []gin.HandlerFunc
	for h := router.Handlers.Front(); h != nil; h = h.Next() {
		handlers = append(handlers, h.Value.(gin.HandlerFunc))
	}
	return handlers
}
func Default(options ...Option) *Router {
	r := &Router{
		Handlers: list.New(),
	}
	for _, option := range options {
		option(r)
	}
	return r
}
