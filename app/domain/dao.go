package domain

import (
	"database/sql"
)

type (
	DAO struct {
		Interest         *interest
		Login            *login
		Session          *session
		User             *user
		UserHasFriends   *userHasFriends
		UserHasInterests *userHasInterests
	}
	interest struct {
		dbRo *sql.DB
		dbRw *sql.DB
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
	userHasFriends struct {
		dbRo *sql.DB
		dbRw *sql.DB
	}
	userHasInterests struct {
		dbRo *sql.DB
		dbRw *sql.DB
	}
)

func New(dbRo *sql.DB, dbRw *sql.DB) *DAO {
	return &DAO{
		Interest:         &interest{dbRo: dbRo, dbRw: dbRw},
		Login:            &login{dbRo: dbRo, dbRw: dbRw},
		Session:          &session{dbRo: dbRo, dbRw: dbRw},
		User:             &user{dbRo: dbRo, dbRw: dbRw},
		UserHasFriends:   &userHasFriends{dbRo: dbRo, dbRw: dbRw},
		UserHasInterests: &userHasInterests{dbRo: dbRo, dbRw: dbRw},
	}
}
