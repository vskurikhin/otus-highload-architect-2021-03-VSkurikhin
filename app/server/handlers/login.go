package handlers

import (
	"encoding/json"
	"fmt"
	sa "github.com/savsgio/atreugo/v11"
	"github.com/savsgio/go-logger/v2"
	"github.com/valyala/fasthttp"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app/config"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app/domain"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app/security"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app/utils"
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

	login, err := h.login(ctx)

	if err == nil {
		sessionId := utils.RandomSessionId()
		jwtCookie := ctx.Request.Header.Cookie(config.ACCESS_TOKEN_COOKIE)

		if len(jwtCookie) == 0 {

			token := h.generateToken(ctx, sessionId)
			err = h.Server.DAO.Session.UpdateOrCreate(login, sessionId)

			if err != nil {
				logger.Errorf("Bad password or error: %v", err)

				errorCase := domain.ApiMessage{
					Code:    fasthttp.StatusForbidden,
					Message: "Bad password or username",
				}
				return ctx.HTTPResponse(errorCase.String(), fasthttp.StatusForbidden)
			}
			if logger.DebugEnabled() {
				logger.Debugf("jwt for session %d created", sessionId)
			}
			return ctx.HTTPResponse(token.String())
		}
		token := domain.Token{Token: string(jwtCookie)}
		err = h.Server.DAO.Session.UpdateOrCreate(login, sessionId)

		if logger.DebugEnabled() {
			logger.Debugf("jwt for session %d updated", sessionId)
		}
		return ctx.HTTPResponse(token.String())
	}
	logger.Errorf("Bad password or error: %v", err)

	errorCase := domain.ApiMessage{
		Code:    fasthttp.StatusForbidden,
		Message: "Bad password or username",
	}
	return ctx.HTTPResponse(errorCase.String(), fasthttp.StatusForbidden)
}

func (h *Handlers) login(ctx *sa.RequestCtx) (*domain.Login, error) {

	var dto login
	err := json.Unmarshal(ctx.PostBody(), &dto)

	if err != nil {
		return nil, err // правильная обработка ошибок вместо паники
	}
	err = security.CheckValue("Username", dto.Username)
	if err != nil {
		return nil, err
	}

	login, err := h.Server.DAO.Login.Read(dto.Username)

	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(login.Password), []byte(dto.Password))
	if err != nil {
		return nil, err
	}

	return login, nil
}

func (h *Handlers) generateToken(ctx *sa.RequestCtx, sessionId uint64) *domain.Token {

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
