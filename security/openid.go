package security

import (
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gin-gonic/gin"
)

type OpenID struct {
	Security
	ConnectUrl string
}

func (i *OpenID) Authorize(c *gin.Context) {

}
func (i *OpenID) Provider() string {
	return OpenIDAuth
}

func (i *OpenID) Scheme() *openapi3.SecurityScheme {
	return &openapi3.SecurityScheme{
		Type:             "openIdConnect",
		OpenIdConnectUrl: i.ConnectUrl,
	}
}
