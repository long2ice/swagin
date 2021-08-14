package main

import "github.com/long2ice/fastgo/router"

var (
	query = router.Default(
		router.API(&TestQuery{}),
		router.Summary("Test Query"),
		router.Description("Test Query Model"),
		router.Tags("Test"),
	)
	form = router.Default(
		router.API(&TestForm{}),
		router.Summary("Test Form"),
		router.Description("Test Form Model"),
		router.Tags("Test"),
	)
)
