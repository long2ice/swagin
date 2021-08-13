package main

import (
	"fmt"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/long2ice/fastgo"
	"github.com/long2ice/fastgo/constants"
	"github.com/long2ice/fastgo/router"
	"github.com/long2ice/fastgo/swagger"
	"net/http"
)

type Model struct {
	Name string `query:"name" binding:"required" json:"name" description:"name of model" default:"test"`
}

func main() {
	s := swagger.Default("FastGo", "Gin + Swagger = FastGo", "0.1.0",
		swagger.License(&openapi3.License{
			Name: "Apache License 2.0",
			URL:  "https://www.apache.org/licenses/LICENSE-2.0",
		}),
		swagger.Contact(&openapi3.Contact{
			Name:  "long2ice",
			URL:   "https://github.com/long2ice",
			Email: "long2ice@gmail.com",
		}),
		swagger.TermsOfService("https://github.com/long2ice"),
	)
	app := fastgo.Default(s)
	app.GET("/test", router.Default(
		router.Handlers(func(c *gin.Context) {
			model := c.MustGet(constants.Model)
			fmt.Println(model)
			c.JSON(http.StatusOK, model)
		}),
		router.Summary("Test"),
		router.Description("Test api"),
		router.Model(&Model{}),
		router.Tags("Test"),
	))
	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))
	if err := app.Run(); err != nil {
		panic(err)
	}
}
