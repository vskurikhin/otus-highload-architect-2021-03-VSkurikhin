package domain

import "database/sql"

type (
	DAO struct {
		DialogMessage *dialogMessage
		Login         *login
		Session       *session
		ShardingMap   *shardingMap
		User          *user
	}
	dialogMessage struct {
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
	shardingMap struct {
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
		DialogMessage: &dialogMessage{dbRo: dbRo, dbRw: dbRw},
		Login:         &login{dbRo: dbRo, dbRw: dbRw},
		Session:       &session{dbRo: dbRo, dbRw: dbRw},
		ShardingMap:   &shardingMap{dbRo: dbRo, dbRw: dbRw},
		User:          &user{dbRo: dbRo, dbRw: dbRw},
	}
}
