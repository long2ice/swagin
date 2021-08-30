# Gin + Swagger = SwaGin

[![deploy](https://github.com/long2ice/swagin/actions/workflows/deploy.yml/badge.svg)](https://github.com/long2ice/swagin/actions/workflows/deploy.yml)

## Introduction

`SwaGin` is a web framework based on `Gin` and `Swagger`, which wraps `Gin` and provides built-in swagger api docs and
request model validation.

## Why I build this project?

Previous I have used [FastAPI](https://github.com/tiangolo/fastapi), which gives me a great experience in api docs
generation, because nobody like writing api docs.

Now I use `Gin` but I can't found anything like that, I found [swag](https://github.com/swaggo/swag) but which write
docs with comment is so stupid. So there is `SwaGin`.

## Installation

```shell
go get -u github.com/long2ice/swagin
```

## Online Demo

You can see online demo at <https://swagin.long2ice.io/docs> or <https://swagin.long2ice.io/redoc>.

![](https://raw.githubusercontent.com/long2ice/swagin/dev/images/docs.png)
![](https://raw.githubusercontent.com/long2ice/swagin/dev/images/redoc.png)

And you can reference all usage in [examples](https://github.com/long2ice/swagin/tree/dev/examples).

## Usage

### Build Swagger

Firstly, build a swagger object with basic information.

```go
package examples

import (
  "github.com/getkin/kin-openapi/openapi3"
  "github.com/long2ice/swagin/swagger"
)

func NewSwagger() *swagger.Swagger {
  return swagger.New("SwaGin", "Gin + Swagger = SwaGin", "0.1.0",
    swagger.License(&openapi3.License{
      Name: "Apache License 2.0",
      URL:  "https://github.com/long2ice/swagin/blob/dev/LICENSE",
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

type TestQuery struct {
  Name string `query:"name" binding:"required" json:"name" description:"name of model" default:"test"`
}

func (t *TestQuery) Handler(c *gin.Context) {
  c.JSON(http.StatusOK, t)
}
```

Note that the attributes in `TestQuery`? `SwaGin` will validate request and inject it automatically, then you can use it
in handler easily.

### Write Router

Then write router with some docs configuration and api.

```go
package examples

var query = router.New(
  &TestQuery{},
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
  &TestQuery{},
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
  c.JSON(http.StatusOK, t)
}
```

### Mount Router

Then you can mount router in your application or group.

```go
app := swagin.New(NewSwagger())
queryGroup := app.Group("/query", swagin.Tags("Query"))
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
  "github.com/long2ice/swagin"
)

func main() {
  app := swagin.New(NewSwagger())
  app.Use(cors.New(cors.Config{
    AllowOrigins:     []string{"*"},
    AllowMethods:     []string{"*"},
    AllowHeaders:     []string{"*"},
    AllowCredentials: true,
  }))

  queryGroup := app.Group("/query", swagin.Tags("Query"))
  queryGroup.GET("", query)
  queryGroup.GET("/:id", queryPath)
  queryGroup.DELETE("", query)

  formGroup := app.Group("/form", swagin.Tags("Form"))
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

### Disable Docs

In some cases you may want to disable docs such as in production, just put `nil` to `swagin.New`.

```go
app = swagin.New(nil)
```

### SubAPP Mount

If you want to use sub application, you can mount another `SwaGin` instance to main application, and their swagger docs
is also separate.

```go
package main

func main() {
  app := swagin.New(NewSwagger())
  subApp := swagin.New(NewSwagger())
  subApp.GET("/noModel", noModel)
  app.Mount("/sub", subApp)
}

```

## ThanksTo

- [kin-openapi](https://github.com/getkin/kin-openapi), OpenAPI 3.0 implementation for Go (parsing, converting,
  validation, and more).
- [Gin](https://github.com/gin-gonic/gin), an HTTP web framework written in Go (Golang).

## License

This project is licensed under the
[Apache-2.0](https://github.com/long2ice/swagin/blob/master/LICENSE)
License.