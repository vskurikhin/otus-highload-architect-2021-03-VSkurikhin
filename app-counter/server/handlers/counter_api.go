package handlers

import (
	"fmt"
	sa "github.com/savsgio/atreugo/v11"
	"github.com/savsgio/go-logger/v2"
	"github.com/valyala/fasthttp"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app-counter/domain"
)

func (h *Handlers) GetCounter(ctx *sa.RequestCtx) error {

	counter, err := h.getCounter(ctx)

	if err != nil {
		logger.Error(err)
		errorCase := domain.ApiMessage{
			Code:    fasthttp.StatusPreconditionFailed,
			Message: err.Error(),
		}
		return ctx.HTTPResponse(errorCase.String(), fasthttp.StatusPreconditionFailed)
	}
	return ctx.HTTPResponse(counter.String())
}

func (h *Handlers) getCounter(ctx *sa.RequestCtx) (*domain.Counter, error) {

	username := fmt.Sprintf("%v", ctx.UserValue("username"))
	counter, err := h.Server.DAO.Counter.ReadByUserName(username)
	if err != nil {
		return nil, err
	}
	return counter, nil
}
