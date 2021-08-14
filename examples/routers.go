package main

import "github.com/long2ice/fastgo/router"

var (
	query = router.Default(
		router.API(&TestQuery{}),
		router.Summary("Test query"),
		router.Description("Test query model"),
	)
	queryPath = router.Default(
		router.API(&TestQueryPath{}),
		router.Summary("Test query path"),
		router.Description("Test query path model"),
	)
	formEncode = router.Default(
		router.API(&TestForm{}),
		router.Summary("Test form"),
		router.ContentType("application/x-www-form-urlencoded"),
	)
	body = router.Default(
		router.API(&TestBody{}),
		router.Summary("Test json body"),
	)
)
