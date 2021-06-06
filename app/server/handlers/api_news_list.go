package handlers

import (
	"fmt"
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

	offsetString := fmt.Sprintf("%v", ctx.UserValue("offset"))
	offset, err := utils.ParseInt(offsetString)
	if err != nil {
		return nil, err
	}
	limitString := fmt.Sprintf("%v", ctx.UserValue("limit"))
	limit, err := utils.ParseInt(limitString)
	if err != nil {
		return nil, err
	}
	newsList, err := h.Server.DAO.News.ReadNewsList(offset, limit)
	if err != nil {
		return nil, err
	}
	var result []string

	for _, user := range newsList {
		result = append(result, user.String())
	}
	return result, nil
}
