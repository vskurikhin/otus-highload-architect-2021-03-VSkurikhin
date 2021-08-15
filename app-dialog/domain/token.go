package domain

import (
	"encoding/json"
	"github.com/savsgio/go-logger/v2"
)

type Token struct {
	Token string `json:"token"`
}

func (t *Token) String() string {
	return string(t.Marshal())
}

func (t *Token) Marshal() []byte {

	token, err := json.Marshal(*t)
	if err != nil {
		logger.Error(err)
		return nil
	}
	return token
}
