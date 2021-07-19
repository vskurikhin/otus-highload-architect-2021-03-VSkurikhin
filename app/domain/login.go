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
