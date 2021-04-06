package handlers

import (
	"github.com/google/uuid"
	sa "github.com/savsgio/atreugo/v11"
	"github.com/valyala/fasthttp"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app/config"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app/message"
	"log"
)

func (h *Handlers) Login(ctx *sa.RequestCtx) error {
	sessionId := uuid.New()

	jwtCookie := ctx.Request.Header.Cookie(config.ACCESS_TOKEN_COOKIE)

	if len(jwtCookie) == 0 {
		if ctx.IsBodyStream() {
			log.Println(ctx.PostBody())
		}

		tokenString, expireAt := h.Server.JWT.GenerateToken(sessionId)

		// Set cookie for domain
		cookie := fasthttp.AcquireCookie()
		defer fasthttp.ReleaseCookie(cookie)

		cookie.SetKey(config.ACCESS_TOKEN_COOKIE)
		cookie.SetValue(tokenString)
		cookie.SetExpire(expireAt)
		ctx.Response.Header.SetCookie(cookie)

		token := message.Token{Token: tokenString}

		return ctx.HTTPResponse(token.String())
	}
	token := message.Token{Token: string(jwtCookie)}

	return ctx.HTTPResponse(token.String())
}
