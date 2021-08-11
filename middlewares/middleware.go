package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/long2ice/fastgo/constants"
	"net/http"
)

func Validate(model interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := c.ShouldBind(model); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.Set(constants.ParsedModel, model)
		c.Next()
	}
}
