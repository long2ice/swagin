package main

import "github.com/long2ice/fastgo/router"

var (
	query = router.New(
		router.API(&TestQuery{}),
		router.Summary("Test query"),
		router.Description("Test query model"),
	)
	queryPath = router.New(
		router.API(&TestQueryPath{}),
		router.Summary("Test query path"),
		router.Description("Test query path model"),
	)
	formEncode = router.New(
		router.API(&TestForm{}),
		router.Summary("Test form"),
		router.ContentType("application/x-www-form-urlencoded"),
	)
	body = router.New(
		router.API(&TestBody{}),
		router.Summary("Test json body"),
	)
)
