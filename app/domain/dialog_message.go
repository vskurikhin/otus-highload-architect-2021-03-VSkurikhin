package domain

import (
	"encoding/json"
	"github.com/savsgio/go-logger/v2"
)

type DialogMessage struct {
	Id       uint64
	ShardId  uint8
	FromUser uint64
	ToUser   uint64
	Message  string
	ParentId *uint64
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

	switch shardId {
	case 1:
		return d.createOnShardId1(dm)
	default:
		return d.create(dm)
	}
}

const INSERT_INTO_DIALOG_MESSAGE = `
	INSERT INTO dialog_message
	    (shard_id, id, from_user, to_user, message, parent_id)
	     VALUES
	    (0, ?, ?, ?, ?, ?)`

func (d *dialogMessage) create(dm *DialogMessage) error {
	// Подготовить оператор для вставки данных
	stmtIns, err := d.dbRw.Prepare(INSERT_INTO_DIALOG_MESSAGE) // ? = заполнитель

	if err != nil {
		return err // правильная обработка ошибок вместо паники
	}
	defer func() { _ = stmtIns.Close() }() // Закрывается оператор, когда выйдете из функции
	_, err = stmtIns.Exec(dm.Id, dm.FromUser, dm.ToUser, dm.Message, dm.ParentId)

	if err != nil {
		return err
	}

	return nil
}

const INSERT_INTO_DIALOG_MESSAGE_SHARD_ID_1 = "INSERT INTO dialog_message" +
	" (shard_id, id, from_user, to_user, message, parent_id)" +
	" VALUES (1, ?, ?, ?, ?, ?)"

func (d *dialogMessage) createOnShardId1(dm *DialogMessage) error {
	// Подготовить оператор для вставки данных
	stmtIns, err := d.dbRw.Prepare(INSERT_INTO_DIALOG_MESSAGE_SHARD_ID_1) // ? = заполнитель

	if err != nil {
		return err // правильная обработка ошибок вместо паники
	}
	defer func() { _ = stmtIns.Close() }() // Закрывается оператор, когда выйдете из функции
	_, err = stmtIns.Exec(dm.Id, dm.FromUser, dm.ToUser, dm.Message, dm.ParentId)

	if err != nil {
		return err
	}

	return nil
}

const UPDATE_DIALOG_MESSAGE_MESSAGE = `
	UPDATE dialog_message
	   SET message = ?
	 WHERE shard_id = ? AND id = ?`

func (d *dialogMessage) update(dm *DialogMessage) error {
	// Подготовить оператор для вставки данных
	stmtIns, err := d.dbRw.Prepare(UPDATE_DIALOG_MESSAGE_MESSAGE) // ? = заполнитель

	if err != nil {
		return err // правильная обработка ошибок вместо паники
	}
	defer func() { _ = stmtIns.Close() }() // Закрывается оператор, когда выйдете из функции

	_, err = stmtIns.Exec(dm.Message, dm.ShardId, dm.Id)

	if err != nil {
		return err
	}

	return nil
}
