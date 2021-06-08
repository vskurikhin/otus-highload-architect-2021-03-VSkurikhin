package handlers

import (
	"fmt"
	sa "github.com/savsgio/atreugo/v11"
	"github.com/savsgio/go-logger/v2"
	"github.com/valyala/fasthttp"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app/cache"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app/domain"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app/utils"
	"strings"
	"time"
)

func (h *Handlers) NewsList(ctx *sa.RequestCtx) error {

	list, err := h.newsList(ctx)
	if err != nil {
		logger.Errorf("NewsList error %v", err)
	}

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

	p, err := h.profile(ctx)
	if err != nil {
		return nil, err
	}

	o := fmt.Sprintf("%v", ctx.UserValue("offset"))
	l := fmt.Sprintf("%v", ctx.UserValue("limit"))
	limit, offset, err := parsePairInt(o, l)
	if err != nil {
		return nil, err
	}

	return h.getNewsList(p, offset, limit)
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

func parseDate(value string) (time.Time, error) {
	layout := time.RFC3339[:len(value)]
	return time.Parse(layout, value)
}

func (h *Handlers) getNewsList(p *domain.Profile, offset, limit int) ([]string, error) {

	sk := cache.CreateSetKey(p.Cache())
	sizeInCache, _ := h.Server.Cache.SCardNews(sk)
	sizeInDB, _ := h.Server.DAO.News.SizeMyNewsList(p, offset, limit)
	if sizeInCache != sizeInDB {
		logger.Debugf("sizeInCache != sizeInDB -> %d != %d", sizeInCache, sizeInDB)
		// Update Cache
		return h.updateCache(p, offset, limit)
	}
	lastNews, err := h.Server.DAO.News.LastMyNews(p)
	if err != nil {
		logger.Errorf("getNewsList error %v lastNews", err)
		// Update Cache
		return h.updateCache(p, offset, limit)
	}
	logger.Debugf("lastNews %s", lastNews)
	rfc3339t := strings.Replace(lastNews.PublicAt, " ", "T", 1) + "Z"
	logger.Debugf("getNewsList rfc3339t: %s", rfc3339t)
	t, err := time.Parse(time.RFC3339, rfc3339t)
	if err != nil {
		logger.Errorf("getNewsList lastNews.PublicAt: %s error %v", lastNews.PublicAt, err)
		// Update Cache
		return h.updateCache(p, offset, limit)
	}
	lastNewsKey, err := h.Server.Cache.LastMyNewsKey(p.Cache())
	if err != nil {
		logger.Errorf("getNewsList lastNewsKey error %v", err)
		// Update Cache
		return h.updateCache(p, offset, limit)
	}
	nk := cache.CreateNewsKey(&t, lastNews.Id())
	if nk.Key != lastNewsKey.Key {
		logger.Debugf("nk.Key != lastNewsKey.Key -> %s != %s", nk.Key, lastNewsKey.Key)
		// Update Cache
		return h.updateCache(p, offset, limit)
	}
	// Read from Cache
	return h.readFromCache(p, offset, limit), nil
}

func (h *Handlers) updateCache(p *domain.Profile, offset, limit int) ([]string, error) {

	var result []string
	var cacheUpdateError error
	newsList, err := h.Server.DAO.News.ReadMyNewsList(p, offset, limit)
	if err != nil {
		cacheUpdateError = err
	}
	for _, i := range newsList {
		n := domain.NewsConvert(&i)
		result = append(result, n.String())
		rfc3339t := strings.Replace(i.PublicAt, " ", "T", 1) + "Z"
		t, err := time.Parse(time.RFC3339, rfc3339t)
		if err != nil {
			cacheUpdateError = err
		}
		key := cache.CreateNewsKey(&t, i.Id())
		h.Server.Cache.PutNews(p.Cache(), key, n)
	}
	return result, cacheUpdateError
}

func (h *Handlers) readFromCache(p *domain.Profile, offset int, limit int) []string {

	var result []string
	keys := h.Server.Cache.SortMyNewsKeys(p.Cache(), int64(offset), int64(limit))
	logger.Debugf("read from cache keys: %s", keys)
	for _, key := range keys {
		result = h.getNews(result, key)
	}
	return result
}

func (h *Handlers) getNews(result []string, key string) []string {

	logger.Debugf("read from cache key: %s", key)
	news, err := h.Server.Cache.GetNewsJson(key)
	if err == nil && news != nil {
		result = append(result, *news)
	} else {
		logger.Errorf("handlers.getNews for key %s error: %s", key, err)
	}
	return result
}
