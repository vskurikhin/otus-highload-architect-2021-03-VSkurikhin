package handlers

import (
	sa "github.com/savsgio/atreugo/v11"
	"github.com/savsgio/go-logger/v2"
	"github.com/valyala/fasthttp"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app/domain"
	"strings"
)

func (h *Handlers) List(ctx *sa.RequestCtx) error {

	list, err := h.list(ctx)

	if err != nil {
		logger.Error(err)
		errorCase := domain.ApiMessage{
			Code:    fasthttp.StatusPreconditionFailed,
			Message: err.Error(),
		}
		return ctx.HTTPResponse(errorCase.String(), fasthttp.StatusPreconditionFailed)
	}
	result := "[" + strings.Join(list, ", ") + "]"

	return ctx.HTTPResponse(result)
}

func (h *Handlers) list(ctx *sa.RequestCtx) ([]string, error) {

	p, err := h.profile(ctx)

	if err != nil {
		return nil, err
	}
	users, err := h.Server.DAO.User.ReadUserList(p.Id)

	if err != nil {
		return nil, err
	}
	var result []string

	for _, user := range users {
		result = append(result, user.String())
	}
	return result, nil
}
