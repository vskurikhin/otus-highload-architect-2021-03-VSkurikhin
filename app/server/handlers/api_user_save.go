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

func (h *Handlers) UserSave(ctx *sa.RequestCtx) error {

	user, err := h.userSave(ctx)

	if err != nil {
		logger.Error(err)
		errorCase := domain.ApiMessage{
			Code:    fasthttp.StatusPreconditionFailed,
			Message: err.Error(),
		}
		return ctx.HTTPResponse(errorCase.String(), fasthttp.StatusPreconditionFailed)
	}
	msg := fmt.Sprintf("updated with id: %s", user.Id().String())
	created := domain.ApiMessage{
		Code:    fasthttp.StatusCreated,
		Message: msg,
	}
	return ctx.HTTPResponse(created.String())
}

type User struct {
	Id        uuid.UUID
	Username  string
	Name      *string
	SurName   *string
	Age       int
	Sex       int
	Interests []string
	City      *string
}

func (h *Handlers) userSave(ctx *sa.RequestCtx) (*domain.User, error) {

	var u User
	err := json.Unmarshal(ctx.PostBody(), &u)

	if err != nil {
		return nil, err
	}
	user := domain.User{Username: u.Username, Name: u.Name, SurName: u.SurName, Age: u.Age, Sex: u.Sex, City: u.City}
	user.SetId(u.Id)

	if logger.DebugEnabled() {
		logger.Debugf("userSave: got user: %s", user.String())
	}
	err = h.Server.DAO.User.Update(&user)

	if err != nil {
		return nil, err
	}
	im, err := h.Server.DAO.Interest.NewInterestMap(u.Interests)

	if err != nil {
		return nil, err
	}
	err = h.Server.DAO.Interest.CreateInterests(im.NewSet)

	interests, err := h.Server.DAO.Interest.NewInterestMap(u.Interests)

	if err != nil {
		return nil, err
	}
	mapInterests, err := h.Server.DAO.UserHasInterests.LinkedInterestMap(user.Id(), interests)

	if err != nil {
		return nil, err
	}
	forLink := mapInterests.ConcatInterests()

	if logger.DebugEnabled() {
		logger.Debugf("userSave: forLink: %v", forLink)
	}
	err = h.Server.DAO.UserHasInterests.LinkInterests(&user, forLink)

	if err != nil {
		return nil, err
	}
	return &user, nil
}
