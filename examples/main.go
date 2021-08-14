package main

import (
	"github.com/gin-contrib/cors"
	"github.com/long2ice/fastgo"
)

func main() {
	app := fastgo.Default(NewSwagger())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))
	app.GET("/query", query)
	app.DELETE("/query", query)
	app.POST("/form", form)
	app.PUT("/form", form)
	if err := app.Run(); err != nil {
		panic(err)
	}
}
