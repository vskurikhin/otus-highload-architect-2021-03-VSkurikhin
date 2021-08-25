package domain

import (
	"encoding/json"
	"fmt"
	"github.com/savsgio/go-logger/v2"
)

type DialogMessage struct {
	Id       uint64
	ShardId  uint8
	FromUser uint64
	ToUser   uint64
	Message  string
	ParentId *uint64
	HashId   uint8
}

func (d *DialogMessage) String() string {
	return string(d.Marshal())
}

func (d *DialogMessage) Marshal() []byte {

	session, err := json.Marshal(*d)
	if err != nil {
		logger.Error(err)
		return nil
	}
	return session
}

func (d *dialogMessage) Create(shardId uint8, dm *DialogMessage) error {
	return d.create(dm, shardId)
}

const INSERT_INTO_DIALOG_MESSAGE = "INSERT INTO dialog_message " +
	" (shard_id, id, from_user, to_user, message, parent_id, hash_id) " +
	"  VALUES " +
	" (%d, ?, ?, ?, ?, ?, ?)"

func (d *dialogMessage) create(dm *DialogMessage, shardId uint8) error {
	// Подготовить оператор для вставки данных
	sql := fmt.Sprintf(INSERT_INTO_DIALOG_MESSAGE, shardId)
	logger.Debugf("create sql: %s", sql)
	stmtIns, err := d.dbRw.Prepare(sql) // ? = заполнитель

	if err != nil {
		return err // правильная обработка ошибок вместо паники
	}
	defer func() { _ = stmtIns.Close() }() // Закрывается оператор, когда выйдете из функции
	_, err = stmtIns.Exec(dm.Id, dm.FromUser, dm.ToUser, dm.Message, dm.ParentId, dm.HashId)

	if err != nil {
		return err
	}

	return nil
}

func (d *dialogMessage) UpdateReadMessage(id uint64, shardId uint8) error {
	return d.updateReadMessage(id, shardId)
}

const UPDATE_DIALOG_MESSAGE_SET_ALREADY_READ_TRUE = "UPDATE dialog_message " +
	"   SET already_read = true " +
	" WHERE shard_id = %d AND id = ?"

func (d *dialogMessage) updateReadMessage(id uint64, shardId uint8) error {
	// Подготовить оператор для вставки данных
	sql := fmt.Sprintf(UPDATE_DIALOG_MESSAGE_SET_ALREADY_READ_TRUE, shardId)
	logger.Debugf("create sql: %s", sql)
	stmtIns, err := d.dbRw.Prepare(sql) // ? = заполнитель

	if err != nil {
		return err // правильная обработка ошибок вместо паники
	}
	defer func() { _ = stmtIns.Close() }() // Закрывается оператор, когда выйдете из функции
	_, err = stmtIns.Exec(id)

	if err != nil {
		return err
	}

	return nil
}
