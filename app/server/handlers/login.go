package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	sa "github.com/savsgio/atreugo/v11"
	"github.com/savsgio/go-logger/v2"
	"github.com/valyala/fasthttp"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app/config"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app/domain"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app/security"
	"golang.org/x/crypto/bcrypt"
)

type login struct {
	Username string
	Password string
}

func (l *login) String() string {
	return fmt.Sprintf("{username: %s, password: %s}", l.Username, l.Password)
}

func (h *Handlers) Login(ctx *sa.RequestCtx) error {

	err := h.login(ctx)

	if err == nil {
		sessionId := uuid.New()
		jwtCookie := ctx.Request.Header.Cookie(config.ACCESS_TOKEN_COOKIE)

		if len(jwtCookie) == 0 {
			token := h.generateToken(ctx, sessionId)

			return ctx.HTTPResponse(token.String())
		}
		token := domain.Token{Token: string(jwtCookie)}

		return ctx.HTTPResponse(token.String())
	}
	logger.Errorf("Bad password or error: %v", err)

	return ctx.HTTPResponse("{}", fasthttp.StatusForbidden)
}

func (h *Handlers) login(ctx *sa.RequestCtx) error {

	var dto login
	err := json.Unmarshal(ctx.PostBody(), &dto)

	if err != nil {
		return err // правильная обработка ошибок вместо паники
	}
	err = security.CheckValue(dto.Username)
	if err != nil {
		return err
	}

	login, err := h.Server.DAO.Login.Read(dto.Username)

	if err != nil {
		return err
	}
	err = bcrypt.CompareHashAndPassword([]byte(login.Password()), []byte(dto.Password))

	return err
}

func (h *Handlers) generateToken(ctx *sa.RequestCtx, sessionId uuid.UUID) *domain.Token {

	tokenString, expireAt := h.Server.JWT.GenerateToken(sessionId)

	// Set cookie for domain
	cookie := fasthttp.AcquireCookie()
	defer fasthttp.ReleaseCookie(cookie)

	cookie.SetKey(config.ACCESS_TOKEN_COOKIE)
	cookie.SetValue(tokenString)
	cookie.SetExpire(expireAt)
	ctx.Response.Header.SetCookie(cookie)

	token := domain.Token{Token: tokenString}

	return &token
}
