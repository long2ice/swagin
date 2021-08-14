# Gin + Swagger = FastGo

## Introduction

`FastGo` is a web framework based on `Gin` and `Swagger`, which wraps `Gin` and provides built-in swagger api docs.

## Why I build this project?

Before I have used [FastAPI](https://github.com/tiangolo/fastapi), which gives me a good experience in api docs
generation, because nobody like writing api docs.

Now I use `Gin` but I can't found anything like that, I found [swag](https://github.com/swaggo/swag) but write docs with
comment is so stupid. So there is `FastGo`.

## Installation

```shell
go get -u github.com/long2ice/fastgo
```

## Online Demo

You can see online demo at <https://fastgo.long2ice.cn>

## Usage

### Build Swagger

Firstly, build a swagger object with basic information.

```go
package examples

import (
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/long2ice/fastgo/swagger"
)

func NewSwagger() *swagger.Swagger {
	return swagger.Default("FastGo", "Gin + Swagger = FastGo", "0.1.0",
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
}
```

### Write API

Then make api struct which implement `router.IAPI`.

```go
package examples

type QueryModel struct {
	Name string `query:"name" binding:"required" json:"name" description:"name of model" default:"test"`
}
type TestQuery struct {
	Model *QueryModel
}

func (t *TestQuery) NewModel() interface{} {
	return &QueryModel{}
}

func (t *TestQuery) Handler(c *gin.Context) {
	c.JSON(http.StatusOK, t.Model)
}
```

Note that the `QueryModel`? `FastGo` will validate request and inject it automatically, then you can use it in handler
easily.

### Write Router

Then write router with some docs configuration and api.

```go
package examples

var query = router.Default(
	router.API(&TestQuery{}),
	router.Summary("Test Query"),
	router.Description("Test Query Model"),
	router.Tags("Test"),
)
```

## Start APP

Finally, start the application with routes defined.

```go
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

```

That's all! Now you can visit <http://127.0.0.1:8080/docs> or <http://127.0.0.1:8080/redoc> to see the api docs. Have
fun!

## License

This project is licensed under the
[Apache-2.0](https://github.com/long2ice/fastgo/blob/master/LICENSE)
License.