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
	    (shard_id, id, from_user, to_user, message)
	     VALUES
	    (0, ?, ?, ?, ?)`

func (d *dialogMessage) create(dm *DialogMessage) error {
	// Подготовить оператор для вставки данных
	stmtIns, err := d.dbRw.Prepare(INSERT_INTO_DIALOG_MESSAGE) // ? = заполнитель

	if err != nil {
		return err // правильная обработка ошибок вместо паники
	}
	defer func() { _ = stmtIns.Close() }() // Закрывается оператор, когда выйдете из функции
	_, err = stmtIns.Exec(dm.Id, dm.FromUser, dm.ToUser, dm.Message)

	if err != nil {
		return err
	}

	return nil
}

const INSERT_INTO_DIALOG_MESSAGE_SHARD_ID_1 = "INSERT INTO dialog_message" +
	" (shard_id, id, from_user, to_user, message)" +
	" VALUES (1, ?, ?, ?, ?)"

func (d *dialogMessage) createOnShardId1(dm *DialogMessage) error {
	// Подготовить оператор для вставки данных
	stmtIns, err := d.dbRw.Prepare(INSERT_INTO_DIALOG_MESSAGE_SHARD_ID_1) // ? = заполнитель

	if err != nil {
		return err // правильная обработка ошибок вместо паники
	}
	defer func() { _ = stmtIns.Close() }() // Закрывается оператор, когда выйдете из функции
	_, err = stmtIns.Exec(dm.Id, dm.FromUser, dm.ToUser, dm.Message)

	if err != nil {
		return err
	}

	return nil
}

const SELECT_DIALOG_MESSAGE_BY_ID_SHARD_ID = `
        SELECT from_user, to_user, message
          FROM dialog_message
         WHERE shard_id = ? AND id = ?`

func (s *session) getByIdShardId(id uint64, shardId uint8) (*DialogMessage, error) {

	stmtOut, err := s.dbRo.Prepare(SELECT_DIALOG_MESSAGE_BY_ID_SHARD_ID)

	if err != nil {
		return nil, err // правильная обработка ошибок вместо паники
	}
	defer func() { _ = stmtOut.Close() }() // Закрывается оператор, когда выйдете из функции

	dm := DialogMessage{Id: id, ShardId: shardId}
	err = stmtOut.QueryRow(shardId, id).Scan(&dm.FromUser, &dm.ToUser, &dm.Message)

	if err != nil {
		return nil, err
	}

	return &dm, nil
}

const SELECT_ALL_DIALOG_MESSAGE_BY_SHARD_ID = `
        SELECT id, from_user, to_user, message
          FROM dialog_message
         WHERE shard_id = ?`

func (s *session) getAll(shardId uint8) (*DialogMessage, error) {

	stmtOut, err := s.dbRo.Prepare(SELECT_ALL_DIALOG_MESSAGE_BY_SHARD_ID)

	if err != nil {
		return nil, err // правильная обработка ошибок вместо паники
	}
	defer func() { _ = stmtOut.Close() }() // Закрывается оператор, когда выйдете из функции

	dm := DialogMessage{ShardId: shardId}
	err = stmtOut.QueryRow(shardId).Scan(&dm.Id, &dm.FromUser, &dm.ToUser, &dm.Message)

	if err != nil {
		return nil, err
	}

	return &dm, nil
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
