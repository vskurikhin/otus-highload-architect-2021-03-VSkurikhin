package model

import (
	"encoding/json"
	"github.com/google/uuid"
	"log"
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
		log.Println(err)
		return nil
	}
	return login
}
