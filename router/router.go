package router

import (
	"container/list"
	"github.com/gin-gonic/gin"
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
}

func BindModel(api IAPI) gin.HandlerFunc {
	return func(c *gin.Context) {
		model := api.NewModel()
		params := make(map[string][]string)
		for _, v := range c.Params {
			params[v.Key] = []string{v.Value}
		}

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
	for h := router.Handlers.Front(); h != nil; h = h.Next() {
		handlers = append(handlers, h.Value.(gin.HandlerFunc))
	}
	handlers = append(handlers, router.API.Handler)
	return handlers
}
func New(options ...Option) *Router {
	r := &Router{
		Handlers: list.New(),
	}
	for _, option := range options {
		option(r)
	}
	return r
}
