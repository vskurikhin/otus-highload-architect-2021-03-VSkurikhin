package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	sa "github.com/savsgio/atreugo/v11"
	"github.com/savsgio/go-logger/v2"
	"github.com/valyala/fasthttp"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app/domain"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app/security"
	"strconv"
	"strings"
)

func (h *Handlers) List(ctx *sa.RequestCtx) error {

	u, err := h.Server.DAO.User.ReadListAsString()

	if err != nil {
		logger.Error(err)
		errorCase := domain.ApiMessage{
			Code:    fasthttp.StatusPreconditionFailed,
			Message: err.Error(),
		}
		return ctx.HTTPResponse(errorCase.String(), fasthttp.StatusPreconditionFailed)
	}
	return ctx.HTTPResponse(u)
}

func (h *Handlers) User(ctx *sa.RequestCtx) error {

	id := fmt.Sprintf("%v", ctx.UserValue("id"))
	if logger.DebugEnabled() {
		logger.Debugf("got user id: %s", id)
	}
	u, err := h.user(ctx)

	if err != nil {
		logger.Error(err)
		errorCase := domain.ApiMessage{
			Code:    fasthttp.StatusPreconditionFailed,
			Message: err.Error(),
		}
		return ctx.HTTPResponse(errorCase.String(), fasthttp.StatusPreconditionFailed)
	}
	return ctx.HTTPResponse(u.String())
}

func (h *Handlers) user(ctx *sa.RequestCtx) (*domain.User, error) {

	id := fmt.Sprintf("%v", ctx.UserValue("id"))

	if logger.DebugEnabled() {
		logger.Debugf("got user id: %s", id)
	}
	u, err := h.Server.DAO.User.ReadUser(id)

	if err != nil {
		return nil, err
	}
	p, _ := h.profile(ctx)

	if p != nil {
		f, _ := h.Server.DAO.UserHasFriends.IsFriendship(p.Id, u.Id())
		u.Friend = f
	}
	return u, nil
}

func (h *Handlers) SignIn(ctx *sa.RequestCtx) error {

	token, err := h.signIn(ctx)

	if err != nil {
		logger.Error(err)
		errorCase := domain.ApiMessage{
			Code:    fasthttp.StatusPreconditionFailed,
			Message: err.Error(),
		}
		return ctx.HTTPResponse(errorCase.String(), fasthttp.StatusPreconditionFailed)
	}

	return ctx.HTTPResponse(token.String())
}

func (h *Handlers) signIn(ctx *sa.RequestCtx) (*domain.Token, error) {

	var s domain.Signin
	err := json.Unmarshal(ctx.PostBody(), &s)

	if err != nil {
		return nil, err // правильная обработка ошибок вместо паники
	}
	err = security.CheckSignIn(&s)

	if err != nil {
		return nil, err
	}

	if logger.DebugEnabled() {
		logger.Debugf("got s: %s", s.String())
	}
	id := uuid.New()
	password := security.HashAndSalt([]byte(s.Password))
	login := domain.Login{Username: s.Username}
	login.SetId(id)
	login.SetPassword(password)
	err = h.Server.DAO.Login.Create(&login)

	if err != nil {
		return nil, err
	}
	age, err := strconv.ParseInt(s.Age, 10, 64)

	if err != nil {
		return nil, err
	}
	sex, err := strconv.ParseInt(s.Sex, 10, 64)

	if err != nil {
		return nil, err
	}
	ins := strings.Split(s.Interests, "\n")
	user := domain.Create(id, s.Username, &s.Name, &s.Surname, int(age), int(sex), ins, &s.City, true)

	if logger.DebugEnabled() {
		logger.Debugf("got user: %s", user.String())
	}
	err = h.Server.DAO.User.Create(user)

	if err != nil {
		return nil, err
	}
	_ = h.Server.DAO.Interest.CreateInterests(ins)
	interests, _ := h.Server.DAO.Interest.GetExistsInterests(ins)
	_ = h.Server.DAO.UserHasInterests.LinkInterests(user, interests)

	return h.generateToken(ctx, id), nil
}

func (h *Handlers) Create(ctx *sa.RequestCtx) error {

	var user domain.User
	err := json.Unmarshal(ctx.PostBody(), &user)

	if err != nil {
		logger.Error(err)
		errorCase := domain.ApiMessage{
			Code:    fasthttp.StatusPreconditionFailed,
			Message: err.Error(),
		}
		return ctx.HTTPResponse(errorCase.String(), fasthttp.StatusPreconditionFailed)
	}
	user.SetId(uuid.New())

	if logger.DebugEnabled() {
		logger.Debugf("got user: %s", user.String())
	}
	err = h.Server.DAO.User.Create(&user)

	if err != nil {
		logger.Error(err)
		errorCase := domain.ApiMessage{
			Code:    fasthttp.StatusPreconditionFailed,
			Message: err.Error(),
		}
		return ctx.HTTPResponse(errorCase.String(), fasthttp.StatusPreconditionFailed)
	}
	msg := fmt.Sprintf("created with id: %s", user.Id().String())
	created := domain.ApiMessage{
		Code:    fasthttp.StatusCreated,
		Message: msg,
	}
	return ctx.HTTPResponse(created.String())
}

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
