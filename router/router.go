package router

import (
	"container/list"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/long2ice/swagin/security"
	"net/http"
	"reflect"
)

type IAPI interface {
	Handler(context *gin.Context)
}
type Router struct {
	Handlers    *list.List
	Path        string
	Method      string
	Summary     string
	Description string
	Deprecated  bool
	ContentType string
	Tags        []string
	API         IAPI
	OperationID string
	Exclude     bool
	Securities  []security.ISecurity
	Response    Response
}

func BindModel(api IAPI) gin.HandlerFunc {
	return func(c *gin.Context) {
		model := reflect.New(reflect.TypeOf(api).Elem()).Interface()
		if err := c.ShouldBindRequest(model); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err := copier.Copy(api, model)
		if err != nil {
			return
		}
		c.Next()
	}
}

func (router *Router) GetHandlers() []gin.HandlerFunc {
	var handlers []gin.HandlerFunc
	for _, s := range router.Securities {
		handlers = append(handlers, s.Authorize)
	}
	for h := router.Handlers.Front(); h != nil; h = h.Next() {
		if f, ok := h.Value.(gin.HandlerFunc); ok {
			handlers = append(handlers, f)
		}
	}
	handlers = append(handlers, router.API.Handler)
	return handlers
}

func New(api IAPI, options ...Option) *Router {
	r := &Router{
		Handlers: list.New(),
		API:      api,
		Response: make(Response),
	}
	r.Handlers.PushBack(BindModel(api))
	for _, option := range options {
		option(r)
	}
	return r
}
