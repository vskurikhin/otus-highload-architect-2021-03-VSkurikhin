package domain

import (
	"github.com/google/uuid"
	"github.com/savsgio/go-logger/v2"
	"strings"
)

func (u *userHasInterests) Link(user *User, interest *Interest) error {
	// Подготовить оператор для вставки данных
	stmtIns, err := u.dbRw.Prepare(`INSERT INTO user_has_interests (id, user_id, interest_id) VALUES (?, ?, ?)`)
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

const SELECT_ALL_INTERESTS_CONTAINS_IN_SET = `
    SELECT i.interests
      FROM user_has_interests uhi
      JOIN interest i ON i.id = uhi.interest_id
     WHERE uhi.user_id = ?
       AND JSON_CONTAINS(?, JSON_ARRAY(i.interests))`

const DELETE_ALL_INTERESTS_CONTAINS_NOT_IN_SET = `
    DELETE FROM user_has_interests
     WHERE user_id = ?
       AND interest_id NOT IN (
           SELECT i.id
             FROM interest i
            WHERE JSON_CONTAINS(?, JSON_ARRAY(i.interests)))`

func (u *userHasInterests) LinkedInterestMap(userId uuid.UUID, is *InterestMap) (*InterestMap, error) {

	id, err := userId.MarshalBinary()
	if err != nil {
		return nil, err // правильная обработка ошибок вместо паники
	}

	stmtOut1, err := u.dbRo.Prepare(SELECT_ALL_INTERESTS_CONTAINS_IN_SET)

	if err != nil {
		return nil, err // правильная обработка ошибок вместо паники
	}
	defer func() { _ = stmtOut1.Close() }()

	ar := `["` + strings.Join(is.SavedSetLabels(), `", "`) + `"]`
	rows, err := stmtOut1.Query(id, ar)

	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	for rows.Next() {
		var i string

		err = rows.Scan(&i)
		if err != nil {
			continue
		}
		delete(is.SavedMap, i)
	}
	stmtOut2, err := u.dbRw.Prepare(DELETE_ALL_INTERESTS_CONTAINS_NOT_IN_SET)

	if err != nil {
		return nil, err // правильная обработка ошибок вместо паники
	}
	defer func() { _ = stmtOut2.Close() }()
	count, err := stmtOut2.Query(id, ar)

	if logger.DebugEnabled() {
		logger.Debugf("LinkedInterestMap: count: %d", count)
	}
	return is, nil
}
