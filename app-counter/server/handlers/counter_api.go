package handlers

import (
	sa "github.com/savsgio/atreugo/v11"
	"github.com/savsgio/go-logger/v2"
	"github.com/valyala/fasthttp"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app-counter/domain"
)

func (h *Handlers) GetCounter(ctx *sa.RequestCtx) error {

	// list, err := h.getMessages(ctx)
	var err error
	var result string

	if err != nil {
		logger.Error(err)
		errorCase := domain.ApiMessage{
			Code:    fasthttp.StatusPreconditionFailed,
			Message: err.Error(),
		}
		return ctx.HTTPResponse(errorCase.String(), fasthttp.StatusPreconditionFailed)
	}
	// result := "[" + strings.Join(list, ", ") + "]"

	return ctx.HTTPResponse(result)
}
