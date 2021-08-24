package dto

import (
	"encoding/json"
	"github.com/savsgio/go-logger/v2"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app-counter/domain"
	"strconv"
)

type KafkaUser struct {
	Id       string
	Username string
}

func (u *KafkaUser) String() string {
	return string(u.Marshal())
}

func (u *KafkaUser) Marshal() []byte {

	user, err := json.Marshal(*u)
	if err != nil {
		logger.Error(err)
		return nil
	}
	return user
}

func (u *KafkaUser) User() *domain.User {
	id, err := strconv.ParseUint(u.Id, 10, 64)
	if err != nil {
		logger.Error(err)
		return nil
	}
	return &domain.User{
		Id:       id,
		Username: u.Username,
	}
}

func UnmarshalKafkaUser(bytes []byte) (*KafkaUser, error) {

	var value KafkaUser
	err := json.Unmarshal(bytes, &value)

	if logger.DebugEnabled() {
		logger.Debugf("bytes: %v", value)
	}
	if err != nil {
		return nil, err
	}
	return &value, nil
}
