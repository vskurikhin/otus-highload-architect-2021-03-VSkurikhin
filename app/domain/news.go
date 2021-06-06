package domain

import (
	"database/sql"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/savsgio/go-logger/v2"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app/cache"
)

type News struct {
	id       uuid.UUID
	Title    string
	Content  string
	PublicAt string
	Username string
}

func (l *News) Id() uuid.UUID {
	return l.id
}

func (l *News) SetId(id uuid.UUID) {
	l.id = id
}

func NewsConvert(d *News) *cache.News {
	return &cache.News{Id: d.Id().String(), Title: d.Title, Content: d.Content, PublicAt: d.PublicAt, Username: d.Username}
}

func ConvertNews(c *cache.News) *News {
	id, err := uuid.Parse(c.Id)
	if err != nil {
		id = uuid.New()
	}
	return &News{id: id, Title: c.Title, Content: c.Content, PublicAt: c.PublicAt, Username: c.Username}
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
	stmtIns, err := l.dbRw.Prepare("INSERT INTO news (id, title, content, public_at, username) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return err // правильная обработка ошибок вместо паники
	}
	defer func() { _ = stmtIns.Close() }() // Закрывается оператор, когда выйдете из функции

	id, err := news.Id().MarshalBinary()
	_, err = stmtIns.Exec(id, news.Title, news.Content, news.PublicAt, news.Username)
	if err != nil {
		return err
	}

	return nil
}

const SELECT_NEWS = `
    SELECT id, title, content, public_at, username
      FROM news
      WHERE id = ?`

func (u *news) ReadNews(id uuid.UUID) (*News, error) {

	i, err := id.MarshalBinary()
	stmtOut, err := u.dbRo.Prepare(SELECT_NEWS)
	if err != nil {
		return nil, err // правильная обработка ошибок вместо паники
	}
	defer func() { _ = stmtOut.Close() }()

	var n News
	err = stmtOut.QueryRow(i).
		Scan(&n.id, &n.Title, &n.Content, &n.PublicAt)
	if err != nil {
		return nil, err
	}

	return &n, nil
}

const SELECT_NEWS_LIST = `
    SELECT id, title, content, public_at, username
      FROM news
      ORDER BY public_at DESC
      LIMIT ? OFFSET ?`

func (u *news) ReadNewsList(offset, rowcount int) ([]News, error) {

	stmtOut, err := u.dbRo.Prepare(SELECT_NEWS_LIST)
	if err != nil {
		return nil, err // правильная обработка ошибок вместо паники
	}
	defer func() { _ = stmtOut.Close() }()

	return readNewsList(stmtOut, offset, rowcount)
}

func (u *news) RefreshNewsList(offset, rowcount int) ([]News, error) {

	stmtOut, err := u.dbRw.Prepare(SELECT_NEWS_LIST)
	if err != nil {
		return nil, err // правильная обработка ошибок вместо паники
	}
	defer func() { _ = stmtOut.Close() }()

	return readNewsList(stmtOut, offset, rowcount)
}

func readNewsList(stmtOut *sql.Stmt, offset, rowcount int) ([]News, error) {

	rows, err := stmtOut.Query(rowcount, offset)

	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	var newsSet []News
	for rows.Next() {

		var n News
		err = rows.Scan(&n.id, &n.Title, &n.Content, &n.PublicAt, &n.Username)
		if err != nil {
			return nil, err
		}
		newsSet = append(newsSet, n)
	}
	return newsSet, nil
}
