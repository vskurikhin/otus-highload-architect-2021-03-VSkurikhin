package handlers

import (
	"fmt"
	"github.com/google/uuid"
	sa "github.com/savsgio/atreugo/v11"
	"github.com/savsgio/go-logger/v2"
	"github.com/valyala/fasthttp"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app/domain"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app/utils"
	"strings"
)

func (h *Handlers) NewsList(ctx *sa.RequestCtx) error {

	list, err := h.newsList(ctx)

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

func (h *Handlers) newsList(ctx *sa.RequestCtx) ([]string, error) {

	_, err := h.profile(ctx)
	if err != nil {
		return nil, err
	}

	o := fmt.Sprintf("%v", ctx.UserValue("offset"))
	l := fmt.Sprintf("%v", ctx.UserValue("limit"))
	limit, offset, err := parsePairInt(o, l)
	if err != nil {
		return nil, err
	}

	return h.getNewsList(offset, limit)
}

func (h *Handlers) getNewsList(offset, limit int) ([]string, error) {
	var result []string
	intCmd := h.Server.Cache.SCardNews()
	if intCmd.Err() != nil {
		logger.Errorf("SCard error %v", intCmd.Err())
	}
	if intCmd.Val() < 1 {
		logger.Infof("read news from db")
		newsList, err := h.Server.DAO.News.ReadNewsList(offset, limit)
		if err != nil {
			return nil, err
		}

		for _, i := range newsList {
			n := domain.NewsConvert(&i)
			result = append(result, n.String())
			h.Server.Cache.PutNews(n)
		}
	} else {
		logger.Infof("read news from cache")
		keys := h.Server.Cache.SortNewsKeys(int64(offset), int64(limit))
		logger.Debugf("read from cache keys: %s", keys)
		for _, key := range keys {
			result = h.getNews(result, key)
		}
	}
	return result, nil
}

func parsePairInt(o, l string) (int, int, error) {
	offset, err := utils.ParseInt(o)
	if err != nil {
		return -1, 0, err
	}
	limit, err := utils.ParseInt(l)
	if err != nil {
		return offset, 0, err
	}
	return offset, limit, nil
}

func (h *Handlers) getNews(result []string, key string) []string {
	logger.Debugf("read from cache key: %s", key)
	news, err := h.Server.Cache.GetNewsJson(key)
	if err == nil && news != nil {
		result = append(result, *news)
	} else {
		result = h.prepareNewsFromDB(result, key)
	}
	return result
}

func (h *Handlers) prepareNewsFromDB(result []string, key string) []string {
	arr := strings.Split(key, "|")
	logger.Debugf("read from cache arr: %s", arr)
	if len(arr) > 2 {
		u, err := uuid.Parse(arr[2])
		logger.Debugf("read from cache u: %s, err: %s", u, err)
		if err == nil {
			return h.appendNewsFromDB(result, u)
		}
	}
	return result
}

func (h *Handlers) appendNewsFromDB(result []string, id uuid.UUID) []string {
	n, err := h.Server.DAO.News.ReadNews(id)
	logger.Debugf("read from cache n: %s", n)
	if err == nil && n != nil {
		c := domain.NewsConvert(n)
		result = append(result, n.String())
		h.Server.Cache.PutNews(c)
	}
	return result
}
