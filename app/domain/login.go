package domain

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/savsgio/go-logger/v2"
)

type Login struct {
	id       uuid.UUID
	Username string
	password string
}

func (l *Login) Id() uuid.UUID {
	return l.id
}

func (l *Login) SetId(id uuid.UUID) {
	l.id = id
}

func (l *Login) Password() string {
	return l.password
}

func (l *Login) SetPassword(password string) {
	l.password = password
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

func (l *login) Create(login *Login) error {
	// Подготовить оператор для вставки данных
	stmtIns, err := l.dbRw.Prepare("INSERT INTO login (id, username, password) VALUES (?, ?, ?)") // ? = заполнитель

	if err != nil {
		return err // правильная обработка ошибок вместо паники
	}
	defer func() { _ = stmtIns.Close() }() // Закрывается оператор, когда выйдете из функции

	id, err := login.Id().MarshalBinary()
	_, err = stmtIns.Exec(id, login.Username, login.Password())
	if err != nil {
		return err
	}
	return nil
}

func (l *login) Read(username string) (*Login, error) {

	stmtOut, err := l.dbRw.Prepare(`SELECT id, username, password FROM login WHERE username = ?`)

	if err != nil {
		return nil, err // правильная обработка ошибок вместо паники
	}
	defer func() { _ = stmtOut.Close() }() // Закрывается оператор, когда выйдете из функции

	var login Login
	err = stmtOut.QueryRow(username).
		Scan(&login.id, &login.Username, &login.password)
	if err != nil {
		return nil, err
	}

	return &login, nil
}
