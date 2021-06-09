package domain

import (
	"errors"
	"github.com/google/uuid"
	"github.com/savsgio/go-logger/v2"
)

type Friend struct {
	UserId   string
	FriendId string
}

func (u *userHasFriends) LinkToFriend(user *User, friend *User) error {

	userId, err := user.Id().MarshalBinary()

	if err != nil {
		return err
	}
	friendId, err := friend.Id().MarshalBinary()

	if err != nil {
		return err
	}
	f, err := u.isFriendship(userId, friendId)

	if err != nil {
		return err
	} else if f {
		return errors.New("friendship exists")
	}
	// Подготовить оператор для вставки данных
	stmtIns, err := u.dbRw.Prepare(`INSERT INTO user_has_friends (id, user_id, friend_id) VALUES (?, ?, ?)`)
	if err != nil {
		return err // правильная обработка ошибок вместо паники
	}
	defer func() { _ = stmtIns.Close() }() // Закрывается оператор, когда выйдете из функции

	id, _ := uuid.New().MarshalBinary()

	_, err = stmtIns.Exec(id, userId, friendId)
	if err != nil {
		return err
	}
	return nil
}

func (u *userHasFriends) IsFriendship(userId *uuid.UUID, friendId *uuid.UUID) (bool, error) {

	uId, err := userId.MarshalBinary()

	if err != nil {
		return false, err
	}
	fId, err := friendId.MarshalBinary()

	if err != nil {
		return false, err
	}
	if logger.DebugEnabled() {
		logger.Debugf("userId: %s, friendId: %s", userId, friendId)
	}
	return u.isFriendship(uId, fId)
}

const SELECT_COUNT_FROM_USER_HAS_FRIENDS_WHERE_USER_ID_AND_FRIEND_ID = `
	SELECT COUNT(*)
	  FROM user_has_friends
	 WHERE user_id = ? AND friend_id = ?`

func (u *userHasFriends) isFriendship(userId []byte, friendId []byte) (bool, error) {

	stmtOut, err := u.dbRo.Prepare(SELECT_COUNT_FROM_USER_HAS_FRIENDS_WHERE_USER_ID_AND_FRIEND_ID)
	if err != nil {
		return false, err // правильная обработка ошибок вместо паники
	}
	defer func() { _ = stmtOut.Close() }()

	var count int
	err = stmtOut.QueryRow(userId, friendId).Scan(&count)

	if err != nil {
		return false, err
	}
	if logger.DebugEnabled() {
		logger.Debugf("count %d", count)
	}
	return count > 0, nil
}

func (u *userHasFriends) LinkFriends(user *User, friends []User) error {
	for _, friend := range friends {
		err := u.LinkToFriend(user, &friend)
		if err != nil {
			return err
		}
	}
	return nil
}
