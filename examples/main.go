package main

import (
	"github.com/gin-contrib/cors"
	"github.com/long2ice/swagin"
	"github.com/long2ice/swagin/security"
)

func main() {
	app := swagin.New(NewSwagger())
	subApp := swagin.New(NewSwagger())
	subApp.GET("/noModel", noModel)
	app.Mount("/sub", subApp)
	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))
	queryGroup := app.Group("/query", swagin.Tags("Query"))
	queryGroup.GET("/list", queryList)
	queryGroup.GET("/:id", queryPath)
	queryGroup.DELETE("", query)

	app.GET("/noModel", noModel)

	formGroup := app.Group("/form", swagin.Tags("Form"), swagin.Security(&security.Bearer{}))
	formGroup.POST("/encoded", formEncode)
	formGroup.PUT("", body)
	formGroup.POST("/file", file)

	if err := app.Run(); err != nil {
		panic(err)
	}
}
