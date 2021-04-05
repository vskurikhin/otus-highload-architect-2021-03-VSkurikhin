package server

import (
	"github.com/fate-lovely/phi"
	"github.com/valyala/fasthttp"
	"golang.org/x/sync/errgroup"
)

// Server определяет параметры для запуска HTTP-сервера.
type Server struct {
	Host    string
	Handler phi.Handler
}

// ListenAndServe инициализирует сервер для ответа на сетевые запросы HTTP.
func (s Server) ListenAndServe() error {
	var g errgroup.Group
	g.Go(func() error {
		return fasthttp.ListenAndServe(s.Host, s.Handler.ServeFastHTTP)
	})
	return g.Wait()
}
