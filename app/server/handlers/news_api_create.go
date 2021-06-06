package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	sa "github.com/savsgio/atreugo/v11"
	"github.com/savsgio/go-logger/v2"
	"github.com/valyala/fasthttp"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app/cache"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app/domain"
)

const layout = "1970-01-01T00:00:00.000Z"

func (h *Handlers) CreateNews(ctx *sa.RequestCtx) error {

	news, err := h.createNews(ctx)

	if err != nil {
		logger.Error(err)
		errorCase := domain.ApiMessage{
			Code:    fasthttp.StatusPreconditionFailed,
			Message: err.Error(),
		}
		return ctx.HTTPResponse(errorCase.String(), fasthttp.StatusPreconditionFailed)
	}
	msg := fmt.Sprintf("updated with id: %s", news.Id())
	created := domain.ApiMessage{
		Code:    fasthttp.StatusCreated,
		Message: msg,
	}
	return ctx.HTTPResponse(created.String())
}

func (h *Handlers) createNews(ctx *sa.RequestCtx) (*domain.News, error) {

	_, err := h.profile(ctx)
	if err != nil {
		return nil, err
	}

	var n cache.News
	err = json.Unmarshal(ctx.PostBody(), &n)

	if err != nil {
		return nil, err
	}
	news := domain.ConvertNews(&n)
	news.SetId(uuid.New())
	n.Id = news.Id().String()

	if logger.DebugEnabled() {
		logger.Debugf("createNews: got news: %s", news.String())
	}
	err = h.Server.DAO.News.Create(news)
	if err == nil {
		h.Server.Cache.PutNews(&n)
	}
	h.Server.PubSub.Publish("/ws-newslist", "push")
	return news, nil
}
