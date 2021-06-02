package domain

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/savsgio/go-logger/v2"
)

type News struct {
	id       uuid.UUID
	Title    string
	Content  string
	PublicAt string
}

func (l *News) Id() uuid.UUID {
	return l.id
}

func (l *News) SetId(id uuid.UUID) {
	l.id = id
}

func (n *News) Marshal() []byte {

	login, err := json.Marshal(*n)
	if err != nil {
		logger.Error(err)
		return nil
	}
	return login
}

func (n *News) String() string {
	return string(n.Marshal())
}

func (l *news) Create(news *News) error {
	// Подготовить оператор для вставки данных
	stmtIns, err := l.dbRw.Prepare("INSERT INTO news (id, title, content, public_at) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err // правильная обработка ошибок вместо паники
	}
	defer func() { _ = stmtIns.Close() }() // Закрывается оператор, когда выйдете из функции

	id, err := news.Id().MarshalBinary()
	_, err = stmtIns.Exec(id, news.Title, news.Content, news.PublicAt)
	if err != nil {
		return err
	}

	return nil
}

const SELECT_NEWS = `
    SELECT id, title, content, public_at 
      FROM news
      ORDER BY public_at DESC
      LIMIT ? OFFSET ?`

func (u *news) ReadNewsList(offset, rowcount int) ([]News, error) {

	stmtOut, err := u.dbRo.Prepare(SELECT_NEWS)
	if err != nil {
		return nil, err // правильная обработка ошибок вместо паники
	}
	defer func() { _ = stmtOut.Close() }()

	rows, err := stmtOut.Query(rowcount, offset)

	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	var newsSet []News
	for rows.Next() {

		var n News
		err = rows.Scan(&n.id, &n.Title, &n.Content, &n.PublicAt)
		if err != nil {
			return nil, err
		}
		newsSet = append(newsSet, n)
	}
	return newsSet, nil
}
