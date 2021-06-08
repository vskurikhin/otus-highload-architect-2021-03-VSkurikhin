package domain

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/savsgio/go-logger/v2"
)

type WsMessage struct {
	Code    int
	Id      uuid.UUID
	Message string
}

func (w *WsMessage) String() string {
	return string(w.Marshal())
}

func (w *WsMessage) Marshal() []byte {

	apiMessage, err := json.Marshal(*w)
	if err != nil {
		logger.Errorf("%v", err)
		return nil
	}
	return apiMessage
}
