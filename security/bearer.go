package security

import (
	"errors"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gin-gonic/gin"
	"strings"
)

type Bearer struct {
	Security
}

func (b *Bearer) Authorize(c *gin.Context) {
	auth := c.Request.Header.Get("Authorization")
	if auth == "" {
		b.Callback(c, nil, errors.New("empty authentication"))
	} else {
		splits := strings.Split(auth, "Bearer ")
		if len(splits) != 2 {
			b.Callback(c, nil, errors.New("invalid authentication string"))
		} else {
			b.Callback(c, splits[1], nil)
		}
	}
}
func (b *Bearer) Provider() string {
	return BearerAuth
}

func (b *Bearer) Scheme() *openapi3.SecurityScheme {
	return &openapi3.SecurityScheme{
		Type:         "http",
		Scheme:       "bearer",
		BearerFormat: "JWT",
	}
}
