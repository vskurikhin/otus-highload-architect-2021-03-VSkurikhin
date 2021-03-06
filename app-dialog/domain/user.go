package domain

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/savsgio/go-logger/v2"
	"strconv"
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

func (u *user) ReadUserById(id uint64) (*User, error) {

	user, err := u.readUser(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *user) ReadUserBySid(sid string) (*User, error) {

	id, err := strconv.ParseUint(sid, 10, 64)
	if err != nil {
		return nil, err // правильная обработка ошибок вместо паники
	}
	return u.ReadUserById(id)
}

const SELECT_ID_USERNAME_NAME_SURNAME_AGE_SEX_CITY_FROM_USER = `
    SELECT id, username, name, surname, age, sex, city
      FROM user
     WHERE id = ?`

func (u *user) readUser(id uint64) (*User, error) {

	stmtOut, err := u.dbRo.Prepare(SELECT_ID_USERNAME_NAME_SURNAME_AGE_SEX_CITY_FROM_USER)
	if err != nil {
		return nil, err // правильная обработка ошибок вместо паники
	}
	defer func() { _ = stmtOut.Close() }()

	var user User

	err = stmtOut.QueryRow(id).
		Scan(&user.Id, &user.Username, &user.Name, &user.SurName, &user.Age, &user.Sex, &user.City)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *user) ReadUserByName(name string) (*User, error) {
	return u.readUserByName(name)
}

const SELECT_ID_USERNAME_NAME_SURNAME_AGE_SEX_CITY_FROM_USER_BY_NAME = `
    SELECT u.id, u.username, u.name, u.surname, u.age, u.sex, u.city
      FROM user u
      JOIN login l ON l.id = u.id
     WHERE l.username = ?`

func (u *user) readUserByName(name string) (*User, error) {

	stmtOut, err := u.dbRo.Prepare(SELECT_ID_USERNAME_NAME_SURNAME_AGE_SEX_CITY_FROM_USER_BY_NAME)
	if err != nil {
		return nil, err // правильная обработка ошибок вместо паники
	}
	defer func() { _ = stmtOut.Close() }()

	var user User

	err = stmtOut.QueryRow(name).
		Scan(&user.Id, &user.Username, &user.Name, &user.SurName, &user.Age, &user.Sex, &user.City)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

const INSERT_INTO_USER_USERNAME_NAME_SURNAME_AGE_SEX_CITY = `
    INSERT INTO user
       (id, username, name, surname, age, sex, city)
      VALUES
       (?, ?, ?, ?, ?, ?, ?)`

func (u *user) Create(user *User) (*User, error) {
	// Подготовить оператор для вставки данных
	stmtIns, err := u.dbRw.Prepare(INSERT_INTO_USER_USERNAME_NAME_SURNAME_AGE_SEX_CITY) // ? = заполнитель

	if err != nil {
		return user, err // правильная обработка ошибок вместо паники
	}
	defer func() { _ = stmtIns.Close() }() // Закрывается оператор, когда выйдете из функции

	res, err := stmtIns.Exec(user.Id, user.Username, user.Name, user.SurName, user.Age, user.Sex, user.City)
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

const SELECT_USER_JOIN_INTERESTS_WHERE = `
    SELECT id, username, name, surname, age, sex, city
      FROM user
     WHERE name LIKE '%s%%%%' AND surname LIKE '%s%%%%'`

func (u *user) SearchUserList(id uint64, name, surname string) ([]User, error) {

	query := fmt.Sprintf(SELECT_USER_JOIN_INTERESTS_WHERE, name, surname)
	if logger.DebugEnabled() {
		logger.Debugf("SearchUserList query: %s", query)
	}
	stmtOut, err := u.dbRo.Prepare(query)
	if err != nil {
		return nil, err // правильная обработка ошибок вместо паники
	}
	defer func() { _ = stmtOut.Close() }()

	rows, err := stmtOut.Query()

	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	var users []User
	for rows.Next() {

		var r User
		err = rows.Scan(&r.Id, &r.Username, &r.Name, &r.SurName, &r.Age, &r.Sex, &r.City)

		if err != nil {
			return nil, err
		}
		users = append(users, r)
	}
	return users, nil
}

const SELECT_USER_JOIN_INTERESTS_WHERE_NAME = `
    SELECT id, username, name, surname, age, sex, city 
      FROM user
     WHERE name LIKE '%s%%%%'`

const SELECT_USER_JOIN_INTERESTS_WHERE_SURNAME = `
    SELECT id, username, name, surname, age, sex, city 
      FROM user
     WHERE surname LIKE '%s%%%%'`

func (u *user) SearchByUserList(id uint64, field, value string) ([]User, error) {

	var err error
	var stmtOut *sql.Stmt
	switch field {
	case "name":
		query := fmt.Sprintf(SELECT_USER_JOIN_INTERESTS_WHERE_NAME, value)
		if logger.DebugEnabled() {
			logger.Debugf("SearchByUserList query: %s", query)
		}
		stmtOut, err = u.dbRo.Prepare(query)
		break
	case "surname":
		query := fmt.Sprintf(SELECT_USER_JOIN_INTERESTS_WHERE_SURNAME, value)
		if logger.DebugEnabled() {
			logger.Debugf("SearchByUserList query: %s", query)
		}
		stmtOut, err = u.dbRo.Prepare(query)
		break
	default:
		return nil, errors.New("Unknown field: " + field)
	}
	if err != nil {
		return nil, err // правильная обработка ошибок вместо паники
	}
	defer func() { _ = stmtOut.Close() }()

	rows, err := stmtOut.Query()

	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	var users []User
	for rows.Next() {

		var r User
		err = rows.Scan(&r.Id, &r.Username, &r.Name, &r.SurName, &r.Age, &r.Sex, &r.City)

		if err != nil {
			return nil, err
		}
		users = append(users, r)
	}
	return users, nil
}
