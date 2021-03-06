package handlers

import (
	"errors"
	sa "github.com/savsgio/atreugo/v11"
	"github.com/savsgio/go-logger/v2"
	"github.com/valyala/fasthttp"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app-dialog/config"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app-dialog/domain"
	"strconv"
)

func (h *Handlers) Profile(ctx *sa.RequestCtx) error {

	user, err := h.profile(ctx)

	if err != nil {

		logger.Errorf("Not found session. %v", err)
		errorCase := domain.ApiMessage{
			Code:    fasthttp.StatusForbidden,
			Message: "Not found session.",
		}
		return ctx.HTTPResponse(errorCase.String(), fasthttp.StatusForbidden)
	}
	return ctx.HTTPResponse(user.String())
}

func (h *Handlers) profile(ctx *sa.RequestCtx) (*domain.Profile, error) {

	jwtCookie := ctx.Request.Header.Cookie(config.ACCESS_TOKEN_COOKIE)

	if len(jwtCookie) == 0 {
		return nil, errors.New(" JWT Cookie is empty ")
	}
	psid, err := h.Server.JWT.SessionIdFromToken(string(jwtCookie))

	if err != nil {
		return nil, err
	}
	if psid == nil {
		return nil, errors.New(" session id is empty ")
	}
	sessionId, err := strconv.ParseUint(*psid, 10, 64)
	logger.Debugf("sessionId: %d", sessionId)
	if err != nil {
		return nil, err
	}
	return h.GetProfile(sessionId)
}

func (h *Handlers) GetProfile(sessionId uint64) (*domain.Profile, error) {

	profile, err := h.Server.DAO.Session.ProfileBySessionId(sessionId)
	if err != nil {
		return nil, err
	}
	return profile, nil
}
