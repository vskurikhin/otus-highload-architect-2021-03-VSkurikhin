package domain

import (
	"encoding/json"
	"github.com/savsgio/go-logger/v2"
)

type User struct {
	Id       uint64
	Username string
	Name     *string
	SurName  *string
	Age      int
	Sex      int
	City     *string
	Friend   bool
}

func (u *User) String() string {
	return string(u.Marshal())
}

func (u *User) Marshal() []byte {

	user, err := json.Marshal(*u)
	if err != nil {
		logger.Error(err)
		return nil
	}
	return user
}

const INSERT_INTO_USER_USERNAME_NAME_SURNAME_AGE_SEX_CITY = `
    INSERT INTO user
       (username, name, surname, age, sex, city)
      VALUES
       (?, ?, ?, ?, ?, ?)`

func (u *user) Create(user *User) (*User, error) {
	// Подготовить оператор для вставки данных
	stmtIns, err := u.dbRw.Prepare(INSERT_INTO_USER_USERNAME_NAME_SURNAME_AGE_SEX_CITY) // ? = заполнитель

	if err != nil {
		return user, err // правильная обработка ошибок вместо паники
	}
	defer func() { _ = stmtIns.Close() }() // Закрывается оператор, когда выйдете из функции

	res, err := stmtIns.Exec(user.Username, user.Name, user.SurName, user.Age, user.Sex, user.City)
	if err != nil {
		return user, err
	}
	id, err := (res.LastInsertId())
	if err != nil {
		return user, err
	}
	user.Id = uint64(id)

	return user, nil
}
