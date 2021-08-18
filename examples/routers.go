package main

import (
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gin-gonic/gin/binding"
	"github.com/long2ice/fastgo/router"
	"github.com/long2ice/fastgo/security"
)

var (
	query = router.New(
		router.API(&TestQuery{}),
		router.Summary("Test query"),
		router.Description("Test query model"),
		router.Security(&security.Basic{}),
		router.Responses(openapi3.Responses{
			"200": &openapi3.ResponseRef{
				Value: &openapi3.Response{
					Content: openapi3.NewContentWithJSONSchema(&openapi3.Schema{
						Example:     QueryModel{},
						Description: "Response model",
					}),
				},
			},
		}),
	)
	noModel = router.New(
		router.API(&TestNoModel{}),
		router.Summary("Test no model"),
		router.Description("Test no model"),
	)
	queryPath = router.New(
		router.API(&TestQueryPath{}),
		router.Summary("Test query path"),
		router.Description("Test query path model"),
	)
	formEncode = router.New(
		router.API(&TestForm{}),
		router.Summary("Test form"),
		router.ContentType(binding.MIMEPOSTForm),
	)
	body = router.New(
		router.API(&TestForm{}),
		router.Summary("Test json body"),
	)
	file = router.New(
		router.API(&TestFile{}),
		router.Summary("Test file upload"),
		router.ContentType(binding.MIMEMultipartPOSTForm),
	)
)
