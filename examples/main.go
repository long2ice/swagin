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
	app.GET("/query", query)
	app.GET("/query/:id", queryPath)
	app.DELETE("/query", query)
	app.POST("/body", body)
	app.POST("/form/encoded", formEncode)
	app.PUT("/form", body)
	if err := app.Run(); err != nil {
		panic(err)
	}
}
