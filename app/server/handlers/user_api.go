package handlers

import (
	"fmt"
	sa "github.com/savsgio/atreugo/v11"
	"github.com/savsgio/go-logger/v2"
	"github.com/valyala/fasthttp"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app/domain"
)

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
	return u, nil
}
