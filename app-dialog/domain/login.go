package domain

import (
	"encoding/json"
	"github.com/savsgio/go-logger/v2"
)

type Login struct {
	Id       uint64
	Username string
	Password string
}

func (l *Login) String() string {
	return string(l.Marshal())
}

func (l *Login) Marshal() []byte {

	login, err := json.Marshal(*l)
	if err != nil {
		logger.Error(err)
		return nil
	}
	return login
}

const INSERT_INTO_LOGIN_USERNAME_PASSWORD = `
	INSERT INTO login (username, password) VALUES (?, ?)`

func (l *login) Create(login *Login) (*Login, error) {
	// Подготовить оператор для вставки данных
	stmtIns, err := l.dbRw.Prepare(INSERT_INTO_LOGIN_USERNAME_PASSWORD) // ? = заполнитель

	if err != nil {
		return login, err // правильная обработка ошибок вместо паники
	}
	defer func() { _ = stmtIns.Close() }() // Закрывается оператор, когда выйдете из функции

	res, err := stmtIns.Exec(login.Username, login.Password)
	if err != nil {
		return login, err
	}
	id, err := (res.LastInsertId())
	if err != nil {
		return login, err
	}
	login.Id = uint64(id)

	return login, nil
}

const SELECT_ID_USERNAME_PASSWORD_FROM_LOGIN = `
	SELECT id, username, password FROM login WHERE username = ?`

func (l *login) Read(username string) (*Login, error) {

	stmtOut, err := l.dbRw.Prepare(SELECT_ID_USERNAME_PASSWORD_FROM_LOGIN)

	if err != nil {
		return nil, err // правильная обработка ошибок вместо паники
	}
	defer func() { _ = stmtOut.Close() }() // Закрывается оператор, когда выйдете из функции

	var login Login
	err = stmtOut.QueryRow(username).Scan(&login.Id, &login.Username, &login.Password)
	if err != nil {
		return nil, err
	}
	return &login, nil
}
