// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package cmd

import (
	"hl.svn.su/highload-architect/app/config"
)

// Injectors from wire.go:

func InitializeApplication(config2 config.Config) (Application, error) {
	handler := provideRouter()
	server := provideServer(handler, config2)
	mainApplication := newApplication(server)
	return mainApplication, nil
}
