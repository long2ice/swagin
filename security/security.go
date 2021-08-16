package security

import "github.com/gin-gonic/gin"

type Security interface {
	Authorize(g *gin.Context)
	Callback(g *gin.Context, credentials interface{}, err error)
}
