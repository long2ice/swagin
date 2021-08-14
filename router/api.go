package router

import "github.com/gin-gonic/gin"

type IAPI interface {
	NewModel() interface{}
	Handler(context *gin.Context)
}
