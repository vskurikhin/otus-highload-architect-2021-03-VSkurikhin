package handlers

import (
	"errors"
	"github.com/go-resty/resty/v2"
	sa "github.com/savsgio/atreugo/v11"
	"github.com/savsgio/go-logger/v2"
	"github.com/valyala/fasthttp"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app-main/config"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app-main/domain"
	"net/http"
)

func (h *Handlers) PostMessage(ctx *sa.RequestCtx) error {

	jwtCookie := ctx.Request.Header.Cookie(config.ACCESS_TOKEN_COOKIE)

	if len(jwtCookie) == 0 {
		return errors.New(" JWT Cookie is empty ")
	}

	// Создать клиента Resty
	client := resty.New()

	// Файлы cookie для всех запросов
	client.SetCookie(&http.Cookie{
		Name:     config.ACCESS_TOKEN_COOKIE,
		Value:    string(jwtCookie),
		Path:     "/",
		Domain:   "localhost",
		MaxAge:   36000,
		HttpOnly: true,
		Secure:   false,
	})
	resp, err := client.R().
		EnableTrace().
		SetBody(ctx.PostBody()).
		Post("http://localhost:8079/message")
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

	jwtCookie := ctx.Request.Header.Cookie(config.ACCESS_TOKEN_COOKIE)

	if len(jwtCookie) == 0 {
		return errors.New(" JWT Cookie is empty ")
	}

	// Создать клиента Resty
	client := resty.New()

	// Файлы cookie для всех запросов
	client.SetCookie(&http.Cookie{
		Name:     config.ACCESS_TOKEN_COOKIE,
		Value:    string(jwtCookie),
		Path:     "/",
		Domain:   "localhost",
		MaxAge:   36000,
		HttpOnly: true,
		Secure:   false,
	})
	resp, err := client.R().Get("http://localhost:8079/messages")
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
