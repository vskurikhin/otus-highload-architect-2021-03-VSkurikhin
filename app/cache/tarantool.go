package cache

import (
	"fmt"
	"github.com/savsgio/go-logger/v2"
	"github.com/tarantool/go-tarantool"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app/config"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app/domain"
	"reflect"
	"strconv"
)

const MAX_UINT32 = ^uint32(0)

type Tarantool struct {
	Cache *tarantool.Connection
}

func CreateTarantoolCacheClient(cfg *config.Config) *Tarantool {
	client, err := newTarantoolClient(cfg)
	if err != nil {
		panic(err.Error())
	}
	return &Tarantool{Cache: client}
}

func newTarantoolClient(cfg *config.Config) (*tarantool.Connection, error) {
	opts := tarantool.Opts{
		User: cfg.Cache.Username,
	}
	addr := fmt.Sprintf("%s:%d", cfg.Cache.Host, cfg.Cache.Port)
	conn, err := tarantool.Connect(addr, opts)
	if err != nil {
		panic(err)
	}
	return conn, nil
}

func (t *Tarantool) UserIndexNameSearch(name, surname string) ([]domain.User, error) {
	list, err := t.userIndexNameSearch(name, surname)
	if err != nil {
		return nil, err
	}
	var users []domain.User
	for _, value := range list {
		object := reflect.ValueOf(value)
		var user domain.User
		for index := 0; index < object.Len(); index++ {
			record := fmt.Sprintf("%v", object.Index(index).Interface())
			switch index {
			case 0:
				id, err := strconv.ParseUint(record, 10, 64)
				if err != nil {
					return nil, err
				}
				user.Id = id
			case 1:
				user.Name = &record
			case 2:
				user.Username = record
			case 3:
				user.SurName = &record
			case 4:
				age, err := strconv.ParseInt(record, 10, 64)
				if err != nil {
					return nil, err
				}
				user.Age = int(age)
			case 5:
				sex, err := strconv.ParseInt(record, 10, 64)
				if err != nil {
					return nil, err
				}
				user.Sex = int(sex)
			case 6:
				user.City = &record
			default:
				logger.Errorf("default record: %v", record)
			}
		}
		users = append(users, user)
	}
	return users, nil
}

func (t *Tarantool) userIndexNameSearch(name, surname string) ([]interface{}, error) {

	resp, err := t.Cache.Call("user_index_name_search", []interface{}{name, surname})
	if err != nil {
		logger.Errorf("Error: %s", err)
		if resp != nil {
			return resp.Data, err
		}
		return nil, err
	}
	return resp.Data, nil
}
