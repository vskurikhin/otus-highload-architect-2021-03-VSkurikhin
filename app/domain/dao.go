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
