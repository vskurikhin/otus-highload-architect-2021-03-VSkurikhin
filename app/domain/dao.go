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
		db *sql.DB
	}
	login struct {
		db *sql.DB
	}
	session struct {
		db *sql.DB
	}
	user struct {
		db *sql.DB
	}
	userHasFriends struct {
		db *sql.DB
	}
	userHasInterests struct {
		db *sql.DB
	}
)

func New(db *sql.DB) *DAO {
	return &DAO{
		Interest:         &interest{db: db},
		Login:            &login{db: db},
		Session:          &session{db: db},
		User:             &user{db: db},
		UserHasFriends:   &userHasFriends{db: db},
		UserHasInterests: &userHasInterests{db: db},
	}
}
