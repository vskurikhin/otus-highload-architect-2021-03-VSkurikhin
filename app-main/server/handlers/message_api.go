package handlers

import (
	"fmt"
	sa "github.com/savsgio/atreugo/v11"
	"github.com/savsgio/go-logger/v2"
	"github.com/valyala/fasthttp"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app-main/domain"
)

type message struct {
	ToUser  string
	Message string
}

func (h *Handlers) PostMessage(ctx *sa.RequestCtx) error {

	// id, err := h.postMessage(ctx)
	var id int
	var err error

	if err != nil {
		logger.Error(err)
		errorCase := domain.ApiMessage{
			Code:    fasthttp.StatusPreconditionFailed,
			Message: err.Error(),
		}
		return ctx.HTTPResponse(errorCase.String(), fasthttp.StatusPreconditionFailed)
	}

	return ctx.HTTPResponse(fmt.Sprintf("%d", id))
}

func (h *Handlers) GetMessages(ctx *sa.RequestCtx) error {
	// TODO
	//list, err := h.getMessages(ctx)
	//
	//if err != nil {
	//	logger.Error(err)
	//	errorCase := domain.ApiMessage{
	//		Code:    fasthttp.StatusPreconditionFailed,
	//		Message: err.Error(),
	//	}
	//	return ctx.HTTPResponse(errorCase.String(), fasthttp.StatusPreconditionFailed)
	//}
	result := "[]"

	return ctx.HTTPResponse(result)
}
