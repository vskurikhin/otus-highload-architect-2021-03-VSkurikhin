package main

import (
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/savsgio/go-logger/v2"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app-counter/config"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app-counter/domain"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app-counter/kafka"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app-counter/server"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app-counter/server/handlers"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app-counter/srv"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app-counter/utils"
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
	dao := domain.New(environ)
	go httpd(environ, dao)
	consumer(environ, dao)
}

func httpd(environ *config.Config, dao *domain.DAO) {

	// Создать инстанс сервера
	s := server.New(environ, dao)

	// Зарегистрируйте для аутентификации перед обработкой запросов.
	s.UseBefore(s.JWT.AuthCheckToken)

	// Обработка статичных фалов из каталога web/public.
	s.StaticCustom()

	// Обработчики запросов.
	h := handlers.Handlers{Server: s}

	// Зарегистрировать индексный маршрут.
	s.GET("/", h.Root)

	s.GET("/counter/{username}", h.GetCounter)

	// Run
	if err := s.ListenAndServe(); err != nil {
		panic(err)
	}
}

func consumer(environ *config.Config, dao *domain.DAO) {
	// kafka.StartConsumer(tlsConfig, dao, environ)
	// Конфигурация подключения к Kafka для sarama
	kafkaConfig := kafka.NewConsumerConfig()
	// Создание Consumer group-ы
	consumerGroup, err := kafka.NewConsumerGroup(environ, kafkaConfig)
	utils.PanicCheck(err)

	consumer := kafka.NewConsumer(srv.NewService(dao))
	consumerGroup.WaitConsumerGroup(environ.Kafka.Topic, consumer)
}
