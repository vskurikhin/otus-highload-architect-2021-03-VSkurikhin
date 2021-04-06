package message

import (
	"encoding/json"
	"log"
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
		log.Println(err)
		return nil
	}
	return token
}
