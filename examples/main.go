package main

import (
	"github.com/gin-gonic/gin"
	"github.com/long2ice/fastgo"
	"github.com/long2ice/fastgo/constants"
	"github.com/long2ice/fastgo/router"
	"github.com/long2ice/fastgo/swagger"
	"net/http"
)

type Model struct {
	Name string `form:"name" binding:"required" json:"name"`
}

func main() {
	app := fastgo.Default(swagger.Default("FastGo", "Gin + Swagger = FastGo", "0.1.0"))
	app.GET("/ping", router.Default(
		router.Handlers(func(c *gin.Context) {
			model := c.MustGet(constants.Model)
			c.JSON(http.StatusOK, model)
		}),
		router.Summary("Ping"),
		router.Model(&Model{}),
	))
	if err := app.Run(); err != nil {
		panic(err)
	}
}
