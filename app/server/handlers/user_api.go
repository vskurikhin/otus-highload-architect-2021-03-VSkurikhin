package handlers

import (
	"encoding/json"
	"fmt"
	sa "github.com/savsgio/atreugo/v11"
	"github.com/savsgio/go-logger/v2"
	"github.com/valyala/fasthttp"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app/domain"
)

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
	user.NewId()

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

func (h *Handlers) User(ctx *sa.RequestCtx) error {

	id := fmt.Sprintf("%v", ctx.UserValue("id"))
	if logger.DebugEnabled() {
		logger.Debugf("got user id: %s", id)
	}
	u, err := h.user(ctx)

	if err != nil {
		logger.Error(err)
		errorCase := domain.ApiMessage{
			Code:    fasthttp.StatusPreconditionFailed,
			Message: err.Error(),
		}
		return ctx.HTTPResponse(errorCase.String(), fasthttp.StatusPreconditionFailed)
	}
	return ctx.HTTPResponse(u.String())
}

func (h *Handlers) user(ctx *sa.RequestCtx) (*domain.User, error) {

	id := fmt.Sprintf("%v", ctx.UserValue("id"))

	if logger.DebugEnabled() {
		logger.Debugf("got user id: %s", id)
	}
	u, err := h.Server.DAO.User.ReadUser(id)

	if err != nil {
		return nil, err
	}
	p, _ := h.profile(ctx)

	if p != nil {
		f, _ := h.Server.DAO.UserHasFriends.IsFriendship(p.Id, u.Id())
		u.Friend = f
	}
	return u, nil
}
