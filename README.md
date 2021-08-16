# Gin + Swagger = FastGo

[![deploy](https://github.com/long2ice/fastgo/actions/workflows/deploy.yml/badge.svg)](https://github.com/long2ice/fastgo/actions/workflows/deploy.yml)

## Introduction

`FastGo` is a web framework based on `Gin` and `Swagger`, which wraps `Gin` and provides built-in swagger api docs and
request model validation.

## Why I build this project?

Previous I have used [FastAPI](https://github.com/tiangolo/fastapi), which gives me a great experience in api docs
generation, because nobody like writing api docs.

Now I use `Gin` but I can't found anything like that, I found [swag](https://github.com/swaggo/swag) but which write
docs with comment is so stupid. So there is `FastGo`.

## Installation

```shell
go get -u github.com/long2ice/fastgo
```

## Online Demo

You can see online demo at <https://fastgo.long2ice.cn/docs> or <https://fastgo.long2ice.cn/redoc>.

![](https://raw.githubusercontent.com/long2ice/fastgo/dev/images/docs.png)
![](https://raw.githubusercontent.com/long2ice/fastgo/dev/images/redoc.png)

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

var query = router.New(
  router.API(&TestQuery{}),
  router.Summary("Test Query"),
  router.Description("Test Query Model"),
  router.Tags("Test"),
)
```

### Security

If you want to project your api with a security policy, you can use security, also they will be shown in swagger docs.

Current there is five kinds of security policies.

- `Basic`
- `Bearer`
- `ApiKey`
- `OpenID`
- `OAuth2`

```go
package main

var query = router.New(
  router.API(&TestQuery{}),
  router.Summary("Test query"),
  router.Description("Test query model"),
  router.Security(&security.Basic{}),
)
```

Then you can get the authentication string by `context.MustGet(security.Credentials)` depending on your auth type.

```go
package main

func (t *TestQuery) Handler(c *gin.Context) {
  user := c.MustGet(security.Credentials).(*security.User)
  fmt.Println(user)
  c.JSON(http.StatusOK, t.Model)
}
```

### Mount Router

Then you can mount router in your application or group.

```go
app := fastgo.New(NewSwagger())
queryGroup := app.Group("/query", fastgo.Tags("Query"))
queryGroup.GET("", query)
queryGroup.GET("/:id", queryPath)
queryGroup.DELETE("", query)

app.GET("/noModel", noModel)
```

### Start APP

Finally, start the application with routes defined.

```go
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

  formGroup := app.Group("/form", fastgo.Tags("Form"))
  formGroup.POST("/encoded", formEncode)
  formGroup.PUT("", body)

  app.GET("/noModel", noModel)
  app.POST("/body", body)
  if err := app.Run(); err != nil {
    panic(err)
  }
}
```

That's all! Now you can visit <http://127.0.0.1:8080/docs> or <http://127.0.0.1:8080/redoc> to see the api docs. Have
fun!

## ThanksTo

- [kin-openapi](https://github.com/getkin/kin-openapi), OpenAPI 3.0 implementation for Go (parsing, converting,
  validation, and more).
- [Gin](https://github.com/gin-gonic/gin), an HTTP web framework written in Go (Golang).

## License

This project is licensed under the
[Apache-2.0](https://github.com/long2ice/fastgo/blob/master/LICENSE)
License.