package domain

import (
	"encoding/json"
	"github.com/savsgio/go-logger/v2"
)

type Counter struct {
	Username string
	Total    uint64
	Unread   uint64
}

func (c *Counter) String() string {
	return string(c.Marshal())
}

func (c *Counter) Marshal() []byte {

	counter, err := json.Marshal(*c)
	if err != nil {
		logger.Error(err)
		return nil
	}
	return counter
}

func (c *counter) ReadByUserId(id uint64) (*Counter, error) {

	counter, err := c.readByUserId(id)
	if err != nil {
		return nil, err
	}
	return counter, nil
}

const SELECT_USERNAME_TOTAL_UNREAD_FROM_COUNTER_BY_ID = `
    SELECT c.username, c.total, c.unread
      FROM user u
      JOIN counter c ON u.username = c.username
     WHERE u.id = ?`

func (c *counter) readByUserId(id uint64) (*Counter, error) {

	stmtOut, err := c.dbRo.Prepare(SELECT_USERNAME_TOTAL_UNREAD_FROM_COUNTER_BY_ID)
	if err != nil {
		return nil, err // правильная обработка ошибок вместо паники
	}
	defer func() { _ = stmtOut.Close() }()

	var value Counter

	err = stmtOut.QueryRow(id).Scan(&value.Username, &value.Total, &value.Unread)
	if err != nil {
		return nil, err
	}
	return &value, nil
}

func (c *counter) ReadByUserName(username string) (*Counter, error) {
	return c.readUserByName(username)
}

const SELECT_USERNAME_TOTAL_UNREAD_FROM_COUNTER_BY_USERNAME = `
    SELECT c.username, c.total, c.unread
      FROM user u
      JOIN counter c ON u.username = c.username
     WHERE username = ?`

func (c *counter) readUserByName(username string) (*Counter, error) {

	stmtOut, err := c.dbRo.Prepare(SELECT_USERNAME_TOTAL_UNREAD_FROM_COUNTER_BY_USERNAME)
	if err != nil {
		return nil, err // правильная обработка ошибок вместо паники
	}
	defer func() { _ = stmtOut.Close() }()

	var value Counter

	err = stmtOut.QueryRow(username).Scan(&value.Username, &value.Total, &value.Unread)
	if err != nil {
		return nil, err
	}
	return &value, nil
}

const INSERT_INTO_COUNTER_USERNAME_TOTAL_UNREAD = `
    INSERT INTO counter
       (username, total, unread)
      VALUES
       (?, ?, ?)`

func (c *counter) Create(counter *Counter) (*Counter, error) {
	// Подготовить оператор для вставки данных
	stmtIns, err := c.dbRw.Prepare(INSERT_INTO_COUNTER_USERNAME_TOTAL_UNREAD) // ? = заполнитель

	if err != nil {
		return counter, err // правильная обработка ошибок вместо паники
	}
	defer func() { _ = stmtIns.Close() }() // Закрывается оператор, когда выйдете из функции

	_, err = stmtIns.Exec(counter.Username, counter.Total, counter.Unread)
	if err != nil {
		return counter, err
	}

	return counter, nil
}

const UPDATE_COUNTER_TOTAL_UNREAD_BY_USERNAME = `
    UPDATE counter
       SET total = ?, unread = ?
     WHERE username = ?`

func (c *counter) Update(counter *Counter) (*Counter, error) {
	// Подготовить оператор для вставки данных
	stmtIns, err := c.dbRw.Prepare(UPDATE_COUNTER_TOTAL_UNREAD_BY_USERNAME) // ? = заполнитель

	if err != nil {
		return counter, err // правильная обработка ошибок вместо паники
	}
	defer func() { _ = stmtIns.Close() }() // Закрывается оператор, когда выйдете из функции

	_, err = stmtIns.Exec(counter.Total, counter.Unread, counter.Username)
	if err != nil {
		return counter, err
	}

	return counter, nil
}
