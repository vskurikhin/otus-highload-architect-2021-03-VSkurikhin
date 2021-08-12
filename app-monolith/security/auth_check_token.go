package security

import (
	"errors"
	sa "github.com/savsgio/atreugo/v11"
	"github.com/valyala/fasthttp"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app-monolith/config"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app-monolith/domain"
	"log"
	"regexp"
)

var css = regexp.MustCompile(`^/css/.*$`)
var generated = regexp.MustCompile(`^/generated/.*$`)

// AuthCheckToken проверка авторизации по токену jwt
func (j *JWT) AuthCheckToken(ctx *sa.RequestCtx) error {

	// пропускаем точку авторизации.
	path := string(ctx.Path())
	log.Println(path)
	switch {
	case path == "/":
		return ctx.Next()
	case path == "/favicon.ico":
		return ctx.Next()
	case path == "/login":
		return ctx.Next()
	case path == "/index.html":
		return ctx.Next()
	case path == "/signin":
		return ctx.Next()
	case css.MatchString(path):
		return ctx.Next()
	case generated.MatchString(path):
		return ctx.Next()
	}

	jwtCookie := ctx.Request.Header.Cookie(config.ACCESS_TOKEN_COOKIE)
	ctx.SetUserValueBytes([]byte("jwtCookie"), jwtCookie)

	if len(jwtCookie) == 0 {
		return ctx.ErrorResponse(errors.New(domain.LoginRequired.String()), fasthttp.StatusForbidden)
	}

	token, err := j.ValidateToken(string(jwtCookie))
	if err != nil {
		return err
	}

	if !token.Valid {
		return ctx.ErrorResponse(errors.New(domain.YourSessionIsExpired.String()), fasthttp.StatusForbidden)
	}

	return ctx.Next()
}
