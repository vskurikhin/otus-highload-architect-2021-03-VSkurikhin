package cache

import (
	"encoding/json"
	"github.com/savsgio/go-logger/v2"
)

type News struct {
	Id       string
	Title    string
	Content  string
	PublicAt string
	Username string
}

func (n *News) Marshal() []byte {

	bytes, err := json.Marshal(*n)
	if err != nil {
		logger.Error(err)
		return nil
	}
	return bytes
}

func (n *News) String() string {
	return string(n.Marshal())
}
