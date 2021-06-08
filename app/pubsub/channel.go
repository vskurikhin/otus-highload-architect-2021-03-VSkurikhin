package pubsub

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/savsgio/go-logger/v2"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app/domain"
)

type Channel struct {
	Name string
}

func CreateChannel(prefix string, p *domain.Profile) *Channel {
	name := fmt.Sprintf("%s/%s", prefix, p.Id)
	return &Channel{Name: name}
}

func CreateChannelById(prefix string, id *uuid.UUID) *Channel {
	name := fmt.Sprintf("%s/%s", prefix, id)
	return &Channel{Name: name}
}

func (m *Channel) Marshal() []byte {

	bytes, err := json.Marshal(*m)
	if err != nil {
		logger.Error(err)
		return nil
	}
	return bytes
}

func (m *Channel) String() string {
	return string(m.Marshal())
}
