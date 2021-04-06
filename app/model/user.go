package model

import (
	"encoding/json"
	"github.com/google/uuid"
	"log"
	"strings"
)

type User struct {
	id        uuid.UUID
	Username  string
	Name      *string
	SurName   *string
	Age       int
	Sex       int
	Interests []string
	City      *string
}

func Create(id uuid.UUID, username string, name *string, surname *string, age int, sex int, inters []string, city *string,
) *User {
	return &User{
		id:        id,
		Username:  username,
		Name:      name,
		SurName:   surname,
		Age:       age,
		Sex:       sex,
		Interests: inters,
		City:      city,
	}
}

func New(username string) *User {
	return &User{id: uuid.New(), Username: username}
}

func Setup(username string, name *string, surname *string, age int, sex int, inters []string, city *string) *User {
	return &User{
		Username:  username,
		Name:      name,
		SurName:   surname,
		Age:       age,
		Sex:       sex,
		Interests: inters,
		City:      city,
	}
}

func (u *User) Id() uuid.UUID {
	return u.id
}

func (u *User) SetId(id uuid.UUID) {
	u.id = id
}

func (u *User) StringInterests() string {
	if len(u.Interests) > 0 {
		return `["` + strings.Join(u.Interests, `", "`) + `"]`
	}
	return ""
}

func (u *User) String() string {
	return string(u.Marshal())
}

func (u *User) Marshal() []byte {

	user, err := json.Marshal(struct {
		Id        uuid.UUID
		Username  string
		Name      *string
		SurName   *string
		Age       int
		Sex       int
		Interests []string
		City      *string
	}{
		Id:        u.id,
		Username:  u.Username,
		Name:      u.Name,
		SurName:   u.SurName,
		Age:       u.Age,
		Sex:       u.Sex,
		Interests: u.Interests,
		City:      u.City,
	})
	if err != nil {
		log.Println(err)
		return nil
	}
	return user
}
