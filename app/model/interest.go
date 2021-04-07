package model

import (
	"encoding/json"
	"github.com/google/uuid"
	"log"
)

type Interest struct {
	id        uuid.UUID
	Interests string
}

func (i *Interest) Id() uuid.UUID {
	return i.id
}

func (i *Interest) SetId(id uuid.UUID) {
	i.id = id
}

func (i *Interest) String() string {
	return string(i.Marshal())
}

func (i *Interest) Marshal() []byte {

	interest, err := json.Marshal(*i)
	if err != nil {
		log.Println(err)
		return nil
	}
	return interest
}
