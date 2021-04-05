package cmd

import "hl.svn.su/highload-architect/app/server"

// Application является основной main структурой для сервера.
type Application struct {
	Server *server.Server
}

// newApplication создает новую структуру приложения Application.
func newApplication(Server *server.Server) Application {
	return Application{Server: Server}
}
