package domain

import (
	"encoding/json"
	"github.com/savsgio/go-logger/v2"
)

type User struct {
	id       uint64
	Username string
	Name     *string
	SurName  *string
	Age      int
	Sex      int
	City     *string
	Friend   bool
}

func (u *User) String() string {
	return string(u.Marshal())
}

func (u *User) Marshal() []byte {

	user, err := json.Marshal(*u)
	if err != nil {
		logger.Error(err)
		return nil
	}
	return user
}
