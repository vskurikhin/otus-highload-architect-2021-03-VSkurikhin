package handlers

import (
	"encoding/json"
	"errors"
	sa "github.com/savsgio/atreugo/v11"
	"github.com/savsgio/go-logger/v2"
	"github.com/valyala/fasthttp"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app-dialog/config"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app-dialog/domain"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app-dialog/security"
	"strconv"
)

func (h *Handlers) UserSignIn(ctx *sa.RequestCtx) error {

	token, err := h.userSignIn(ctx)

	if err != nil {
		logger.Error(err)
		errorCase := domain.ApiMessage{
			Code:    fasthttp.StatusPreconditionFailed,
			Message: err.Error(),
		}
		return ctx.HTTPResponse(errorCase.String(), fasthttp.StatusPreconditionFailed)
	}

	return ctx.HTTPResponse(token.String())
}

func (h *Handlers) userSignIn(ctx *sa.RequestCtx) (*domain.Token, error) {

	var s domain.Signin
	err := json.Unmarshal(ctx.PostBody(), &s)

	if err != nil {
		return nil, err // правильная обработка ошибок вместо паники
	}
	err = security.CheckSignIn(&s)

	if err != nil {
		return nil, err
	}
	l, err := h.Server.DAO.Login.Read(s.Username)

	if l != nil {
		return nil, errors.New(" User with Username: " + s.Username + " already exists ")
	}

	if logger.DebugEnabled() {
		logger.Debugf("got s: %s", s.String())
	}
	password := security.HashAndSalt([]byte(s.Password))
	login := domain.Login{Username: s.Username, Password: password}
	l, err = h.Server.DAO.Login.Create(&login)
	logger.Debugf("l: %s", l)

	if err != nil {
		return nil, err
	}
	age, err := strconv.ParseInt(s.Age, 10, 64)

	if err != nil {
		return nil, err
	}
	sex, err := strconv.ParseInt(s.Sex, 10, 64)

	if err != nil {
		return nil, err
	}
	user := domain.User{
		Id:       l.Id,
		Username: s.Username,
		Name:     &s.Name,
		SurName:  &s.Surname,
		Age:      int(age),
		Sex:      int(sex),
		City:     &s.City,
		Friend:   true,
	}

	if logger.DebugEnabled() {
		logger.Debugf("got user: %s", user.String())
	}
	_, err = h.Server.DAO.User.Create(&user)

	if err != nil {
		return nil, err
	}
	return h.generateToken(ctx, l.Id), nil
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
