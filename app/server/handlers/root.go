package handlers

import (
	"fmt"
	sa "github.com/savsgio/atreugo/v11"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app/config"
)

func (h *Handlers) Root(ctx *sa.RequestCtx) error {
	return ctx.HTTPResponse(root(ctx))
}

func root(ctx *sa.RequestCtx) string {
	return rootOut(ctx.Request.Header.Cookie(config.ACCESS_TOKEN_COOKIE))
}

func rootOut(jwt []byte) string {
	return fmt.Sprintf(`<h1>You are login with JWT</h1> JWT cookie value: %s`, jwt)
}
