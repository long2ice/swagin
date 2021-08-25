package router

import (
	"container/list"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gin-gonic/gin"
	"github.com/long2ice/fastgo/security"
	"net/http"
	"reflect"
)

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
	Responses   openapi3.Responses
}

func BindModel(api IAPI) gin.HandlerFunc {
	return func(c *gin.Context) {
		model := api.NewModel()
		if err := c.ShouldBindRequest(model); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		getType := reflect.TypeOf(api).Elem()
		getValue := reflect.ValueOf(api).Elem()
		for i := 0; i < getType.NumField(); i++ {
			field := getType.Field(i)
			value := getValue.Field(i)
			if field.Type == reflect.TypeOf(model) {
				value.Set(reflect.ValueOf(model))
				break
			}
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
		Handlers:  list.New(),
		API:       api,
		Responses: openapi3.NewResponses(),
	}
	r.Handlers.PushBack(BindModel(api))
	for _, option := range options {
		option(r)
	}
	return r
}
