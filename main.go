package main

import (
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/savsgio/go-logger/v2"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app/config"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app/server"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app/server/handlers"
	"log"
	"os"
	"strconv"
)

type TestRepl struct {
	Id   uuid.UUID
	Test string `fake:"{regex:[0-9a-zA-Z]{128}}"`
}

func (l *TestRepl) String() string {
	return string(l.Marshal())
}

func (l *TestRepl) Marshal() []byte {

	test, err := json.Marshal(l)
	if err != nil {
		return nil
	}
	return test
}

type dao struct {
	db *sql.DB
}

func openDB(cfg config.DataBase) *sql.DB {

	dsn := fmt.Sprintf(`%s:%s@tcp(%s:%d)/%s`, cfg.Username, cfg.Password, cfg.HostRw, cfg.PortRw, cfg.DBName)
	fmt.Println(dsn)
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		panic(err.Error())
	}
	return db
}

func (d *dao) CreateTestRepl(test *TestRepl) error {
	// Подготовить оператор для вставки данных
	stmtIns, err := d.db.Prepare("INSERT INTO test_repl (id, test) VALUES (?, ?)") // ? = заполнитель
	if err != nil {
		return err // правильная обработка ошибок вместо паники
	}
	defer func() { _ = stmtIns.Close() }() // Закрывается оператор, когда выйдете из функции

	id, err := test.Id.MarshalBinary()
	_, err = stmtIns.Exec(id, test.Test)
	if err != nil {
		return err
	}
	return nil
}

func gracefulClose(db *sql.DB) {
	// Настраиваем канал для отправки сигнальных уведомлений.
	// Нужно использовать буферизованный канал или есть риск пропустить сигнал
	// если не готовы принять сигнал при отправке.
	c := make(chan os.Signal, 1)

	// Блокировать до получения сигнала.
	s := <-c
	fmt.Println("Got signal:", s)
	err := db.Close()
	if err != nil {
		log.Println(err)
	}
}

func test(upperBound int64) {

	gofakeit.Seed(0)

	// Загрузка конфигурации
	var envFile string
	flag.StringVar(&envFile, "env-file", ".env", "Read in a file of environment variables")
	flag.Parse()
	err := godotenv.Load(envFile)

	if err != nil {
		logger.Debug("main: can't load configuration")
	}
	environ, err := config.Environ()

	if environ == nil {
		panic("main: can't load configuration")
	}
	// Если логгирование на уровне трассировки включено, вывести
	// параметры конфигурации.
	if environ.Logging.Debug {
		fmt.Println(environ.String())
		logger.SetLevel(logger.DEBUG)
	}

	db := openDB(environ.DataBase)
	go gracefulClose(db)
	dao := dao{db: db}

	var i = 0
	defer func() {
		if recover() != nil {
			fmt.Printf("error count = %d\n", i)
		} else {
			fmt.Printf("ok count = %d\n", i)
		}
	}()
	for ; i < int(upperBound); i++ {

		var t TestRepl
		gofakeit.Struct(&t)

		id := uuid.New()
		t.Id = id

		err := dao.CreateTestRepl(&t)
		if err != nil {
			panic(err)
		}
		fmt.Println(t.String())
	}
}

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

	// Зарегистрировать login маршрут.
	s.POST("/login", h.Login)

	// Зарегистрировать маршрут для профиля пользователя.
	s.GET("/profile", h.Profile)

	// Зарегистрировать маршрут для списка пользователей.
	s.GET("/users/all", h.List)

	// Зарегистрировать маршруты для поиска пользователей.
	s.GET("/users/search/{name}/{surname}", h.Search)
	s.GET("/users/search-by/{field}/{value}", h.SearchBy)

	// Зарегистрировать маршрут для списка пользователей.
	s.GET("/user/{id}", h.User)

	// Зарегистрировать маршрут для создания пользователя.
	s.POST("/user", h.Create)

	// Зарегистрировать маршрут для добавления друга.
	s.POST("/friend", h.Friend)

	// Зарегистрировать маршрут для Sign-in пользователя.
	s.POST("/signin", h.SignIn)

	// Зарегистрировать маршрут для Sign-in пользователя.
	s.POST("/save", h.Save)

	// Run
	if err := s.ListenAndServe(); err != nil {
		panic(err)
	}
}

func main() {

	upperBound, err := strconv.ParseInt(os.Args[1], 10, 32)

	if err == nil {
		test(upperBound)
	} else {
		httpd()
	}
}
