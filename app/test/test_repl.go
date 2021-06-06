package test

import (
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/savsgio/go-logger/v2"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app/config"
	"log"
	"os"
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

func FakeIt(upperBound int64) {

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
