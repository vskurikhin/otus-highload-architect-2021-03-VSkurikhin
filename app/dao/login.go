package dao

import (
	"github.com/google/uuid"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app/model"
)

func (l *login) Create(login *model.Login) error {
	// Подготовить оператор для вставки данных
	stmtIns, err := l.db.Prepare("INSERT INTO login (id, username, password) VALUES (?, ?, ?)") // ? = заполнитель
	if err != nil {
		return err // правильная обработка ошибок вместо паники
	}
	defer func() { _ = stmtIns.Close() }() // Закрывается оператор, когда выйдете из функции

	id, err := login.Id().MarshalBinary()
	_, err = stmtIns.Exec(id, login.Username, login.Password())
	if err != nil {
		return err
	}
	return nil
}

func (l *login) Read(username string) (*model.Login, error) {

	stmtOut, err := l.db.Prepare(`
 		SELECT id, username, password FROM login WHERE username = ?
 	`)
	if err != nil {
		return nil, err // правильная обработка ошибок вместо паники
	}
	defer func() { _ = stmtOut.Close() }() // Закрывается оператор, когда выйдете из функции

	var id uuid.UUID
	var login model.Login
	var password string
	err = stmtOut.QueryRow(username).
		Scan(&id, &login.Username, &password)
	if err != nil {
		return nil, err
	}
	login.SetId(id)
	login.SetPassword(password)

	return &login, nil
}
