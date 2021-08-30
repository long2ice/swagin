package main

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
			URL:   "https://github.com/long2ice/swagin",
			Email: "long2ice@gmail.com",
		}),
		swagger.TermsOfService("https://github.com/long2ice"),
	)
}
