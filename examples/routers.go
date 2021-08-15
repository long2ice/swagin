package main

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/long2ice/fastgo/router"
)

var (
	query = router.New(
		router.API(&TestQuery{}),
		router.Summary("Test query"),
		router.Description("Test query model"),
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
)
