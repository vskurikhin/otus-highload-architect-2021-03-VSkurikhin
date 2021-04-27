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

func (h *Handlers) Save(ctx *sa.RequestCtx) error {

	user, err := h.save(ctx)

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

func (h *Handlers) save(ctx *sa.RequestCtx) (*domain.User, error) {

	var u User
	err := json.Unmarshal(ctx.PostBody(), &u)

	if err != nil {
		return nil, err
	}
	user := domain.User{Username: u.Username, Name: u.Name, SurName: u.SurName, Age: u.Age, Sex: u.Sex, City: u.City}
	user.SetId(u.Id)

	if logger.DebugEnabled() {
		logger.Debugf("save: got user: %s", user.String())
	}
	err = h.Server.DAO.User.Update(&user)

	if err != nil {
		return nil, err
	}
	is, err := h.Server.DAO.Interest.NewInterestMap(u.Interests)

	if logger.DebugEnabled() {
		logger.Debugf("save: is: %v", is)
	}
	if err != nil {
		return nil, err
	}
	err = h.Server.DAO.Interest.CreateInterests(is.NewSet)

	interests, err := h.Server.DAO.Interest.NewInterestMap(u.Interests)

	if logger.DebugEnabled() {
		logger.Debugf("save: interests: %v", interests)
	}
	if err != nil {
		return nil, err
	}
	mapInterests, err := h.Server.DAO.UserHasInterests.LinkedInterestMap(user.Id(), interests)

	if logger.DebugEnabled() {
		logger.Debugf("save: mapInterests: %v", mapInterests)
	}
	if err != nil {
		return nil, err
	}
	forLink := mapInterests.ConcatInterests()

	if logger.DebugEnabled() {
		logger.Debugf("save: forLink: %v", forLink)
	}
	err = h.Server.DAO.UserHasInterests.LinkInterests(&user, forLink)

	if err != nil {
		return nil, err
	}
	return &user, nil
}