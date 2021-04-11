[![hits](https://hits.deltapapa.io/github/vskurikhin/otus-highload-architect-2021-03-VSkurikhin.svg)](https://hits.deltapapa.io)

# Домашняя работа

## TOC

- [Решение домашних заданий](solutions_of_homework.md)
  1. [Заготовка для социальной сети](solutions_of_homework.md#заготовка-для-социальной-сети)

### [Заготовка для социальной сети](solutions_of_homework.md#заготовка-для-социальной-сети)

Заготовка для социальной сети
Цель:
В результате выполнения ДЗ вы создадите базовый скелет социальной сети, который будет развиваться в дальнейших ДЗ. В данном задании тренируются навыки:

1. Проведена декомпозиции предметной области:\
   созданы таблицы:
   - login;
   - user;
   - interest;
   - user_has_interests.
1. Построенна архитектура проекта, состоящая из:\
   - БД (MySQL);
   - монолитного приложения (otus-highload-architect-2021-03-VSkurikhin);
   - UI на стороне web-браузера (JavaScript + React).

Требуется разработать создание и просмотр анект в социальной сети.

Реализованы функциональные требования:

1. Авторизация по паролю (API /login).
1. Страница регистрации (API /signin).
1. Страницa со списком пользователей (API /userlist).
1. Страницы с анкетой (API /userform/:id).

Нефункциональные требования:
- [x] Любой язык программирования
- [x] В качестве базы данных использовать MySQL
- [x] Не использовать ORM
- [x] Программа должна представлять из себя монолитное приложение.

Не рекомендуется использовать следующие технологии:

- [-] Репликация
- [-] Шардинг
- [-] Индексы
- [-] Кэширование

- [x] Разместить приложение на любом хостинге. Например, heroku:\
  [highload-architect](https://highload-architect.herokuapp.com)
- [x] ДЗ принимается в виде исходного кода на github и демонстрации проекта на хостинге:\
  бранч [homework/1](https://github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/tree/homework/1) на [github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin](https://github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin)

- [ ] Оценка происходит по принципу зачет/незачет.
