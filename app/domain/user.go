package domain

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/savsgio/go-logger/v2"
	"strings"
)

type User struct {
	id        uuid.UUID
	Username  string
	Name      *string
	SurName   *string
	Age       int
	Sex       int
	Interests []string
	City      *string
}

func Create(id uuid.UUID, username string, name *string, surname *string, age int, sex int, inters []string, city *string,
) *User {
	return &User{
		id:        id,
		Username:  username,
		Name:      name,
		SurName:   surname,
		Age:       age,
		Sex:       sex,
		Interests: inters,
		City:      city,
	}
}

func NewUser(username string) *User {
	return &User{id: uuid.New(), Username: username}
}

func Setup(username string, name *string, surname *string, age int, sex int, inters []string, city *string) *User {
	return &User{
		Username:  username,
		Name:      name,
		SurName:   surname,
		Age:       age,
		Sex:       sex,
		Interests: inters,
		City:      city,
	}
}

func (u *User) Id() uuid.UUID {
	return u.id
}

func (u *User) SetId(id uuid.UUID) {
	u.id = id
}

func (u *User) StringInterests() string {
	if len(u.Interests) > 0 {
		return `["` + strings.Join(u.Interests, `", "`) + `"]`
	}
	return ""
}

func (u *User) String() string {
	return string(u.Marshal())
}

func (u *User) Marshal() []byte {

	user, err := json.Marshal(struct {
		Id        uuid.UUID
		Username  string
		Name      *string
		SurName   *string
		Age       int
		Sex       int
		Interests []string
		City      *string
	}{
		Id:        u.id,
		Username:  u.Username,
		Name:      u.Name,
		SurName:   u.SurName,
		Age:       u.Age,
		Sex:       u.Sex,
		Interests: u.Interests,
		City:      u.City,
	})
	if err != nil {
		logger.Errorf("%v", err)
		return nil
	}
	return user
}

func (u *user) Create(user *User) error {
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

func (u *user) Read(sid string) (string, error) {

	user, err := u.ReadUser(sid)

	if err != nil {
		return "{}", err
	}
	return user.String(), nil
}

func (u *user) ReadUser(sid string) (*User, error) {

	id, err := uuid.Parse(sid)

	if err != nil {
		return nil, err // правильная обработка ошибок вместо паники
	}
	user, err := u.read(id)

	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *user) read(id uuid.UUID) (*User, error) {

	stmtOut, err := u.db.Prepare(`
		SELECT u.id, username, name, surname, age, sex, city, JSON_ARRAYAGG(interests)
		  FROM user u
		  JOIN user_has_interests uhi ON u.id = uhi.user_id
		  JOIN interest i ON i.id = uhi.interest_id
		 WHERE u.id = ?
         GROUP BY u.id, username, name, surname, age, sex, city
	`)
	if err != nil {
		return nil, err // правильная обработка ошибок вместо паники
	}
	defer func() { _ = stmtOut.Close() }()

	var user User
	var interests string
	idBytes, err := id.MarshalBinary()
	err = stmtOut.QueryRow(idBytes).
		Scan(&user.id, &user.Username, &user.Name, &user.SurName, &user.Age, &user.Sex, &user.City, &interests)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(interests), &user.Interests)

	if err != nil {
		return nil, err
	}
	return &user, nil
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

func (u *user) readList() ([]User, error) {

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

	var users []User
	for rows.Next() {

		var user User
		var interests string

		err = rows.Scan(&user.id, &user.Username, &user.Name, &user.SurName, &user.Age, &user.Sex, &user.City, &interests)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal([]byte(interests), &user.Interests)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
