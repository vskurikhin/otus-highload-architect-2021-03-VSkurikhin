package dao

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app/model"
	"strings"
)

func (u *user) Create(user *model.User) error {
	// Подготовить оператор для вставки данных
	stmtIns, err := u.db.Prepare(`
		INSERT INTO user
			(id, username, name, surname, age, sex, city)
		  VALUES
			(?, ?, ?, ?, ?, ?, ?)`, // ? = заполнитель
	)
	if err != nil {
		return err // правильная обработка ошибок вместо паники
	}
	defer func() { _ = stmtIns.Close() }() // Закрывается оператор, когда выйдете из функции

	id, err := user.Id().MarshalBinary()
	_, err = stmtIns.Exec(id, user.Username, user.Name, user.SurName, user.Age, user.Sex, user.City)
	if err != nil {
		return err
	}
	return nil
}

func (u *user) ReadListAsString() (string, error) {

	users, err := u.readList()
	if err != nil {
		return "{}", err // правильная обработка ошибок вместо паники
	}
	var result []string

	for _, user := range users {
		result = append(result, user.String())
	}
	return "[" + strings.Join(result, ", ") + "]", nil
}

func (u *user) readList() ([]model.User, error) {

	stmtOut, err := u.db.Prepare(`
		SELECT u.id, username, name, surname, age, sex, city, JSON_ARRAYAGG(interests)
		  FROM user u
		  JOIN user_has_interests uhi ON u.id = uhi.user_id
		  JOIN interest i ON i.id = uhi.interest_id
		 GROUP BY u.id, username, name, surname, age, sex, city
	`)
	if err != nil {
		return nil, err // правильная обработка ошибок вместо паники
	}
	defer func() { _ = stmtOut.Close() }()

	rows, err := stmtOut.Query()
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	var users []model.User
	for rows.Next() {
		var id uuid.UUID
		var u model.User
		var interests string

		err = rows.Scan(&id, &u.Username, &u.Name, &u.SurName, &u.Age, &u.Sex, &u.City, &interests)
		if err != nil {
			return nil, err
		}
		u.SetId(id)

		err = json.Unmarshal([]byte(interests), &u.Interests)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}
