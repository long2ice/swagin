package security

import (
	"errors"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gin-gonic/gin"
)

type ApiKey struct {
	Security
	Name string
}

func (k *ApiKey) Authorize(c *gin.Context) {
	auth := c.Request.Header.Get(k.Name)
	if auth == "" {
		k.Callback(c, nil, errors.New("empty apikey"))
	} else {
		k.Callback(c, auth, nil)
	}
}
func (k *ApiKey) Provider() string {
	return ApiKeyAuth
}

func (k *ApiKey) Scheme() *openapi3.SecurityScheme {
	return &openapi3.SecurityScheme{
		Type: "http",
		In:   "header",
		Name: k.Name,
	}
}
