package controllers

import (
	"github.com/kataras/iris"
)

func NotFound(ctx iris.Context) {
	type NotFound struct {
		Message string
	}

	notFound := NotFound{
		Message: "Route not found.",
	}
	ctx.JSON(notFound)
}
