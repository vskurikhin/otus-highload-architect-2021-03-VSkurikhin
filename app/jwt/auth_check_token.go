package jwt

import (
	"errors"
	sa "github.com/savsgio/atreugo/v11"
	"github.com/valyala/fasthttp"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app/config"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app/message"
)

// AuthCheckToken проверка авторизации по токену jwt
func (j *JWT) AuthCheckToken(ctx *sa.RequestCtx) error {
	// пропускаем точку авторизации.
	if string(ctx.Path()) == "/login" {
		return ctx.Next()
	}

	jwtCookie := ctx.Request.Header.Cookie(config.ACCESS_TOKEN_COOKIE)

	if len(jwtCookie) == 0 {
		return ctx.ErrorResponse(errors.New(message.LoginRequired.String()), fasthttp.StatusForbidden)
	}

	token, err := j.ValidateToken(string(jwtCookie))
	if err != nil {
		return err
	}

	if !token.Valid {
		return ctx.ErrorResponse(errors.New(message.YourSessionIsExpired.String()), fasthttp.StatusForbidden)
	}

	return ctx.Next()
}
