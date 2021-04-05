// Copyright 2021 Victor N. Skurikhin
//
// Licensed under the Unlicense;
// For more information, please refer to <https://unlicense.org>

package main

import (
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
	"hl.svn.su/highload-architect/app/config"
	"hl.svn.su/highload-architect/cmd"
)

func main() {
	// Загрузка конфигурации
	var envFile string
	flag.StringVar(&envFile, "env-file", "../.env", "Read in a file of environment variables")
	flag.Parse()

	err := godotenv.Load(envFile)
	if err != nil {
		logger := logrus.WithError(err)
		logger.Fatalln("main: can't load configuration")
	}
	environ, err := config.Environ()
	if err != nil {
		logger := logrus.WithError(err)
		logger.Fatalln("main: invalid configuration")
	}
	// Инициализация логгирования
	cmd.InitLogging(environ)

	// Если логгирование на уровне трассировки включено, вывести
	// параметры конфигурации.
	if logrus.IsLevelEnabled(logrus.TraceLevel) {
		fmt.Println(environ.String())
	}

	// Инициализация приложения
	app, err := cmd.InitializeApplication(environ)
	if err != nil {
		logger := logrus.WithError(err)
		logger.Fatalln("main: cannot initialize server")
	}

	// Запуск сервера
	g := errgroup.Group{}
	g.Go(func() error {
		logrus.WithFields(
			logrus.Fields{
				"Host": environ.Server.Host,
			},
		).Info("starting the http server")
		return app.Server.ListenAndServe()
	})

	// Ожидение the gorouitine
	if err := g.Wait(); err != nil {
		logrus.WithError(err).Fatalln("program terminated")
	}
}
