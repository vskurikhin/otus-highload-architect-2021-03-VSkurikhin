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
	Friend    bool
}

func Create(
	id uuid.UUID,
	username string,
	name *string,
	surname *string,
	age int,
	sex int,
	interests []string,
	city *string,
	friend bool,
) *User {
	return &User{
		id:        id,
		Username:  username,
		Name:      name,
		SurName:   surname,
		Age:       age,
		Sex:       sex,
		Interests: interests,
		City:      city,
		Friend:    friend,
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
		Friend    bool
	}{
		Id:        u.id,
		Username:  u.Username,
		Name:      u.Name,
		SurName:   u.SurName,
		Age:       u.Age,
		Sex:       u.Sex,
		Interests: u.Interests,
		City:      u.City,
		Friend:    u.Friend,
	})
	if err != nil {
		logger.Errorf("%v", err)
		return nil
	}
	return user
}

const INSERT_INTO_USER = `
    INSERT INTO user
	  (id, username, name, surname, age, sex, city)
      VALUES
      (?, ?, ?, ?, ?, ?, ?)`

func (u *user) Create(user *User) error {
	// Подготовить оператор для вставки данных
	stmtIns, err := u.db.Prepare(INSERT_INTO_USER) // ? = заполнитель

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

func (u *user) ReadUser(sid string) (*User, error) {

	id, err := uuid.Parse(sid)

	if err != nil {
		return nil, err // правильная обработка ошибок вместо паники
	}
	user, err := u.readUser(id)

	if err != nil {
		return nil, err
	}
	return user, nil
}

const SELECT_USER_JOIN_INTERESTS_WHERE_ID = `
    SELECT u.id, username, name, surname, age, sex, city, JSON_ARRAYAGG(interests)
      FROM user u
      LEFT JOIN user_has_interests uhi ON uhi.user_id = u.id
      LEFT JOIN interest i ON i.id = uhi.interest_id
     WHERE u.id = ?
     GROUP BY u.id, username, name, surname, age, sex, city`

func (u *user) readUser(id uuid.UUID) (*User, error) {

	stmtOut, err := u.db.Prepare(SELECT_USER_JOIN_INTERESTS_WHERE_ID)
	if err != nil {
		return nil, err // правильная обработка ошибок вместо паники
	}
	defer func() { _ = stmtOut.Close() }()

	var user User
	var interests string
	idBytes, err := id.MarshalBinary()
	if err != nil {
		return nil, err
	}
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

const SELECT_USER_JOIN_INTERESTS = `
    SELECT u.id, username, name, surname, age, sex, city, JSON_ARRAYAGG(interests), NOT isnull(uhf.id) AS is_friend 
      FROM user u
      LEFT JOIN user_has_interests uhi ON uhi.user_id = u.id
      LEFT JOIN interest i ON i.id = uhi.interest_id
      LEFT JOIN user_has_friends uhf ON uhf.friend_id = u.id AND uhf.user_id = ?
     GROUP BY u.id, username, name, surname, age, sex, city, uhf.id`

func (u *user) ReadUserList(id uuid.UUID) ([]User, error) {

	idBytes, err := id.MarshalBinary()
	if err != nil {
		return nil, err
	}
	stmtOut, err := u.db.Prepare(SELECT_USER_JOIN_INTERESTS)
	if err != nil {
		return nil, err // правильная обработка ошибок вместо паники
	}
	defer func() { _ = stmtOut.Close() }()

	rows, err := stmtOut.Query(idBytes)

	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	var users []User
	for rows.Next() {

		var r User
		var interests string

		err = rows.Scan(&r.id, &r.Username, &r.Name, &r.SurName, &r.Age, &r.Sex, &r.City, &interests, &r.Friend)

		if err != nil {
			return nil, err
		}
		err = json.Unmarshal([]byte(interests), &r.Interests)

		if err != nil {
			return nil, err
		}
		users = append(users, r)
	}
	return users, nil
}

const UPDATE_USER_WHERE_ID = `
    UPDATE user SET
      name = ?,
      surname = ?,
      age = ?,
      sex = ?,
      city = ?
     WHERE id = ?`

func (u *user) Update(user *User) error {
	// Подготовить оператор для вставки данных
	stmtIns, err := u.db.Prepare(UPDATE_USER_WHERE_ID) // ? = заполнитель

	if err != nil {
		return err // правильная обработка ошибок вместо паники
	}
	defer func() { _ = stmtIns.Close() }() // Закрывается оператор, когда выйдете из функции

	id, err := user.Id().MarshalBinary()
	_, err = stmtIns.Exec(user.Name, user.SurName, user.Age, user.Sex, user.City, id)
	if err != nil {
		return err
	}
	return nil
}
