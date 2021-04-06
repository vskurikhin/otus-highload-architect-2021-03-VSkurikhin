package handlers

import (
	"database/sql"
	"encoding/json"
	"github.com/google/uuid"
	sa "github.com/savsgio/atreugo/v11"
	"github.com/savsgio/go-logger/v2"
	"github.com/valyala/fasthttp"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app/message"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app/model"
)

func (h *Handlers) List(ctx *sa.RequestCtx) error {

	u, err := list(h.Server.DB)

	if err != nil {
		errorCase := message.ApiMessage{
			Code:    fasthttp.StatusPreconditionFailed,
			Message: err.Error(),
		}
		return ctx.HTTPResponse(errorCase.String(), fasthttp.StatusPreconditionFailed)
	}

	return ctx.HTTPResponse(u)
}

func list(db *sql.DB) (string, error) {

	stmtOut, err := db.Prepare(`
		SELECT id, username, name, surname, age, sex, interests, city FROM users`)
	if err != nil {
		logger.Errorf("%v", err.Error())
		return "{}", err
	}
	defer func() { _ = stmtOut.Close() }()

	var id uuid.UUID
	var u model.User
	var interests string
	err = stmtOut.QueryRow().
		Scan(&id, &u.Username, &u.Name, &u.SurName, &u.Age, &u.Sex, &interests, &u.City)
	if err != nil {
		logger.Errorf("%v", err.Error())
		return "{}", err
	}
	u.SetId(id)
	err = json.Unmarshal([]byte(interests), &u.Interests)
	if err != nil {
		logger.Errorf("%v", err.Error())
		return "{}", err
	}
	return u.String(), nil
}
