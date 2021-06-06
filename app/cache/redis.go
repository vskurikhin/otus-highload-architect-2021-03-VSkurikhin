package cache

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/savsgio/go-logger/v2"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app/config"
)

var ctx = context.Background()

type Redis struct {
	Cache *redis.Client
}

func NewRedis(cfg *config.Config) *Redis {
	client, err := newRedisClient(cfg)
	if err != nil {
		panic(err.Error())
	}
	return &Redis{Cache: client}
}

func newRedisClient(cfg *config.Config) (*redis.Client, error) {
	addr := fmt.Sprintf("%s:%d", cfg.Cache.Host, cfg.Cache.Port)
	logger.Debugf(addr)
	client := redis.NewClient(&redis.Options{
		Addr:       addr,
		Username:   cfg.Cache.Username,
		Password:   cfg.Cache.Password,
		DB:         0,
		MaxRetries: 2,
	})

	if _, err := client.Ping(ctx).Result(); err != nil {
		client.Close()
		return nil, err
	}
	return client, nil
}

func (r *Redis) PutNews(news *News) {
	if news == nil {
		return
	}
	key := fmt.Sprintf("News|%s|%s", news.PublicAt, news.Id)
	logger.Debugf("PutNews: key: %s", key)
	err := r.Cache.Set(ctx, key, news.String(), redis.KeepTTL).Err()
	if err != nil {
		logger.Errorf("PutNews: 1: %v", err)
	}
	err = r.Cache.SAdd(ctx, "News", key).Err()
	if err != nil {
		logger.Errorf("PutNews: 2: %v", err)
	}
}

func (r *Redis) SortNewsKeys(offset, limit int64) []string {
	logger.Debugf("SortNewsKeys: %d, %d", offset, limit)
	cmd := r.Cache.Sort(ctx, "News", &redis.Sort{Offset: offset, Count: limit, Order: "DESC", Alpha: true})
	if cmd != nil && cmd.Err() != nil {
		logger.Errorf("SortNewsKeys: %v", cmd.Err())
	}
	return cmd.Val()
}

func (r *Redis) SetOfNews(offset, limit int64) []string {
	var result []string
	keys := r.SortNewsKeys(offset, limit)
	for _, key := range keys {
		cmd := r.Cache.Get(ctx, key)
		if cmd != nil && cmd.Err() != nil {
			logger.Errorf("SetOfNews: key: %s %v", key, cmd.Err())
		} else {
			result = append(result, cmd.Val())
		}
	}
	return result
}

func (r *Redis) GetNewsJson(key string) (*string, error) {
	cmd := r.Cache.Get(ctx, key)
	if cmd != nil && cmd.Err() != nil {
		logger.Errorf("GetNewsJson: key: %s %v", key, cmd.Err())
		return nil, cmd.Err()
	} else {
		result := cmd.Val()
		return &result, nil
	}
}

func (r *Redis) SCardNews() *redis.IntCmd {
	return r.Cache.SCard(ctx, "News")
}
