package main

import (
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/hashicorp/consul/api"
	"github.com/joho/godotenv"
	"github.com/savsgio/go-logger/v2"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app-dialog/config"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app-dialog/server"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app-dialog/server/handlers"
	"log"
	"net"
	"os"
	"strconv"
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

	s.GET("/messages", h.GetMessages)
	s.POST("/message", h.PostMessage)
	s.PUT("/message", h.PutMessage)

	hostname, err := os.Hostname()
	if err != nil {
		logger.Error(err)
	}
	addr, err := net.LookupIP(hostname)
	if err != nil {
		fmt.Println("Unknown host")
	} else {
		fmt.Println("IP address: ", addr)
	}
	port, err := strconv.ParseInt(environ.Server.Port, 10, 64)
	if err != nil {
		logger.Error(err)
	}

	service := &api.AgentServiceRegistration{
		ID:      "my-app-dialog",
		Name:    "my-app-dialog",
		Port:    int(port),
		Address: addr[0].String(),
		Check: &api.AgentServiceCheck{
			HTTP:     "http://" + addr[0].String() + ":" + fmt.Sprintf("%d", port),
			Interval: "5s",
			Timeout:  "1s",
		},
	}
	client, _ := api.NewClient(api.DefaultConfig())

	if err := client.Agent().ServiceRegister(service); err != nil {
		log.Fatal(err)
	}

	// Run
	if err := s.ListenAndServe(); err != nil {
		panic(err)
	}
}
