package main

import (
	"github.com/gin-gonic/gin"
	"github.com/long2ice/fastgo"
	"github.com/long2ice/fastgo/options"
	"github.com/long2ice/fastgo/swagger"
)

type Model struct {
	Name string `form:"name"`
}

func main() {
	app := fastgo.Default(swagger.Default("FastGo", "FastGo", "0.1.0"))
	app.GET("/ping", options.WithHandlers(func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	}))
	if err := app.Run(); err != nil {
		panic(err)
	}
}
