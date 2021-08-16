package security

import (
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	Credentials = "credentials"
	BasicAuth   = "BasicAuth"
	BearerAuth  = "BearerAuth"
	ApiKeyAuth  = "ApiKeyAuth"
	OpenIDAuth  = "OpenIDAuth"
	OAuth2Auth  = "OAuth2Auth"
)

type ISecurity interface {
	Authorize(g *gin.Context)
	Callback(c *gin.Context, credentials interface{}, err error)
	Provider() string
	Scheme() *openapi3.SecurityScheme
}

type Security struct {
	ISecurity
}

func (s *Security) Callback(c *gin.Context, credentials interface{}, err error) {
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	} else {
		c.Set(Credentials, credentials)
	}
}
