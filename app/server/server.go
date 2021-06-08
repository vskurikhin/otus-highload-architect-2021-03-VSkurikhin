package server

import (
	"database/sql"
	"fmt"
	"github.com/atreugo/websocket"
	_ "github.com/go-sql-driver/mysql"
	sa "github.com/savsgio/atreugo/v11"
	"github.com/savsgio/go-logger/v2"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app/cache"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app/config"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app/domain"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app/pubsub"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app/security"
	"log"
	"os"
)

// Server определяет параметры для запуска HTTP-сервера.
type Server struct {
	Cache  *cache.Redis
	DAO    *domain.DAO
	JWT    *security.JWT
	PubSub *pubsub.Redis
	Server *sa.Atreugo
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
	redis := cache.CreateRedisCacheClient(cfg)

	return &Server{
		Cache:  redis,
		DAO:    domain.New(dbRo, dbRw),
		JWT:    security.New(cfg),
		PubSub: pubsub.NewRedis(redis.Cache),
		Server: sa.New(c),
	}
}

func (s *Server) UseBefore(fns sa.Middleware) *sa.Router {
	return s.Server.UseBefore(fns)
}

func (s *Server) StaticCustom() *sa.Path {

	pathRewriteCalled := false

	return s.Server.StaticCustom("/", &sa.StaticFS{
		Root:               "web/public",
		GenerateIndexPages: true,
		AcceptByteRange:    true,
		PathRewrite: func(ctx *sa.RequestCtx) []byte {
			pathRewriteCalled = true

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
