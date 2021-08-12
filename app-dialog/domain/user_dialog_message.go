package domain

import (
	"encoding/json"
	"fmt"
	"github.com/savsgio/go-logger/v2"
)

type UserDialogMessage struct {
	Id       uint64
	FromUser string
	ToUser   string
	Message  string
	ParentId *uint64
}

func (d *UserDialogMessage) String() string {
	return string(d.Marshal())
}

func (d *UserDialogMessage) Marshal() []byte {

	session, err := json.Marshal(*d)
	if err != nil {
		logger.Error(err)
		return nil
	}
	return session
}

const SELECT_ALL_DIALOG_MESSAGE_BY_SHARD_ID = `
        SELECT dm.id, fu.username AS from_user, tu.username AS to_user, dm.message, dm.parent_id
          FROM dialog_message dm
          JOIN ` + "`user`" + ` fu ON fu.id = dm.from_user
          JOIN ` + "`user`" + ` tu ON tu.id = dm.to_user
         WHERE dm.shard_id = %d AND dm.to_user = ?`

func (d *dialogMessage) GetAllByUserId(shardId uint8, userId uint64) ([]UserDialogMessage, error) {

	sql := fmt.Sprintf(SELECT_ALL_DIALOG_MESSAGE_BY_SHARD_ID, shardId)
	stmtOut, err := d.dbRo.Prepare(sql)

	if err != nil {
		return nil, err // правильная обработка ошибок вместо паники
	}
	defer func() { _ = stmtOut.Close() }() // Закрывается оператор, когда выйдете из функции

	rows, err := stmtOut.Query(userId) // .Scan(&dm.Name, &dm.FromUser, &dm.ToUser, &dm.Message, &dm.ParentId)
	if err != nil {
		return nil, err
	}
	var dialogMessages []UserDialogMessage
	for rows.Next() {

		var dm UserDialogMessage
		err = rows.Scan(&dm.Id, &dm.FromUser, &dm.ToUser, &dm.Message, &dm.ParentId)

		if err != nil {
			return nil, err
		}
		dialogMessages = append(dialogMessages, dm)
	}
	return dialogMessages, nil
}
