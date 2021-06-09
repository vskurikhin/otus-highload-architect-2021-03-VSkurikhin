package domain

import (
	"database/sql"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/savsgio/go-logger/v2"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app/cache"
)

type News struct {
	id       *uuid.UUID
	Title    string
	Content  string
	PublicAt string
	Username string
}

func (l *News) Id() *uuid.UUID {
	return l.id
}

func (l *News) NewId() {
	id := uuid.New()
	l.id = &id
}

func (l *News) SetId(id *uuid.UUID) {
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
	return &News{id: &id, Title: c.Title, Content: c.Content, PublicAt: c.PublicAt, Username: c.Username}
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

const COUNT_MY_NEWS_LIST = `
    SELECT COUNT(*) size
      FROM news n
      WHERE n.username IN (
        SELECT r.username
          FROM ` + "`user`" + ` r
         WHERE r.id = ?
        UNION
        SELECT u.username
          FROM user_has_friends uhf
          JOIN ` + "`user`" + ` u ON u.id = uhf.friend_id
         WHERE uhf.user_id = ?)
      ORDER BY n.public_at DESC
      LIMIT ? OFFSET ?`

func (u *news) SizeMyNewsList(p *Profile, offset, limit int) (int64, error) {

	stmtOut, err := u.dbRw.Prepare(COUNT_MY_NEWS_LIST)
	if err != nil {
		return 0, err // правильная обработка ошибок вместо паники
	}
	defer func() { _ = stmtOut.Close() }()

	id, err := p.Id.MarshalBinary()
	if err != nil {
		return 0, err // правильная обработка ошибок вместо паники
	}
	var result int64
	err = stmtOut.QueryRow(id, id, limit, offset).Scan(&result)
	if err != nil {
		return 0, err // правильная обработка ошибок вместо паники
	}
	return result, nil
}

const LAST_MY_NEWS = `
    SELECT n.id, n.title, n.content, n.public_at, n.username, BIN_TO_UUID(n.id) uid
      FROM news n
      WHERE n.username IN (
        SELECT r.username
          FROM ` + "`user`" + ` r
         WHERE r.id = ?
        UNION
        SELECT u.username
          FROM user_has_friends uhf
          JOIN ` + "`user`" + ` u ON u.id = uhf.friend_id
         WHERE uhf.user_id = ?)
      ORDER BY n.public_at DESC, uid ASC
      LIMIT 1 OFFSET 0`

func (u *news) LastMyNews(p *Profile) (*News, error) {

	var n News
	stmtOut, err := u.dbRw.Prepare(LAST_MY_NEWS)
	if err != nil {
		return nil, err // правильная обработка ошибок вместо паники
	}
	defer func() { _ = stmtOut.Close() }()

	id, err := p.Id.MarshalBinary()
	if err != nil {
		return nil, err // правильная обработка ошибок вместо паники
	}
	var uid string
	err = stmtOut.QueryRow(id, id).Scan(&n.id, &n.Title, &n.Content, &n.PublicAt, &n.Username, &uid)
	if err != nil {
		return nil, err // правильная обработка ошибок вместо паники
	}
	return &n, nil
}

const SELECT_MY_NEWS_LIST = `
    SELECT n.id, n.title, n.content, n.public_at, n.username
      FROM news n
      WHERE n.username IN (
        SELECT r.username
          FROM ` + "`user`" + ` r
         WHERE r.id = ?
        UNION
        SELECT u.username
          FROM user_has_friends uhf
          JOIN ` + "`user`" + ` u ON u.id = uhf.friend_id
         WHERE uhf.user_id = ?)
      ORDER BY n.public_at DESC
      LIMIT ? OFFSET ?`

func (u *news) ReadMyNewsList(p *Profile, offset, limit int) ([]News, error) {

	stmtOut, err := u.dbRo.Prepare(SELECT_MY_NEWS_LIST)
	if err != nil {
		return nil, err // правильная обработка ошибок вместо паники
	}
	defer func() { _ = stmtOut.Close() }()

	return readMyNewsList(stmtOut, p, offset, limit)
}

func readMyNewsList(stmtOut *sql.Stmt, p *Profile, offset, limit int) ([]News, error) {

	var result []News

	id, err := p.Id.MarshalBinary()
	if err != nil {
		return nil, err // правильная обработка ошибок вместо паники
	}
	rows, err := stmtOut.Query(id, id, limit, offset)
	logger.Debugf("readMyNewsList %v", rows)

	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	for rows.Next() {
		var n News
		err = rows.Scan(&n.id, &n.Title, &n.Content, &n.PublicAt, &n.Username)
		if err != nil {
			return nil, err
		}
		result = append(result, n)
	}
	return result, nil
}

/*
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

const SELECT_FRIENDS_NEWS_LIST = `
    SELECT n.id, n.title, n.content, n.public_at, n.username
      FROM news n
      WHERE n.username IN (
        SELECT u.username
          FROM user_has_friends uhf
          JOIN ` + "`user`" + ` u ON u.id = uhf.friend_id
         WHERE uhf.user_id = ?)
      ORDER BY n.public_at DESC
      LIMIT ? OFFSET ?`

func (u *news) ReadFriendsNewsListOld(offset, limit int, id uuid.UUID) ([]News, error) {

	stmtOut, err := u.dbRw.Prepare(SELECT_FRIENDS_NEWS_LIST)
	if err != nil {
		return nil, err // правильная обработка ошибок вместо паники
	}
	defer func() { _ = stmtOut.Close() }()

	i, err := id.MarshalBinary()
	if err != nil {
		return nil, err // правильная обработка ошибок вместо паники
	}
	rows, err := stmtOut.Query(i, limit, offset)
	if err != nil {
		return nil, err // правильная обработка ошибок вместо паники
	}

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
*/
