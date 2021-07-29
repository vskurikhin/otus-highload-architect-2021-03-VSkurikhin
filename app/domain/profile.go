package domain

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/savsgio/go-logger/v2"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app/cache"
)

type Profile struct {
	Id       *uuid.UUID
	Username string
}

func (p *Profile) Cache() *cache.Profile {
	return &cache.Profile{Id: p.Id, Username: p.Username}
}

func (p *Profile) String() string {
	return string(p.Marshal())
}

func (p *Profile) Marshal() []byte {

	token, err := json.Marshal(*p)
	if err != nil {
		logger.Error(err)
		return nil
	}
	return token
}
