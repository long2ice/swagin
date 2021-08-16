package main

import (
	"github.com/gin-contrib/cors"
	"github.com/long2ice/fastgo"
)

func main() {
	app := fastgo.New(NewSwagger())
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
	app.POST("/body", body)
	formGroup := app.Group("/form", fastgo.Tags("Form"))
	formGroup.POST("/encoded", formEncode)
	formGroup.PUT("", body)
	if err := app.Run(); err != nil {
		panic(err)
	}
}
