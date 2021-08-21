package handlers

import (
	"fmt"
	sa "github.com/savsgio/atreugo/v11"
	"github.com/savsgio/go-logger/v2"
	"github.com/valyala/fasthttp"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app-main/domain"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app-main/security"
	"strings"
)

func (h *Handlers) UserSearch(ctx *sa.RequestCtx) error {

	list, err := h.userSearch(ctx)

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

func (h *Handlers) userSearch(ctx *sa.RequestCtx) ([]string, error) {

	p, err := h.profile(ctx)

	if err != nil {
		return nil, err
	}
	name := fmt.Sprintf("%v", ctx.UserValue("name"))
	surname := fmt.Sprintf("%v", ctx.UserValue("surname"))
	err = security.CheckValue("name", name)

	if err != nil {
		return nil, err
	}
	err = security.CheckValue("surname", surname)

	if err != nil {
		return nil, err
	}
	if logger.DebugEnabled() {
		logger.Debugf("got user name: %s, surname: %s, profile", name, surname, p)
	}

	if err != nil {
		return nil, err
	}
	users, err := h.Server.DAO.User.SearchUserList(p.Id, name, surname)

	if err != nil {
		return nil, err
	}
	var result []string

	for _, user := range users {
		result = append(result, user.String())
	}
	return result, nil
}

func (h *Handlers) SearchBy(ctx *sa.RequestCtx) error {

	list, err := h.searchBy(ctx)

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

func (h *Handlers) searchBy(ctx *sa.RequestCtx) ([]string, error) {

	p, err := h.profile(ctx)

	if err != nil {
		return nil, err
	}
	field := fmt.Sprintf("%v", ctx.UserValue("field"))
	value := fmt.Sprintf("%v", ctx.UserValue("value"))
	err = security.CheckValue("field", field)

	if err != nil {
		return nil, err
	}
	err = security.CheckValue("value", value)

	if err != nil {
		return nil, err
	}
	if logger.DebugEnabled() {
		logger.Debugf("got user field: %s, value: %s, profile", field, value, p)
	}

	if err != nil {
		return nil, err
	}
	users, err := h.Server.DAO.User.SearchByUserList(p.Id, field, value)

	if err != nil {
		return nil, err
	}
	var result []string

	for _, user := range users {
		result = append(result, user.String())
	}
	return result, nil
}
