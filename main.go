package main

import (
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/savsgio/go-logger/v2"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app/config"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app/server"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app/server/handlers"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app/test"
	"os"
	"strconv"
)

func httpd() {

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

	s.WS("/ws-newslist", h.WsNewsList)

	// Зарегистрировать login маршрут.
	s.POST("/login", h.Login)

	// Зарегистрировать маршрут для профиля пользователя.
	s.GET("/profile", h.Profile)

	// Зарегистрировать маршрут для списка новостей.
	s.GET("/news/range/{offset}/{limit}", h.NewsList)

	s.POST("/news/add", h.CreateNews)

	// Зарегистрировать маршрут для списка пользователей.
	s.GET("/users/all", h.UserList)

	// Зарегистрировать маршруты для поиска пользователей.
	s.GET("/users/search/{name}/{surname}", h.UserSearch)
	s.GET("/users/search-by/{field}/{value}", h.SearchBy)

	// Зарегистрировать маршрут для списка пользователей.
	s.GET("/user/{id}", h.User)

	// Зарегистрировать маршрут для создания пользователя.
	s.POST("/user", h.Create)

	// Зарегистрировать маршрут для добавления друга.
	s.POST("/friend", h.UserFriend)

	// Зарегистрировать маршрут для Sign-in пользователя.
	s.POST("/signin", h.UserSignIn)

	// Зарегистрировать маршрут для Sign-in пользователя.
	s.POST("/save", h.UserSave)

	// Run
	if err := s.ListenAndServe(); err != nil {
		panic(err)
	}
}

func main() {

	size := len(os.Args)
	if size == 2 {
		upperBound, err := strconv.ParseInt(os.Args[1], 10, 32)
		if err == nil {
			test.FakeIt(upperBound)
		}
	} else {
		httpd()
	}
}
