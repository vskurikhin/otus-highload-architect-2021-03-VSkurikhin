package domain

import (
	"database/sql"
	"fmt"
	"github.com/savsgio/go-logger/v2"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app-counter/config"
	"log"
	"os"
)

type (
	DAO struct {
		Counter *counter
		User    *user
	}
	counter struct {
		dbRo *sql.DB
		dbRw *sql.DB
	}
	user struct {
		dbRo *sql.DB
		dbRw *sql.DB
	}
)

func New(cfg *config.Config) *DAO {

	dbRo := openDBRo(cfg)
	dbRw := openDBRw(cfg)
	go gracefulClose(dbRo, dbRw)
	versionDB(dbRw)

	return &DAO{
		Counter: &counter{dbRo: dbRo, dbRw: dbRw},
		User:    &user{dbRo: dbRo, dbRw: dbRw},
	}
}

func gracefulClose(dbRo *sql.DB, dbRw *sql.DB) {
	// Настраиваем канал для отправки сигнальных уведомлений.
	// Нужно использовать буферизованный канал или есть риск пропустить сигнал
	// если не готовы принять сигнал при отправке.
	c := make(chan os.Signal, 1)

	// Блокировать до получения сигнала.
	s := <-c
	fmt.Println("Got signal:", s)
	err := dbRw.Close()
	if err != nil {
		log.Println(err)
	}
	err = dbRo.Close()
	if err != nil {
		log.Println(err)
	}
}

func openDBRw(cfg *config.Config) *sql.DB {

	dbCFG := cfg.DataBase
	dsn := fmt.Sprintf(`%s:%s@tcp(%s:%d)/%s`, dbCFG.Username, dbCFG.Password, dbCFG.HostRw, dbCFG.PortRw, dbCFG.DBName)
	fmt.Println(dsn)
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		panic(err.Error())
	}
	return db
}

func openDBRo(cfg *config.Config) *sql.DB {

	dbCFG := cfg.DataBase
	dsn := fmt.Sprintf(`%s:%s@tcp(%s:%d)/%s`, dbCFG.Username, dbCFG.Password, dbCFG.HostRo, dbCFG.PortRo, dbCFG.DBName)
	fmt.Println(dsn)
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		panic(err.Error())
	}
	return db
}

func versionDB(db *sql.DB) {

	if logger.DebugEnabled() {
		var version string
		err := db.QueryRow("SELECT VERSION()").Scan(&version)

		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		fmt.Println(version)
	}
}
