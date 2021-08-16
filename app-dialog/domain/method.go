package domain

import (
	"encoding/json"
	"github.com/savsgio/go-logger/v2"
)

type Method struct {
	Method string
	Key    string
	Offset int
	Limit  int
}

func (m *Method) String() string {
	return string(m.Marshal())
}

func (m *Method) Marshal() []byte {

	apiMessage, err := json.Marshal(*m)
	if err != nil {
		logger.Errorf("%v", err)
		return nil
	}
	return apiMessage
}
