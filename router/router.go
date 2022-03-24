package router

import (
	"container/list"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/long2ice/swagin/security"
	"github.com/mcuadros/go-defaults"
)

type IAPI interface {
	Handler(context *gin.Context)
}

type ErrorHandlerFunc func(ctx *gin.Context, err error, status int)

type Router struct {
	Handlers            *list.List
	Path                string
	Method              string
	Summary             string
	Description         string
	Deprecated          bool
	RequestContentType  string
	ResponseContentType string
	Tags                []string
	API                 IAPI
	OperationID         string
	Exclude             bool
	Securities          []security.ISecurity
	Response            Response
	ErrorHandler        ErrorHandlerFunc
}

var validate = validator.New()

func errHandle(ctx *gin.Context, err error, f ErrorHandlerFunc) {
	if f != nil {
		f(ctx, err, http.StatusBadRequest)
	} else {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}
}
func BindModel(api IAPI, f ErrorHandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		model := reflect.New(reflect.TypeOf(api).Elem()).Interface()
		if err := c.ShouldBindHeader(model); err != nil {
			errHandle(c, err, f)
			return
		}
		if err := CookiesParser(c, model); err != nil {
			errHandle(c, err, f)
			return
		}
		if err := c.ShouldBindWith(model, Query); err != nil {
			errHandle(c, err, f)
			return
		}
		if c.Request.Method == http.MethodPost || c.Request.Method == http.MethodPut {
			switch c.Request.Header.Get("Content-Type") {
			case binding.MIMEMultipartPOSTForm:
				if err := c.ShouldBindWith(model, binding.FormMultipart); err != nil {
					errHandle(c, err, f)
					return
				}
			case binding.MIMEJSON:
				if err := c.ShouldBindWith(model, binding.JSON); err != nil {
					errHandle(c, err, f)
					return
				}
			case binding.MIMEXML:
				if err := c.ShouldBindWith(model, binding.XML); err != nil {
					errHandle(c, err, f)
					return
				}
			case binding.MIMEPOSTForm:
				if err := c.ShouldBindWith(model, binding.Form); err != nil {
					errHandle(c, err, f)
					return
				}
			case binding.MIMEYAML:
				if err := c.ShouldBindWith(model, binding.YAML); err != nil {
					errHandle(c, err, f)
					return
				}
			case binding.MIMEPROTOBUF:
				if err := c.ShouldBindWith(model, binding.ProtoBuf); err != nil {
					errHandle(c, err, f)
					return
				}
			case binding.MIMEMSGPACK:
				if err := c.ShouldBindWith(model, binding.MsgPack); err != nil {
					errHandle(c, err, f)
					return
				}
			}
		}
		if err := c.ShouldBindUri(model); err != nil {
			errHandle(c, err, f)
			return
		}
		defaults.SetDefaults(model)
		if err := validate.Struct(model); err != nil {
			errHandle(c, err, f)
			return
		}
		if err := copier.Copy(api, model); err != nil {
			errHandle(c, err, f)
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
	handlers = append(handlers, BindModel(router.API, router.ErrorHandler))
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
	for _, option := range options {
		option(r)
	}
	return r
}
func (router *Router) WithSecurity(securities ...security.ISecurity) *Router {
	Security(securities...)(router)
	return router
}
func (router *Router) WithResponses(response Response) *Router {
	Responses(response)(router)
	return router
}
func (router *Router) WithHandlers(handlers ...gin.HandlerFunc) *Router {
	Handlers(handlers...)(router)
	return router
}
func (router *Router) WithTags(tags ...string) *Router {
	Tags(tags...)(router)
	return router
}
func (router *Router) WithSummary(summary string) *Router {
	Summary(summary)(router)
	return router
}
func (router *Router) WithErrorHandler(f ErrorHandlerFunc) *Router {
	ErrorHandler(f)(router)
	return router
}
func (router *Router) WithDescription(description string) *Router {
	Description(description)(router)
	return router
}
func (router *Router) WithDeprecated() *Router {
	Deprecated()(router)
	return router
}
func (router *Router) WithOperationID(ID string) *Router {
	OperationID(ID)(router)
	return router
}
func (router *Router) WithExclude() *Router {
	Exclude()(router)
	return router
}
func (router *Router) WithContentType(contentType string, contentTypeType ContentTypeType) *Router {
	ContentType(contentType, contentTypeType)(router)
	return router
}
