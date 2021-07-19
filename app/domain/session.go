package domain

import (
	"encoding/json"
	"github.com/savsgio/go-logger/v2"
)

type Session struct {
	Id        uint64
	SessionId uint64
}

func (s *Session) String() string {
	return string(s.Marshal())
}

func (s *Session) Marshal() []byte {

	session, err := json.Marshal(*s)
	if err != nil {
		logger.Error(err)
		return nil
	}
	return session
}
