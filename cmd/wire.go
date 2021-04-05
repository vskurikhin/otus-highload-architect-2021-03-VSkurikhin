//+build wireinject

package cmd

import (
	"github.com/google/wire"
	"hl.svn.su/highload-architect/app/config"
)

func InitializeApplication(config config.Config) (Application, error) {
	wire.Build(
		serverSet,
		newApplication,
	)
	return Application{}, nil
}
