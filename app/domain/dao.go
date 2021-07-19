package domain

import "database/sql"

type (
	DAO struct {
		Login   *login
		Session *session
		User    *user
	}
	login struct {
		dbRo *sql.DB
		dbRw *sql.DB
	}
	session struct {
		dbRo *sql.DB
		dbRw *sql.DB
	}
	user struct {
		dbRo *sql.DB
		dbRw *sql.DB
	}
)

func New(dbRo *sql.DB, dbRw *sql.DB) *DAO {
	return &DAO{
		Login:   &login{dbRo: dbRo, dbRw: dbRw},
		Session: &session{dbRo: dbRo, dbRw: dbRw},
		User:    &user{dbRo: dbRo, dbRw: dbRw},
	}
}
