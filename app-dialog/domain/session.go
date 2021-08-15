package domain

import (
	"database/sql"
	"encoding/json"
	"github.com/savsgio/go-logger/v2"
)

type Session struct {
	Id        uint64
	SessionId uint64
}

func (s *Session) String() string {
	return string(s.Marshal())
}

func (s *Session) Marshal() []byte {

	session, err := json.Marshal(*s)
	if err != nil {
		logger.Error(err)
		return nil
	}
	return session
}

const SELECT_ID_FROM_SESSION = `
	SELECT id FROM ` + "`session`" + ` WHERE id = ?`

func (s *session) UpdateOrCreate(login *Login, sessionId uint64) error {

	stmtOut, err := s.dbRo.Prepare(SELECT_ID_FROM_SESSION)

	if err != nil {
		return err // правильная обработка ошибок вместо паники
	}
	defer func() { _ = stmtOut.Close() }() // Закрывается оператор, когда выйдете из функции

	var userId uint64
	err = stmtOut.QueryRow(login.Id).Scan(&userId)

	if err == sql.ErrNoRows {
		return s.create(login, sessionId)
	}
	return s.update(userId, sessionId)
}

const INSERT_INTO_SESSION_ID_SESSION_ID = `
	INSERT INTO ` + "`session`" + ` (id, session_id) VALUES (?, ?)`

func (s *session) create(login *Login, sessionId uint64) error {
	// Подготовить оператор для вставки данных
	stmtIns, err := s.dbRw.Prepare(INSERT_INTO_SESSION_ID_SESSION_ID) // ? = заполнитель

	if err != nil {
		return err // правильная обработка ошибок вместо паники
	}
	defer func() { _ = stmtIns.Close() }() // Закрывается оператор, когда выйдете из функции
	_, err = stmtIns.Exec(login.Id, sessionId)

	if err != nil {
		return err
	}

	if logger.DebugEnabled() {
		logger.Debugf("session %s created", sessionId)
	}
	return nil
}

const UPDATE_SESSION_SET_SESSION_ID = `
	UPDATE ` + "`session`" + ` SET session_id = ? WHERE id = ?`

func (s *session) update(userId uint64, sessionId uint64) error {
	// Подготовить оператор для вставки данных
	stmtIns, err := s.dbRw.Prepare(UPDATE_SESSION_SET_SESSION_ID) // ? = заполнитель

	if err != nil {
		return err // правильная обработка ошибок вместо паники
	}
	defer func() { _ = stmtIns.Close() }() // Закрывается оператор, когда выйдете из функции

	_, err = stmtIns.Exec(sessionId, userId)

	if err != nil {
		return err
	}

	if logger.DebugEnabled() {
		logger.Debugf("session %s updated", sessionId)
	}
	return nil
}

const SELECT_USER_ID_AND_USERNAME_BY_SESSION_ID = `
        SELECT l.id, l.username
          FROM ` + "`session`" + ` s
          JOIN login l ON s.id = l.id
         WHERE session_id = ?`

func (s *session) ProfileBySessionId(sessionId uint64) (*Profile, error) {

	stmtOut, err := s.dbRo.Prepare(SELECT_USER_ID_AND_USERNAME_BY_SESSION_ID)

	if err != nil {
		return nil, err // правильная обработка ошибок вместо паники
	}
	defer func() { _ = stmtOut.Close() }() // Закрывается оператор, когда выйдете из функции

	var loginId uint64
	var username string
	err = stmtOut.QueryRow(sessionId).Scan(&loginId, &username)

	if err != nil {
		return nil, err
	}

	profile := Profile{Id: loginId, Username: username}

	return &profile, nil
}
