package handlers

import (
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	sa "github.com/savsgio/atreugo/v11"
	"github.com/savsgio/go-logger/v2"
	"github.com/valyala/fasthttp"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app/domain"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app/security"
	"strconv"
	"strings"
)

func (h *Handlers) SignIn(ctx *sa.RequestCtx) error {

	token, err := h.signIn(ctx)

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

func (h *Handlers) signIn(ctx *sa.RequestCtx) (*domain.Token, error) {

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
	id := uuid.New()
	password := security.HashAndSalt([]byte(s.Password))
	login := domain.Login{Username: s.Username}
	login.SetId(id)
	login.SetPassword(password)
	err = h.Server.DAO.Login.Create(&login)

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
	ins := strings.Split(s.Interests, "\n")
	user := domain.Create(id, s.Username, &s.Name, &s.Surname, int(age), int(sex), ins, &s.City, true)

	if logger.DebugEnabled() {
		logger.Debugf("got user: %s", user.String())
	}
	err = h.Server.DAO.User.Create(user)

	if err != nil {
		return nil, err
	}
	err = h.Server.DAO.Interest.OldCreateInterests(ins)

	if err != nil && logger.DebugEnabled() {
		logger.Debugf("OldCreateInterests has error: %v", err)
	}
	interests, _ := h.Server.DAO.Interest.GetExistsInterests(ins)
	err = h.Server.DAO.UserHasInterests.LinkInterests(user, interests)

	if err != nil && logger.DebugEnabled() {
		logger.Debugf("LinkInterests has error: %v", err)
	}
	err = h.Server.DAO.Session.UpdateOrCreate(&login, id)

	if err != nil && logger.DebugEnabled() {
		logger.Debugf("UpdateOrCreate has error: %v", err)
	}
	return h.generateToken(ctx, id), nil
}
