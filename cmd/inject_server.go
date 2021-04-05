package cmd

import (
	"github.com/fate-lovely/phi"
	"github.com/google/wire"
	"hl.svn.su/highload-architect/app/config"
	"hl.svn.su/highload-architect/app/handlers/root"
	"hl.svn.su/highload-architect/app/server"
)

// набор wire проводов для загрузки сервера.
var serverSet = wire.NewSet(
	provideRouter,
	provideServer,
)

// provideServer это функция провайдера Wire, которая возвращает
// http сервер, настроенный из среды окружения (environment).
func provideServer(handler phi.Handler, config config.Config) *server.Server {
	return &server.Server{
		Host:    config.Server.Host,
		Handler: handler,
	}
}

// provideRouter это функция провайдера Wire, которая возвращает маршрутизатор,
// обслуживающий предоставленные обработчики.
func provideRouter() phi.Handler {

	router := phi.NewRouter()
	router.Mount("/", root.Handler())

	return router
}
