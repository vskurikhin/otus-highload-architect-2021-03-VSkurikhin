package home

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
)

func Index(ctx *fasthttp.RequestCtx) {

	ctx.SetContentType("text/html")
	_, err := fmt.Fprint(ctx, "Welcome to Home Page!\n")

	if err != nil {
		logger := logrus.WithError(err)
		logger.Fatalln("main: can't load configuration")
	}
}
