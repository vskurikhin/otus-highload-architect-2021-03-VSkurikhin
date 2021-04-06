package main

import (
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/savsgio/go-logger/v2"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app/config"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app/server"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app/server/handlers"
)

func init() { //nolint:gochecknoinits
	logger.SetLevel(logger.DEBUG)
}

func main() {

	// Загрузка конфигурации
	var envFile string
	flag.StringVar(&envFile, "env-file", "../.env", "Read in a file of environment variables")
	flag.Parse()
	err := godotenv.Load(envFile)

	if err != nil {
		logger.Debug("main: can't load configuration")
	}
	environ, err := config.Environ()

	// Если логгирование на уровне трассировки включено, вывести
	// параметры конфигурации.
	if environ != nil && environ.Logging.Debug {
		fmt.Println(environ.String())
	}

	// Создать инстанс сервера
	s := server.New(environ)

	// Зарегистрируйте для аутентификации перед обработкой запросов.
	s.UseBefore(s.JWT.AuthCheckToken)

	// Обработчики запросов.
	h := handlers.Handlers{Server: s}

	// Зарегистрировать индексный маршрут
	s.GET("/", h.Root)

	// Зарегистрировать login маршрут
	s.POST("/login", h.Login)

	// Run
	if err := s.ListenAndServe(); err != nil {
		panic(err)
	}
}
