package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	sa "github.com/savsgio/atreugo/v11"
	"github.com/savsgio/go-logger/v2"
	"github.com/valyala/fasthttp"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app/domain"
)

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

type news struct {
	Title    string
	Content  string
	PublicAt string
}

func (h *Handlers) createNews(ctx *sa.RequestCtx) (*domain.News, error) {

	_, err := h.profile(ctx)
	if err != nil {
		return nil, err
	}

	var n news
	err = json.Unmarshal(ctx.PostBody(), &n)

	if err != nil {
		return nil, err
	}
	news := domain.News{Title: n.Title, Content: n.Content, PublicAt: n.PublicAt}
	news.SetId(uuid.New())

	if logger.DebugEnabled() {
		logger.Debugf("createNews: got news: %s", news.String())
	}
	err = h.Server.DAO.News.Create(&news)
	if err == nil {
		updateCache(news)
	}

	return &news, nil
}

func updateCache(n domain.News) {

}
