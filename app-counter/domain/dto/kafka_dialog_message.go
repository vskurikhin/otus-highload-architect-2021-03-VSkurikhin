package dto

import (
	"encoding/json"
	"github.com/savsgio/go-logger/v2"
)

type KafkaDialogMessage struct {
	Id           string
	To_user      uint64
	Already_read int
	Operation    string `json:"_operation"`
}

func (u *KafkaDialogMessage) String() string {
	return string(u.Marshal())
}

func (u *KafkaDialogMessage) Marshal() []byte {

	user, err := json.Marshal(*u)
	if err != nil {
		logger.Error(err)
		return nil
	}
	return user
}

func UnmarshalKafkaDialogMessage(bytes []byte) (*KafkaDialogMessage, error) {

	var value KafkaDialogMessage
	err := json.Unmarshal(bytes, &value)

	if logger.DebugEnabled() {
		logger.Debugf("bytes: %v", value)
	}
	if err != nil {
		return nil, err
	}
	return &value, nil
}
