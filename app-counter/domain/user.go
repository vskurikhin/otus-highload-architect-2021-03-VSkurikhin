package domain

import (
	"encoding/json"
	"github.com/savsgio/go-logger/v2"
)

type User struct {
	Id       uint64
	Username string
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

func UnmarshalUser(bytes []byte) (*User, error) {

	var value User
	err := json.Unmarshal(bytes, &value)

	if logger.DebugEnabled() {
		logger.Debugf("bytes: %v", value)
	}
	if err != nil {
		return nil, err
	}
	return &value, nil
}

func (u *user) ReadUserById(id uint64) (*User, error) {

	user, err := u.readUser(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

const SELECT_ID_USERNAME_FROM_USER = `
    SELECT id, username
      FROM user
     WHERE id = ?`

func (u *user) readUser(id uint64) (*User, error) {

	stmtOut, err := u.dbRo.Prepare(SELECT_ID_USERNAME_FROM_USER)
	if err != nil {
		return nil, err // правильная обработка ошибок вместо паники
	}
	defer func() { _ = stmtOut.Close() }()

	var user User

	err = stmtOut.QueryRow(id).
		Scan(&user.Id, &user.Username)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *user) ReadUserByName(name string) (*User, error) {
	return u.readUserByName(name)
}

const SELECT_ID_USERNAME_FROM_USER_BY_NAME = `
    SELECT id, username
      FROM user
     WHERE username = ?`

func (u *user) readUserByName(name string) (*User, error) {

	stmtOut, err := u.dbRo.Prepare(SELECT_ID_USERNAME_FROM_USER_BY_NAME)
	if err != nil {
		return nil, err // правильная обработка ошибок вместо паники
	}
	defer func() { _ = stmtOut.Close() }()

	var user User

	err = stmtOut.QueryRow(name).
		Scan(&user.Id, &user.Username)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

const INSERT_INTO_USER_USERNAME_NAME_SURNAME_AGE_SEX_CITY = `
    INSERT INTO user
       (id, username)
      VALUES
       (?, ?)`

func (u *user) Create(user *User) (*User, error) {
	// Подготовить оператор для вставки данных
	stmtIns, err := u.dbRw.Prepare(INSERT_INTO_USER_USERNAME_NAME_SURNAME_AGE_SEX_CITY) // ? = заполнитель

	if err != nil {
		return user, err // правильная обработка ошибок вместо паники
	}
	defer func() { _ = stmtIns.Close() }() // Закрывается оператор, когда выйдете из функции

	res, err := stmtIns.Exec(user.Id, user.Username)
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
