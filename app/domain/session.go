package domain

import (
	"database/sql"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/savsgio/go-logger/v2"
)

type Session struct {
	id        uuid.UUID
	SessionId uuid.UUID
}

func (s *Session) Id() uuid.UUID {
	return s.id
}

func (s *Session) SetId(id uuid.UUID) {
	s.id = id
}

func (s *Session) String() string {
	return string(s.Marshal())
}

func (s *Session) Marshal() []byte {

	login, err := json.Marshal(*s)
	if err != nil {
		logger.Error(err)
		return nil
	}
	return login
}

func (s *session) UpdateOrCreate(login *Login, sessionId uuid.UUID) error {

	stmtOut, err := s.dbRo.Prepare("SELECT id FROM `session` WHERE id = ?")

	if err != nil {
		return err // правильная обработка ошибок вместо паники
	}
	defer func() { _ = stmtOut.Close() }() // Закрывается оператор, когда выйдете из функции

	userId := new(uuid.UUID)
	loginId, err := login.Id().MarshalBinary()
	err = stmtOut.QueryRow(loginId).Scan(userId)

	if err == sql.ErrNoRows {
		return s.create(login, sessionId)
	}
	return s.update(*userId, sessionId)
}

func (s *session) create(login *Login, sessionId uuid.UUID) error {
	// Подготовить оператор для вставки данных
	stmtIns, err := s.dbRw.Prepare("INSERT INTO `session` (id, session_id) VALUES (?, ?)") // ? = заполнитель

	if err != nil {
		return err // правильная обработка ошибок вместо паники
	}
	defer func() { _ = stmtIns.Close() }() // Закрывается оператор, когда выйдете из функции

	id, err := login.Id().MarshalBinary()

	if err != nil {
		return err
	}
	sid, err := sessionId.MarshalBinary()

	if err != nil {
		return err
	}
	_, err = stmtIns.Exec(id, sid)

	if err != nil {
		return err
	}

	if logger.DebugEnabled() {
		logger.Debugf("session %s created", sessionId)
	}
	return nil
}

func (s *session) update(userId uuid.UUID, sessionId uuid.UUID) error {
	// Подготовить оператор для вставки данных
	stmtIns, err := s.dbRw.Prepare("UPDATE `session` SET session_id = ? WHERE id = ?") // ? = заполнитель

	if err != nil {
		return err // правильная обработка ошибок вместо паники
	}
	defer func() { _ = stmtIns.Close() }() // Закрывается оператор, когда выйдете из функции

	uid, err := userId.MarshalBinary()

	if err != nil {
		return err
	}
	sid, err := sessionId.MarshalBinary()

	if err != nil {
		return err
	}
	_, err = stmtIns.Exec(sid, uid)

	if err != nil {
		return err
	}

	if logger.DebugEnabled() {
		logger.Debugf("session %s updated", sessionId)
	}
	return nil
}

const SELECT_USER_ID_AND_SESSION_ID_BY_USERNAME = `
	SELECT s.id, session_id
	  FROM ` + "`session`" + ` s
	  JOIN login l ON s.id = l.id
	 WHERE username = ?`

func (s *session) ReadByUsername(username string) (*Session, error) {

	stmtOut, err := s.dbRo.Prepare(SELECT_USER_ID_AND_SESSION_ID_BY_USERNAME)

	if err != nil {
		return nil, err // правильная обработка ошибок вместо паники
	}
	defer func() { _ = stmtOut.Close() }() // Закрывается оператор, когда выйдете из функции

	var session Session
	err = stmtOut.QueryRow(username).
		Scan(&session.id, &session.SessionId)

	if err != nil {
		return nil, err
	}

	return &session, nil
}

const SELECT_USERNAME_BY_SESSION_ID = `
	SELECT l.username
	  FROM ` + "`session`" + ` s
	  JOIN login l ON s.id = l.id
	 WHERE session_id = ?`

func (s *session) UsernameBySessionId(sessionId uuid.UUID) (*string, error) {

	stmtOut, err := s.dbRo.Prepare(SELECT_USERNAME_BY_SESSION_ID)

	if err != nil {
		return nil, err // правильная обработка ошибок вместо паники
	}
	defer func() { _ = stmtOut.Close() }() // Закрывается оператор, когда выйдете из функции

	var username string
	id, err := sessionId.MarshalBinary()
	if err != nil {
		return nil, err // правильная обработка ошибок вместо паники
	}
	err = stmtOut.QueryRow(id).Scan(&username)

	if err != nil {
		return nil, err
	}

	return &username, nil
}

const SELECT_USER_ID_AND_USERNAME_BY_SESSION_ID = `
	SELECT l.id, l.username
	  FROM ` + "`session`" + ` s
	  JOIN login l ON s.id = l.id
	 WHERE session_id = ?`

func (s *session) ProfileBySessionId(sessionId uuid.UUID) (*Profile, error) {

	stmtOut, err := s.dbRo.Prepare(SELECT_USER_ID_AND_USERNAME_BY_SESSION_ID)

	if err != nil {
		return nil, err // правильная обработка ошибок вместо паники
	}
	defer func() { _ = stmtOut.Close() }() // Закрывается оператор, когда выйдете из функции

	var loginId uuid.UUID
	var username string
	id, err := sessionId.MarshalBinary()
	if err != nil {
		return nil, err // правильная обработка ошибок вместо паники
	}
	err = stmtOut.QueryRow(id).Scan(&loginId, &username)

	if err != nil {
		return nil, err
	}

	profile := Profile{Id: loginId, Username: username}

	return &profile, nil
}
