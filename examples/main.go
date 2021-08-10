package main

import (
	"github.com/gin-gonic/gin"
	"github.com/long2ice/fastgo"
)

type Model struct {
	Name string `form:"name"`
}

func main() {
	r := gin.Default()
	r.GET("/ping", fastgo.Validate(&Model{}), func(c *gin.Context) {
		model := c.MustGet(fastgo.ParsedModel).(*Model)
		c.JSON(200, model)
	})
	r.Run()
}
