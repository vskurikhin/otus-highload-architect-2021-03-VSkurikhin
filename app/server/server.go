package server

import (
	"database/sql"
	"fmt"
	sa "github.com/savsgio/atreugo/v11"
	"github.com/savsgio/go-logger/v2"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app/config"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app/domain"
	"log"
	"os"
)

// Server определяет параметры для запуска HTTP-сервера.
type Server struct {
	DAO *domain.DAO
	// JWT    *security.JWT
	Server *sa.Atreugo
}

// New инициализирует сервер для ответа на сетевые запросы HTTP.
func New(cfg *config.Config) *Server {

	c := sa.Config{
		Addr:             cfg.Server.Host + `:` + cfg.Server.Port,
		Compress:         true,
		Name:             "httpd",
		GracefulShutdown: true,
	}
	dbRo := openDBRo(cfg)
	dbRw := openDBRw(cfg)
	go gracefulClose(dbRo, dbRw)
	versionDB(dbRw)

	return &Server{
		DAO: domain.New(dbRo, dbRw),
		// JWT:    security.New(cfg),
		Server: sa.New(c),
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
