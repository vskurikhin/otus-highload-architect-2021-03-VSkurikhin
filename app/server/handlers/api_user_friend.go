package handlers

import (
	"encoding/json"
	"fmt"
	sa "github.com/savsgio/atreugo/v11"
	"github.com/savsgio/go-logger/v2"
	"github.com/valyala/fasthttp"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app/domain"
)

func (h *Handlers) Friend(ctx *sa.RequestCtx) error {

	friend, err := h.friend(ctx)

	if err != nil {
		logger.Error(err)
		errorCase := domain.ApiMessage{
			Code:    fasthttp.StatusPreconditionFailed,
			Message: err.Error(),
		}
		return ctx.HTTPResponse(errorCase.String(), fasthttp.StatusPreconditionFailed)
	}
	msg := fmt.Sprintf("linked user id: %s and friend id: %s", friend.UserId, friend.FriendId)
	created := domain.ApiMessage{
		Code:    fasthttp.StatusCreated,
		Message: msg,
	}
	return ctx.HTTPResponse(created.String())
}

func (h *Handlers) friend(ctx *sa.RequestCtx) (*domain.Friend, error) {

	var friend domain.Friend
	err := json.Unmarshal(ctx.PostBody(), &friend)

	if err != nil {
		return nil, err
	}
	if logger.DebugEnabled() {
		logger.Debugf("got user id: %s", friend.UserId)
	}
	u, err := h.Server.DAO.User.ReadUser(friend.UserId)

	if err != nil {
		return nil, err
	}
	if logger.DebugEnabled() {
		logger.Debugf("got friend id: %s", friend.FriendId)
	}
	f, err := h.Server.DAO.User.ReadUser(friend.FriendId)
	if err != nil {
		return nil, err
	}

	err = h.Server.DAO.UserHasFriends.LinkToFriend(u, f)

	return &friend, nil
}
