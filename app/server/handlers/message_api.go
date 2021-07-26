package handlers

import (
	"encoding/json"
	"fmt"
	sa "github.com/savsgio/atreugo/v11"
	"github.com/savsgio/go-logger/v2"
	"github.com/valyala/fasthttp"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app/domain"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app/utils"
)

type message struct {
	ToUser  string
	Message string
}

func (h *Handlers) PostMessage(ctx *sa.RequestCtx) error {

	id, err := h.postMessage(ctx)

	if err != nil {
		logger.Error(err)
		errorCase := domain.ApiMessage{
			Code:    fasthttp.StatusPreconditionFailed,
			Message: err.Error(),
		}
		return ctx.HTTPResponse(errorCase.String(), fasthttp.StatusPreconditionFailed)
	}

	return ctx.HTTPResponse(fmt.Sprintf("%d", id))
}

func (h *Handlers) postMessage(ctx *sa.RequestCtx) (*uint64, error) {

	p, err := h.profile(ctx)
	if err != nil {
		return nil, err
	}
	user, err := h.Server.DAO.User.ReadUserById(p.Id)
	if err != nil {
		return nil, err
	}

	var m message
	err = json.Unmarshal(ctx.PostBody(), &m)

	if err != nil {
		return nil, err // правильная обработка ошибок вместо паники
	}
	toUser, err := h.Server.DAO.User.ReadUserByName(m.ToUser)
	if err != nil {
		return nil, err
	}
	shardId := domain.GetShardId(*toUser.City)
	id := utils.RandomIdWithShardId(shardId)
	if logger.DebugEnabled() {
		logger.Debugf("shardId: %d", shardId)
	}
	message := domain.DialogMessage{
		Id:       id,
		ShardId:  shardId,
		FromUser: user.Id,
		ToUser:   toUser.Id,
		Message:  m.Message,
	}
	if logger.DebugEnabled() {
		logger.Debugf("got message: %m", message.String())
	}
	err = h.Server.DAO.DialogMessage.Create(shardId, &message)
	if err != nil {
		return nil, err
	}
	if logger.DebugEnabled() {
		logger.Debugf("got message: %m", message.String())
	}
	return &id, nil
}
