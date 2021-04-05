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
	"hl.svn.su/highload-architect/app/config"
	"hl.svn.su/highload-architect/app/server"
)

func main() {
	// Загрузка конфигурации
	var envfile string
	flag.StringVar(&envfile, "env-file", "../.env", "Read in a file of environment variables")
	flag.Parse()

	godotenv.Load(envfile)
	config, err := config.Environ()
	if err != nil {
		logger := logrus.WithError(err)
		logger.Fatalln("main: invalid configuration")
	}
	// Инициализация логгирования
	initLogging(config)

	// Если логгирование на уровне трассировки включено, вывести
	// параметры конфигурации.
	if logrus.IsLevelEnabled(logrus.TraceLevel) {
		fmt.Println(config.String())
	}
}

// вспомогательная функция настраивает ведение логгирования.
func initLogging(c config.Config) {
	if c.Logging.Debug {
		logrus.SetLevel(logrus.DebugLevel)
	}
	if c.Logging.Trace {
		logrus.SetLevel(logrus.TraceLevel)
	}
	if c.Logging.Text {
		logrus.SetFormatter(&logrus.TextFormatter{
			ForceColors:   c.Logging.Color,
			DisableColors: !c.Logging.Color,
		})
	} else {
		logrus.SetFormatter(&logrus.JSONFormatter{
			PrettyPrint: c.Logging.Pretty,
		})
	}
}

// application является основной main структурой для сервера.
type application struct {
	Server *server.Server
}

// newApplication создает новую структуру приложения application.
func newApplication(
	Server *server.Server,
) application {
	return application{
		Server: Server,
	}
}
