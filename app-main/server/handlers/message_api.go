package handlers

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	sa "github.com/savsgio/atreugo/v11"
	"github.com/savsgio/go-logger/v2"
	"github.com/valyala/fasthttp"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app-main/client"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app-main/domain"
)

const SERVICE_NAME = "my-app-dialog"

func (h *Handlers) PostMessage(ctx *sa.RequestCtx) error {

	resp, err := h.postMessage(ctx)
	if err != nil {
		logger.Error(err)
		errorCase := domain.ApiMessage{
			Code:    fasthttp.StatusPreconditionFailed,
			Message: err.Error(),
		}
		return ctx.HTTPResponse(errorCase.String(), fasthttp.StatusPreconditionFailed)
	}
	return ctx.HTTPResponse(string(resp))
}

func (h *Handlers) postMessage(ctx *sa.RequestCtx) ([]byte, error) {

	// Создать клиента Consul
	cc, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		return nil, err
	}
	// Создать клиента Resty
	rc, err := client.NewDialog(ctx)

	svc, _, err := cc.Health().Service(SERVICE_NAME, "", true, &api.QueryOptions{})
	for _, entry := range svc {
		if SERVICE_NAME != entry.Service.Service {
			continue
		}
		logger.Infof("entry.Service: %s", entry.Service)
		address := fmt.Sprintf("%s:%d", entry.Service.Address, entry.Service.Port)
		resp, err := rc.R().
			SetBody(ctx.PostBody()).
			Post(fmt.Sprintf("http://%s/%s", address, "message"))
		if err != nil {
			return nil, err
		}
		return []byte(resp.String()), nil
	}
	return []byte("{}"), nil
}

func (h *Handlers) GetMessages(ctx *sa.RequestCtx) error {

	resp, err := h.getMessages(ctx)
	if err != nil {
		logger.Error(err)
		errorCase := domain.ApiMessage{
			Code:    fasthttp.StatusPreconditionFailed,
			Message: err.Error(),
		}
		return ctx.HTTPResponse(errorCase.String(), fasthttp.StatusPreconditionFailed)
	}

	return ctx.HTTPResponse(string(resp))
}

func (h *Handlers) getMessages(ctx *sa.RequestCtx) ([]byte, error) {

	// Создать клиента Consul
	cc, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		return nil, err
	}
	// Создать клиента Resty
	rc, err := client.NewDialog(ctx)

	svc, _, err := cc.Health().Service(SERVICE_NAME, "", true, &api.QueryOptions{})
	for _, entry := range svc {
		if SERVICE_NAME != entry.Service.Service {
			continue
		}
		logger.Infof("entry.Service: %s", entry.Service)
		address := fmt.Sprintf("%s:%d", entry.Service.Address, entry.Service.Port)
		resp, err := rc.R().Get(fmt.Sprintf("http://%s/%s", address, "messages"))
		if err != nil {
			return nil, err
		}
		return []byte(resp.String()), nil
	}
	return []byte("{}"), nil
}
