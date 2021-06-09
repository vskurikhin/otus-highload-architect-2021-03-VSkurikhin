package cache

import (
	"encoding/json"
	"fmt"
	"github.com/savsgio/go-logger/v2"
)

type SetKey struct {
	Key string
}

func CreateSetKey(p *Profile) *SetKey {
	key := fmt.Sprintf("News|%s", p.Id)
	return &SetKey{Key: key}
}

func (n *SetKey) Marshal() []byte {

	bytes, err := json.Marshal(*n)
	if err != nil {
		logger.Error(err)
		return nil
	}
	return bytes
}

func (n *SetKey) String() string {
	return string(n.Marshal())
}
