package domain

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/savsgio/go-logger/v2"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app/set"
	"strings"
)

type Interest struct {
	id        uuid.UUID
	Interests string
}

func (i *Interest) Id() uuid.UUID {
	return i.id
}

func (i *Interest) SetId(id uuid.UUID) {
	i.id = id
}

func (i *Interest) String() string {
	return string(i.Marshal())
}

func (i *Interest) Marshal() []byte {

	interest, err := json.Marshal(*i)
	if err != nil {
		logger.Error(err)
		return nil
	}
	return interest
}

func (i *interest) Create(interest *Interest) error {
	// Подготовить оператор для вставки данных
	stmtIns, err := i.db.Prepare("INSERT INTO interest (id, interests) VALUES (?, ?)") // ? = заполнитель
	if err != nil {
		return err // правильная обработка ошибок вместо паники
	}
	defer func() { _ = stmtIns.Close() }() // Закрывается оператор, когда выйдете из функции

	id, err := interest.Id().MarshalBinary()

	if err != nil {
		return err
	}
	_, err = stmtIns.Exec(id, interest.Interests)
	if err != nil {
		return err
	}
	return nil
}

func (i *interest) CreateInterests(interests []string) error {

	err := i.createInterests(interests)
	if err != nil {
		return err
	}
	return nil
}

func (i *interest) GetExistsInterests(interests []string) ([]Interest, error) {

	stmtOut, err := i.db.Prepare(`
		SELECT id, interests FROM interest WHERE JSON_CONTAINS(?, JSON_ARRAY(interests));
	`)
	if err != nil {
		return nil, err // правильная обработка ошибок вместо паники
	}
	defer func() { _ = stmtOut.Close() }()

	ar := `["` + strings.Join(interests, `", "`) + `"]`
	rows, err := stmtOut.Query(ar)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	var ids []Interest
	for rows.Next() {

		var interest Interest
		err = rows.Scan(&interest.id, &interest.Interests)

		if err != nil {
			continue
		}
		ids = append(ids, interest)
	}
	return ids, nil
}

func (i *interest) createInterests(interests []string) error {

	newInterests, err := i.extractNewInterests(interests)
	if err != nil {
		return err // правильная обработка ошибок вместо паники
	}
	for _, interest := range newInterests {
		if logger.DebugEnabled() {
			logger.Debugf("interest: %v", interest)
		}
		newInterest := Interest{id: uuid.New(), Interests: interest}
		err := i.Create(&newInterest)

		if err != nil {
			return err
		}
	}
	return nil
}

func (i *interest) extractNewInterests(interests []string) ([]string, error) {

	existsInterests, err := i.getExistsInterestLabels(interests)
	if err != nil {
		return nil, err // правильная обработка ошибок вместо паники
	}
	if logger.DebugEnabled() {
		logger.Debugf("existsInterests: %v", existsInterests)
	}
	result := set.DifferenceString(interests, existsInterests)
	if logger.DebugEnabled() {
		logger.Debugf("result: %v", result)
	}
	return result, nil
}

func (i *interest) getExistsInterestLabels(interests []string) ([]string, error) {

	stmtOut, err := i.db.Prepare(`SELECT interests FROM interest WHERE JSON_CONTAINS(?, JSON_ARRAY(interests))`)

	if err != nil {
		return nil, err // правильная обработка ошибок вместо паники
	}
	defer func() { _ = stmtOut.Close() }()

	ar := `["` + strings.Join(interests, `", "`) + `"]`
	rows, err := stmtOut.Query(ar)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	var is []string
	for rows.Next() {
		var i string

		err = rows.Scan(&i)
		if err != nil {
			continue
		}
		is = append(is, i)
	}
	return is, nil
}
