package client

import (
	"errors"
	"github.com/go-resty/resty/v2"
	sa "github.com/savsgio/atreugo/v11"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app-main/config"
	"net/http"
)

func NewDialog(ctx *sa.RequestCtx) (*resty.Client, error) {
	// Создать клиента Resty
	client := resty.New()

	jwtCookie := ctx.Request.Header.Cookie(config.ACCESS_TOKEN_COOKIE)

	if len(jwtCookie) == 0 {
		return nil, errors.New(" JWT Cookie is empty ")
	}

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

	return client, nil
}
