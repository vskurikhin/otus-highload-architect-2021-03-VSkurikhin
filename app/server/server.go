package server

import (
	"github.com/savsgio/atreugo/v11"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app/config"
)

// Server определяет параметры для запуска HTTP-сервера.
type Server struct {
	Server *atreugo.Atreugo
}

// New инициализирует сервер для ответа на сетевые запросы HTTP.
func New(cfg *config.Config) *Server {
	c := atreugo.Config{
		Addr: cfg.Server.Host,
	}

	return &Server{Server: atreugo.New(c)}
}

func (s *Server) UseBefore(fns atreugo.Middleware) *atreugo.Router {
	return s.Server.UseBefore(fns)
}

// GET устанавливает обработчик для GET запросов
func (s *Server) GET(url string, viewFn atreugo.View) *atreugo.Path {
	return s.Server.GET(url, viewFn)
}

// POST устанавливает обработчик для POST запросов
func (s *Server) POST(url string, viewFn atreugo.View) *atreugo.Path {
	return s.Server.POST(url, viewFn)
}

// PUT устанавливает обработчик для PUT запросов
func (s *Server) PUT(url string, viewFn atreugo.View) *atreugo.Path {
	return s.Server.PUT(url, viewFn)
}

func (s *Server) DELETE(url string, viewFn atreugo.View) *atreugo.Path {
	return s.Server.DELETE(url, viewFn)
}

// ListenAndServe запускает сервер для ответа на сетевые запросы HTTP.
func (s *Server) ListenAndServe() error {
	return s.Server.ListenAndServe()
}
