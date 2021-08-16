package security

import (
	"errors"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gin-gonic/gin"
)

type Basic struct {
	Security
}

type User struct {
	Username string
	Password string
}

func (b *Basic) Authorize(c *gin.Context) {
	username, password, ok := c.Request.BasicAuth()
	if !ok {
		b.Callback(c, nil, errors.New("parse authentication error"))
	} else {
		b.Callback(c, &User{
			Username: username,
			Password: password,
		}, nil)
	}
}
func (b *Basic) Provider() string {
	return BasicAuth
}
func (b *Basic) Scheme() *openapi3.SecurityScheme {
	return &openapi3.SecurityScheme{
		Type:   "http",
		Scheme: "basic",
	}
}
