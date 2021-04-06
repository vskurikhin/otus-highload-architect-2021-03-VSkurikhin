package handlers

import (
	"fmt"
	sa "github.com/savsgio/atreugo/v11"
	"github.com/valyala/fasthttp"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app/config"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app/message"
)

func (h *Handlers) Root(ctx *sa.RequestCtx) error {
	return ctx.HTTPResponse(root(ctx))
}

func root(ctx *sa.RequestCtx) string {
	return rootOut(ctx.Request.Header.Cookie(config.ACCESS_TOKEN_COOKIE))
}

func rootOut(jwt []byte) string {

	msg := fmt.Sprintf(`You are login with JWT cookie value: %s`, jwt)
	apiMessage := message.ApiMessage{Code: fasthttp.StatusOK, Message: msg}

	return apiMessage.String()
}
