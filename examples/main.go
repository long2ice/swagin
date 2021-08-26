package main

import (
	"github.com/gin-contrib/cors"
	"github.com/long2ice/fastgo"
	"github.com/long2ice/fastgo/security"
)

func main() {
	app := fastgo.New(NewSwagger())
	subApp := fastgo.New(NewSwagger())
	subApp.GET("/noModel", noModel)
	app.Mount("/sub", subApp)
	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))
	queryGroup := app.Group("/query", fastgo.Tags("Query"))
	queryGroup.GET("", query)
	queryGroup.GET("/:id", queryPath)
	queryGroup.DELETE("", query)

	app.GET("/noModel", noModel)

	formGroup := app.Group("/form", fastgo.Tags("Form"), fastgo.Security(&security.Bearer{}))
	formGroup.POST("/encoded", formEncode)
	formGroup.PUT("", body)
	formGroup.POST("/file", file)

	if err := app.Run(); err != nil {
		panic(err)
	}
}
