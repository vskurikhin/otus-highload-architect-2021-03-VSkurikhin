package domain

import (
	"github.com/google/uuid"
)

func (u *userHasInterests) Link(user *User, interest *Interest) error {
	// Подготовить оператор для вставки данных
	stmtIns, err := u.db.Prepare(`
		INSERT INTO user_has_interests (id, user_id, interest_id) VALUES (?, ?, ?)
	`) // ? = заполнитель
	if err != nil {
		return err // правильная обработка ошибок вместо паники
	}
	defer func() { _ = stmtIns.Close() }() // Закрывается оператор, когда выйдете из функции

	id, _ := uuid.New().MarshalBinary()
	userId, _ := user.Id().MarshalBinary()
	interestId, _ := interest.Id().MarshalBinary()

	_, err = stmtIns.Exec(id, userId, interestId)
	if err != nil {
		return err
	}
	return nil
}

func (u *userHasInterests) LinkInterests(user *User, interests []Interest) error {
	for _, interest := range interests {
		err := u.Link(user, &interest)
		if err != nil {
			return err
		}
	}
	return nil
}
