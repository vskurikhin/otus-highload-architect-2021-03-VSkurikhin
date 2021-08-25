package handlers

import (
	"encoding/json"
	"fmt"
	sa "github.com/savsgio/atreugo/v11"
	"github.com/savsgio/go-logger/v2"
	"github.com/valyala/fasthttp"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app-dialog/domain"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app-dialog/utils"
	"strings"
)

type message struct {
	ToUser  string
	Message string
}

type readMessage struct {
	id uint64
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
	logger.Infof("%s", ctx.PostBody())

	err = json.Unmarshal(ctx.PostBody(), &m)
	if err != nil {
		return nil, err // правильная обработка ошибок вместо паники
	}

	toUser, err := h.Server.DAO.User.ReadUserByName(m.ToUser)
	if err != nil {
		return nil, err
	}

	shardId := h.Server.Ring.GetId(*toUser.City)
	hashId := h.Server.Ring.GetHashId(*toUser.City)
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
		HashId:   hashId,
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

func (h *Handlers) PutMessage(ctx *sa.RequestCtx) error {

	id, err := h.putMessage(ctx)

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

func (h *Handlers) putMessage(ctx *sa.RequestCtx) (*uint64, error) {

	p, err := h.profile(ctx)
	if err != nil {
		return nil, err
	}

	user, err := h.Server.DAO.User.ReadUserById(p.Id)
	if err != nil {
		return nil, err
	}

	shardId := h.Server.Ring.GetId(*user.City)

	var m readMessage
	logger.Infof("%s", ctx.PostBody())
	err = json.Unmarshal(ctx.PostBody(), &m)
	err = h.Server.DAO.DialogMessage.UpdateReadMessage(m.id, shardId)

	return &m.id, nil
}

func (h *Handlers) GetMessages(ctx *sa.RequestCtx) error {

	list, err := h.getMessages(ctx)

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

func (h *Handlers) getMessages(ctx *sa.RequestCtx) ([]string, error) {
	p, err := h.profile(ctx)

	if err != nil {
		return nil, err
	}
	messages, err := h.getDialogMessages(p.Id)
	if err != nil {
		return nil, err
	}
	var result []string

	for _, message := range messages {
		result = append(result, message.String())
	}
	return result, nil
}

func (h *Handlers) getDialogMessages(id uint64) ([]domain.UserDialogMessage, error) {
	user, err := h.Server.DAO.User.ReadUserById(id)
	if err != nil {
		return nil, err
	}
	shardId := h.Server.Ring.GetId(*user.City)
	messages, err := h.Server.DAO.DialogMessage.GetAllByUserId(shardId, id)

	if err != nil {
		return nil, err
	}
	return messages, nil
}
