package dao

import (
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app/model"
)

func (i *interest) Create(interest *model.Interest) error {
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
