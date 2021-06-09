package pubsub

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/savsgio/go-logger/v2"
	"time"
)

type Message struct {
	Id      *uuid.UUID
	Code    int
	Message string
	SendAt  *time.Time
}

func CreateMessage(code int, msg string) *Message {
	id := uuid.New()
	now := time.Now()
	return &Message{Id: &id, Code: code, Message: msg, SendAt: &now}
}

func (m *Message) Marshal() []byte {

	bytes, err := json.Marshal(*m)
	if err != nil {
		logger.Error(err)
		return nil
	}
	return bytes
}

func (m *Message) String() string {
	return string(m.Marshal())
}
