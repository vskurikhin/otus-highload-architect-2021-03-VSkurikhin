package main

import (
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/savsgio/go-logger/v2"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app-main/config"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app-main/server"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app-main/server/handlers"
)

func main() {

	// Загрузка конфигурации
	var envFile string
	flag.StringVar(&envFile, "env-file", ".env", "Read in a file of environment variables")
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
		logger.SetLevel(logger.DEBUG)
	}

	// Создать инстанс сервера
	s := server.New(environ)

	// Зарегистрируйте для аутентификации перед обработкой запросов.
	s.UseBefore(s.JWT.AuthCheckToken)

	// Обработка статичных фалов из каталога web/public.
	s.StaticCustom()

	// Обработчики запросов.
	h := handlers.Handlers{Server: s}

	// Зарегистрировать индексный маршрут.
	s.GET("/", h.Root)

	// Зарегистрировать login маршрут.
	s.POST("/login", h.Login)

	// Зарегистрировать маршрут для профиля пользователя.
	s.GET("/profile", h.Profile)

	// Зарегистрировать маршрут для списка пользователей.
	s.GET("/users/all", h.UserList)

	// Зарегистрировать маршруты для поиска пользователей.
	s.GET("/users/search/{name}/{surname}", h.UserSearch)
	s.GET("/users/search-by/{field}/{value}", h.SearchBy)

	// Зарегистрировать маршрут для списка пользователей.
	s.GET("/user/{id}", h.User)

	// Зарегистрировать маршрут для Sign-in пользователя.
	s.POST("/signin", h.UserSignIn)

	s.GET("/messages", h.GetMessages)
	s.POST("/message", h.PostMessage)

	// Run
	if err := s.ListenAndServe(); err != nil {
		panic(err)
	}
}
