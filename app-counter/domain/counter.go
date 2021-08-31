package domain

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
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
    SELECT username, total, unread
      FROM counter
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

const SELECT_USERNAME_TOTAL_UNREAD_FROM_COUNTER_BY_USERNAME_1 = `
    SELECT (total > -1)
      FROM counter
     WHERE username = ?`

const INSERT_INTO_COUNTER_USERNAME_TOTAL_UNREAD = `
    INSERT INTO counter
       (username, total, unread)
      VALUES
       (?, 0, 0)`

const UPDATE_COUNTER_TOTAL_UNREAD_BY_USERNAME = `
    UPDATE counter
       SET total = total + 1,
           unread = unread + 1
     WHERE username = ?`

// Create a helper function for preparing failure results.
func fail(err error, id int) error {
	return fmt.Errorf("Upsert: (%d, %v)", id, err)
}

func (c *counter) Upsert(username string) error {

	ctx := context.Background()

	// Get a Tx for making transaction requests.
	tx, err := c.dbRw.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelRepeatableRead})
	if err != nil {
		return fail(err, 1)
	}
	// Defer a rollback in case anything fails.
	defer func() { _ = tx.Rollback() }()

	// Confirm that album inventory is enough for the order.
	var enough bool
	if err = tx.QueryRowContext(ctx, SELECT_USERNAME_TOTAL_UNREAD_FROM_COUNTER_BY_USERNAME_1, username).
		Scan(&enough); err != nil {
		if err == sql.ErrNoRows {
			// Insert the album inventory to remove the quantity in the order.
			_, err = tx.ExecContext(ctx, INSERT_INTO_COUNTER_USERNAME_TOTAL_UNREAD, username)
			if err != nil {
				return fail(err, 2)
			}
		}
	}

	// Update the album inventory to remove the quantity in the order.
	_, err = tx.ExecContext(ctx, UPDATE_COUNTER_TOTAL_UNREAD_BY_USERNAME, username)
	if err != nil {
		return fail(err, 4)
	}

	// Commit the transaction.
	if err = tx.Commit(); err != nil {
		return fail(err, 5)
	}
	return nil
}

const UPDATE_COUNTER_TOTAL_DOWN_COUNT_UNREAD_BY_USERNAME = `
    UPDATE counter
       SET unread = unread - 1 
     WHERE username = ?`

func (c *counter) Read(username string) error {

	ctx := context.Background()

	// Get a Tx for making transaction requests.
	tx, err := c.dbRw.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelRepeatableRead})
	if err != nil {
		return fail(err, 1)
	}
	// Defer a rollback in case anything fails.
	defer func() { _ = tx.Rollback() }()

	// Confirm that album inventory is enough for the order.
	var enough bool
	if err = tx.QueryRowContext(ctx, SELECT_USERNAME_TOTAL_UNREAD_FROM_COUNTER_BY_USERNAME_1, username).
		Scan(&enough); err != nil {
		if err == sql.ErrNoRows {
			return fail(err, 2)
		}
	}

	// Update the album inventory to remove the quantity in the order.
	_, err = tx.ExecContext(ctx, UPDATE_COUNTER_TOTAL_DOWN_COUNT_UNREAD_BY_USERNAME, username)
	if err != nil {
		return fail(err, 4)
	}

	// Commit the transaction.
	if err = tx.Commit(); err != nil {
		return fail(err, 5)
	}
	return nil
}
