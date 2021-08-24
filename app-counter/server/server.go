package server

import (
	"github.com/atreugo/websocket"
	sa "github.com/savsgio/atreugo/v11"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app-counter/config"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app-counter/domain"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app-counter/security"
)

// Server определяет параметры для запуска HTTP-сервера.
type Server struct {
	DAO    *domain.DAO
	JWT    *security.JWT
	Server *sa.Atreugo
}

// New инициализирует сервер для ответа на сетевые запросы HTTP.
func New(cfg *config.Config, dao *domain.DAO) *Server {

	c := sa.Config{
		Addr:             cfg.Server.Host + `:` + cfg.Server.Port,
		Compress:         true,
		Name:             "httpd",
		GracefulShutdown: true,
	}
	return &Server{
		DAO:    dao,
		JWT:    security.New(cfg),
		Server: sa.New(c),
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
