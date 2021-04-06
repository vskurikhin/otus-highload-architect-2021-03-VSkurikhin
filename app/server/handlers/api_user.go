package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	sa "github.com/savsgio/atreugo/v11"
	"github.com/savsgio/go-logger/v2"
	"github.com/valyala/fasthttp"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app/message"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app/model"
	"strings"
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

	rows, err := stmtOut.Query()
	if err != nil {
		logger.Errorf("%v", err.Error())
		return "{}", err
	}
	defer rows.Close()

	var users []string
	for rows.Next() {
		var id uuid.UUID
		var u model.User
		var interests string

		err = rows.Scan(&id, &u.Username, &u.Name, &u.SurName, &u.Age, &u.Sex, &interests, &u.City)
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
		users = append(users, u.String())
	}
	return "[" + strings.Join(users, ", ") + "]", nil
}

func (h *Handlers) Create(ctx *sa.RequestCtx) error {

	var user model.User
	err := json.Unmarshal(ctx.PostBody(), &user)
	user.SetId(uuid.New())

	if err != nil {
		errorCase := message.ApiMessage{
			Code:    fasthttp.StatusPreconditionFailed,
			Message: err.Error(),
		}
		return ctx.HTTPResponse(errorCase.String(), fasthttp.StatusPreconditionFailed)
	}
	if logger.DebugEnabled() {
		logger.Debugf("got: %s", user.String())
	}
	err = create(h.Server.DB, &user)
	if err != nil {
		errorCase := message.ApiMessage{
			Code:    fasthttp.StatusPreconditionFailed,
			Message: err.Error(),
		}
		return ctx.HTTPResponse(errorCase.String(), fasthttp.StatusPreconditionFailed)
	}
	msg := fmt.Sprintf("created with id: %s", user.Id().String())
	created := message.ApiMessage{
		Code:    fasthttp.StatusCreated,
		Message: msg,
	}
	return ctx.HTTPResponse(created.String())
}

func create(db *sql.DB, user *model.User) error {

	// Prepare statement for inserting data
	stmtIns, err := db.Prepare(`
		INSERT INTO users
		    (id, username, password, name, surname, age, sex, interests, city)
		  VALUES 
		    (?, ?, 'none', ?, ?, 51, 1, ?, ?)`) // ? = placeholder
	if err != nil {
		return err // proper error handling instead of panic in your app
	}
	defer func() { _ = stmtIns.Close() }() // Close the statement when we leave main() / the program terminates

	id, err := user.Id().MarshalBinary()

	if err != nil {
		return err // proper error handling instead of panic in your app
	}
	interests, err := json.Marshal(user.Interests)

	_, err = stmtIns.Exec(id, user.Username, user.Name, user.SurName, interests, user.City)
	if err != nil {
		return err // proper error handling instead of panic in your app
	}
	return nil
}
