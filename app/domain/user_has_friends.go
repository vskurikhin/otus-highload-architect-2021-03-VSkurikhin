package domain

import (
	"github.com/google/uuid"
)

type Friend struct {
	UserId   string
	FriendId string
}

func (u *userHasFriends) Link(user *User, friend *User) error {
	// Подготовить оператор для вставки данных
	stmtIns, err := u.db.Prepare(`
		INSERT INTO user_has_friends (id, user_id, friend_id) VALUES (?, ?, ?)
	`) // ? = заполнитель
	if err != nil {
		return err // правильная обработка ошибок вместо паники
	}
	defer func() { _ = stmtIns.Close() }() // Закрывается оператор, когда выйдете из функции

	id, _ := uuid.New().MarshalBinary()
	userId, _ := user.Id().MarshalBinary()
	friendId, _ := friend.Id().MarshalBinary()

	_, err = stmtIns.Exec(id, userId, friendId)
	if err != nil {
		return err
	}
	return nil
}

func (u *userHasFriends) LinkFriends(user *User, friends []User) error {
	for _, friend := range friends {
		err := u.Link(user, &friend)
		if err != nil {
			return err
		}
	}
	return nil
}
