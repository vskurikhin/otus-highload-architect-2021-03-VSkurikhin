package handlers

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"github.com/hashicorp/consul/connect"
	sa "github.com/savsgio/atreugo/v11"
	"github.com/savsgio/go-logger/v2"
	"github.com/valyala/fasthttp"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app-main/client"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app-main/domain"
)

func (h *Handlers) PostMessage(ctx *sa.RequestCtx) error {

	// Создать клиента Resty
	c, err := client.NewDialog(ctx)
	if err != nil {
		logger.Error(err)
		errorCase := domain.ApiMessage{
			Code:    fasthttp.StatusPreconditionFailed,
			Message: err.Error(),
		}
		return ctx.HTTPResponse(errorCase.String(), fasthttp.StatusPreconditionFailed)
	}
	resp, err := c.R().
		SetBody(ctx.PostBody()).
		Post(fmt.Sprintf("%s/%s", h.Server.Services.Dialog, "message"))
	if err != nil {
		logger.Error(err)
		errorCase := domain.ApiMessage{
			Code:    fasthttp.StatusPreconditionFailed,
			Message: err.Error(),
		}
		return ctx.HTTPResponse(errorCase.String(), fasthttp.StatusPreconditionFailed)
	}

	return ctx.HTTPResponse(resp.String())
}

func (h *Handlers) GetMessages(ctx *sa.RequestCtx) error {

	// Создать клиента Resty
	c, err := client.NewDialog(ctx)
	if err != nil {
		logger.Error(err)
		errorCase := domain.ApiMessage{
			Code:    fasthttp.StatusPreconditionFailed,
			Message: err.Error(),
		}
		return ctx.HTTPResponse(errorCase.String(), fasthttp.StatusPreconditionFailed)
	}
	client, _ := api.NewClient(api.DefaultConfig())

	// Create an instance representing this service. "my-service" is the
	// name of _this_ service. The service should be cleaned up via Close.
	svc, _ := connect.NewService("my-app-dialog", client)
	defer svc.Close()

	<-svc.ReadyWait()
	logger.Infof("%s", svc.HTTPClient())

	resp, err := c.R().Get(fmt.Sprintf("%s/%s", h.Server.Services.Dialog, "messages"))
	if err != nil {
		logger.Error(err)
		errorCase := domain.ApiMessage{
			Code:    fasthttp.StatusPreconditionFailed,
			Message: err.Error(),
		}
		return ctx.HTTPResponse(errorCase.String(), fasthttp.StatusPreconditionFailed)
	}
	return ctx.HTTPResponse(resp.String())
}
