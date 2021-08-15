[![hits](https://hits.deltapapa.io/github/vskurikhin/otus-highload-architect-2021-03-VSkurikhin.svg)](https://hits.deltapapa.io)
[![license](https://img.shields.io/github/license/vskurikhin/otus-highload-architect-2021-03-VSkurikhin)](https://raw.githubusercontent.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/main/LICENSE)
modify_column_name
# OTUS Highload Architect

Группа 2021-03

Виктор Скурихин (Victor Skurikhin)

- [Лекции](doc/lectures.md)
  1. [Проблемы высоких нагрузок](doc/lectures.md#проблемы-высоких-нагрузок)
- [Домашние задания](doc/homeworks.md)
  1. [Заготовка для социальной сети](doc/solutions_of_homework.md#заготовка-для-социальной-сети)
     - [Заготовка для социальной сети](doc/solutions_of_homework.md#заготовка-для-социальной-сети)
     - [Производительность индексов](doc/solutions_of_homework.md#производительность-индексов)
     - [Лента новостей социальной сети](solutions_of_homework.md#лента-новостей-социальной-сети)
     - [Полусинхронная репликация](doc/solutions_of_homework.md#полусинхронная-репликация)
     - [In-Memory СУБД](doc/solutions_of_homework.md#in-memory-субд)
- [Разное](doc/other.md)

# Установка

- [Установить](https://github.com/pressly/goose) :\
  `go get -u github.com/pressly/goose/cmd/goose`
- [Скачать репозиторий](https://github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin.git)
- [Создать БД, пользователя и раздать привелегии](db/create/create.sql) :\
  `cat ./db/create/create.sql | mysql`
- Выполнить скрипты миграции БД:\
  `goose -v -dir ./db/migrations mysql "hl:password@/hl?parseTime=true" up`
- При необходимости обновиить React:
  ```
  npm install
  npm run dev
  ```
  Собранное React приложение уже присутвует в репозитории: `app-bundle.js`.\
  Клиентское JavaScript приложение отрабатывает в браузере клиента.  
- Выкачать зависимости:\
  `go mod download`
- `go run .`

# Библиотеки

- https://github.com/dgrijalva/jwt-go
- https://github.com/google/uuid
- https://github.com/joho/godotenv
- https://github.com/kelseyhightower/envconfig
- https://github.com/savsgio/atreugo/v11
- https://github.com/savsgio/go-logger/v2
- https://github.com/valyala/fasthttp
- https://golang.org/x/crypto/bcrypt
- https://golang.org/x/sync
- https://gopkg.in/yaml.v2
