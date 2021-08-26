package router

import "github.com/getkin/kin-openapi/openapi3"

type Response map[string]ResponseItem

type ResponseItem struct {
	Description string
	Model       interface{}
	Headers     openapi3.Headers
}
