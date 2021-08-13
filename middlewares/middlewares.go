package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/long2ice/fastgo/constants"
	"net/http"
)

func BindModel(model interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := c.ShouldBindRequest(model); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.Set(constants.Model, model)
		c.Next()
	}
}
