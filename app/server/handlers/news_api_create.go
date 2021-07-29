package handlers

import (
	"encoding/json"
	"fmt"
	sa "github.com/savsgio/atreugo/v11"
	"github.com/savsgio/go-logger/v2"
	"github.com/valyala/fasthttp"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app/cache"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app/domain"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app/pubsub"
	"strings"
	"time"
)

const layout = "1970-12-30 00:00:00"

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

	p, err := h.profile(ctx)
	if err != nil {
		return nil, err
	}

	var n cache.News
	err = json.Unmarshal(ctx.PostBody(), &n)
	n.Username = p.Username

	if err != nil {
		return nil, err
	}
	news := domain.ConvertNews(&n)
	news.NewId()
	n.Id = news.Id().String()

	if logger.DebugEnabled() {
		logger.Debugf("createNews: got news: %s", news.String())
	}
	err = h.Server.DAO.News.Create(news)
	if err != nil {
		return nil, err
	}
	rfc3339t := strings.Replace(news.PublicAt, " ", "T", 1) + "Z"
	t, err := time.Parse(time.RFC3339, rfc3339t)
	if err != nil {
		return nil, err
	}
	key := cache.CreateNewsKey(&t, news.Id())
	message := pubsub.CreateMessage(1, key.Key)
	channel := pubsub.CreateChannel("news", p)
	h.Server.Cache.PutNews(p.Cache(), key, &n)
	h.Server.PubSub.Publish(channel.Name, message)
	return news, nil
}
