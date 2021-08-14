package main

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
