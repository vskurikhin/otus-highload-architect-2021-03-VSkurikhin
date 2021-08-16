package cache

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/savsgio/go-logger/v2"
	"time"
)

type NewsKey struct {
	Key string
}

func CreateNewsKey(t *time.Time, id *uuid.UUID) *NewsKey {
	key := fmt.Sprintf("News|%s|%s", t.Format(time.RFC3339), id)
	return &NewsKey{Key: key}
}

func (n *NewsKey) Marshal() []byte {

	bytes, err := json.Marshal(*n)
	if err != nil {
		logger.Error(err)
		return nil
	}
	return bytes
}

func (n *NewsKey) String() string {
	return string(n.Marshal())
}
