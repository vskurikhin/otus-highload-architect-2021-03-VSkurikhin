package server

import (
	"database/sql"
	"fmt"
	"github.com/atreugo/websocket"
	sa "github.com/savsgio/atreugo/v11"
	"github.com/savsgio/go-logger/v2"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app-main/config"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app-main/domain"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app-main/security"
	"log"
	"os"
	"time"
)

// Server определяет параметры для запуска HTTP-сервера.
type Server struct {
	DAO      *domain.DAO
	JWT      *security.JWT
	Server   *sa.Atreugo
	Services *config.Services
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
	dao := domain.New(dbRo, dbRw)
	_, err := dao.ShardingMap.ReadMap()
	if err != nil {
		panic(err)
	}

	return &Server{
		DAO:      dao,
		JWT:      security.New(cfg),
		Server:   sa.New(c),
		Services: &config.Services{Dialog: cfg.Services.Dialog},
	}
}

func (s *Server) UseBefore(fns sa.Middleware) *sa.Router {
	return s.Server.UseBefore(fns)
}

func (s *Server) StaticCustom() *sa.Path {

	return s.Server.StaticCustom("/", &sa.StaticFS{
		Root:               "web/public",
		GenerateIndexPages: true,
		AcceptByteRange:    true,
		PathRewrite: func(ctx *sa.RequestCtx) []byte {
			return ctx.Path()
		},
		PathNotFound: func(ctx *sa.RequestCtx) error {
			return ctx.TextResponse("File not found", 404)
		},
	})
}

// GET устанавливает обработчик для GET запросов
func (s *Server) GET(url string, viewFn sa.View) *sa.Path {
	return s.Server.GET(url, viewFn)
}

// POST устанавливает обработчик для POST запросов
func (s *Server) POST(url string, viewFn sa.View) *sa.Path {
	return s.Server.POST(url, viewFn)
}

// PUT устанавливает обработчик для PUT запросов
func (s *Server) PUT(url string, viewFn sa.View) *sa.Path {
	return s.Server.PUT(url, viewFn)
}

func (s *Server) DELETE(url string, viewFn sa.View) *sa.Path {
	return s.Server.DELETE(url, viewFn)
}

var upgrader = websocket.New(websocket.Config{
	AllowedOrigins: []string{"*"},
})

func (s *Server) WS(url string, viewFn websocket.View) *sa.Path {
	var WsNewsList = upgrader.Upgrade(viewFn)
	return s.Server.GET(url, WsNewsList)
}

// ListenAndServe запускает сервер для ответа на сетевые запросы HTTP.
func (s *Server) ListenAndServe() error {
	return s.Server.ListenAndServe()
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
	dsn := fmt.Sprintf(`%s:%s@tcp(%s:%d)/%s?timeout=5s`, dbCFG.Username, dbCFG.Password, dbCFG.HostRw, dbCFG.PortRw, dbCFG.DBName)
	fmt.Println(dsn)
	db, err := sql.Open("mysql", dsn)
	db.SetMaxOpenConns(37)
	db.SetConnMaxLifetime(time.Hour)

	if err != nil {
		panic(err.Error())
	}
	return db
}

func openDBRo(cfg *config.Config) *sql.DB {

	dbCFG := cfg.DataBase
	dsn := fmt.Sprintf(`%s:%s@tcp(%s:%d)/%s?timeout=1s`, dbCFG.Username, dbCFG.Password, dbCFG.HostRo, dbCFG.PortRo, dbCFG.DBName)
	fmt.Println(dsn)
	db, err := sql.Open("mysql", dsn)
	db.SetMaxOpenConns(66)
	db.SetConnMaxLifetime(time.Hour)

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
