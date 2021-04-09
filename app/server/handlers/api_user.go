package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	sa "github.com/savsgio/atreugo/v11"
	"github.com/savsgio/go-logger/v2"
	"github.com/valyala/fasthttp"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app/domain"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app/security"
	"strconv"
	"strings"
)

func (h *Handlers) List(ctx *sa.RequestCtx) error {

	u, err := h.Server.DAO.User.ReadListAsString()

	if err != nil {
		logger.Error(err)
		errorCase := domain.ApiMessage{
			Code:    fasthttp.StatusPreconditionFailed,
			Message: err.Error(),
		}
		return ctx.HTTPResponse(errorCase.String(), fasthttp.StatusPreconditionFailed)
	}
	return ctx.HTTPResponse(u)
}

func (h *Handlers) Signin(ctx *sa.RequestCtx) error {

	var signIn domain.Signin
	err := json.Unmarshal(ctx.PostBody(), &signIn)

	if err != nil {
		logger.Error(err)
		errorCase := domain.ApiMessage{
			Code:    fasthttp.StatusPreconditionFailed,
			Message: err.Error(),
		}
		return ctx.HTTPResponse(errorCase.String(), fasthttp.StatusPreconditionFailed)
	}
	if logger.DebugEnabled() {
		logger.Debugf("got signIn: %s", signIn.String())
	}
	id := uuid.New()
	password := security.HashAndSalt([]byte(signIn.Password))
	login := domain.Login{Username: signIn.Username}
	login.SetId(id)
	login.SetPassword(password)
	err = h.Server.DAO.Login.Create(&login)
	if err != nil {
		logger.Error(err)
		errorCase := domain.ApiMessage{
			Code:    fasthttp.StatusPreconditionFailed,
			Message: err.Error(),
		}
		return ctx.HTTPResponse(errorCase.String(), fasthttp.StatusPreconditionFailed)
	}
	age, _ := strconv.ParseInt(signIn.Age, 10, 32)
	sex, _ := strconv.ParseInt(signIn.Age, 10, 1)
	ins := strings.Split(signIn.Interests, "\n")
	user := domain.Create(id, signIn.Username, &signIn.Name, &signIn.Surname, int(age), int(sex), ins, &signIn.City)
	if logger.DebugEnabled() {
		logger.Debugf("got user: %s", user.String())
	}
	err = h.Server.DAO.User.Create(user)
	if err != nil {
		logger.Error(err)
		errorCase := domain.ApiMessage{
			Code:    fasthttp.StatusPreconditionFailed,
			Message: err.Error(),
		}
		return ctx.HTTPResponse(errorCase.String(), fasthttp.StatusPreconditionFailed)
	}
	_ = h.Server.DAO.Interest.CreateInterests(ins)
	interests, _ := h.Server.DAO.Interest.GetExistsInterests(ins)
	_ = h.Server.DAO.UserHasInterests.LinkInterests(user, interests)
	token := h.generateToken(ctx, id)

	return ctx.HTTPResponse(token.String())
}

func (h *Handlers) Create(ctx *sa.RequestCtx) error {

	var user domain.User
	err := json.Unmarshal(ctx.PostBody(), &user)

	if err != nil {
		logger.Error(err)
		errorCase := domain.ApiMessage{
			Code:    fasthttp.StatusPreconditionFailed,
			Message: err.Error(),
		}
		return ctx.HTTPResponse(errorCase.String(), fasthttp.StatusPreconditionFailed)
	}
	user.SetId(uuid.New())

	if logger.DebugEnabled() {
		logger.Debugf("got user: %s", user.String())
	}
	err = h.Server.DAO.User.Create(&user)

	if err != nil {
		logger.Error(err)
		errorCase := domain.ApiMessage{
			Code:    fasthttp.StatusPreconditionFailed,
			Message: err.Error(),
		}
		return ctx.HTTPResponse(errorCase.String(), fasthttp.StatusPreconditionFailed)
	}
	msg := fmt.Sprintf("created with id: %s", user.Id().String())
	created := domain.ApiMessage{
		Code:    fasthttp.StatusCreated,
		Message: msg,
	}
	return ctx.HTTPResponse(created.String())
}
