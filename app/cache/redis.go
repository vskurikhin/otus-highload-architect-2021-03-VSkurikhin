package cache

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/savsgio/go-logger/v2"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app/config"
)

var ctx = context.Background()

type Redis struct {
	Cache *redis.Client
}

func CreateRedisCacheClient(cfg *config.Config) *Redis {
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
		_ = client.Close()
		return nil, err
	}
	return client, nil
}

func (r *Redis) SCardNews(sk *SetKey) (int64, error) {
	intCmd := r.Cache.SCard(ctx, sk.Key)
	if intCmd.Err() != nil {
		logger.Errorf("SCard error %v", intCmd.Err())
		return -1, intCmd.Err()
	}
	return intCmd.Val(), nil
}

func (r *Redis) PutNews(p *Profile, key *NewsKey, news *News) {
	sk := CreateSetKey(p)
	logger.Debugf("cache.PutNews: profile: %s, key: %s,  sk: %s, news: %s", p, key, sk, news)
	err := r.Cache.Set(ctx, key.Key, news.String(), redis.KeepTTL).Err()
	if err != nil {
		logger.Errorf("cache.PutNews: r.Cache.Set: %v", err)
	}
	err = r.Cache.SAdd(ctx, sk.Key, key.Key).Err()
	if err != nil {
		logger.Errorf("cache.PutFriendNews: r.Cache.SAdd: %v", err)
	}
}

func (r *Redis) LastMyNewsKey(p *Profile) (*NewsKey, error) {
	sk := CreateSetKey(p)
	logger.Debugf("cache.LastMyNews: sk: %s", sk)
	cmd := r.Cache.Sort(ctx, sk.Key, &redis.Sort{Offset: 0, Count: 1, Order: "DESC", Alpha: true})
	if cmd != nil && cmd.Err() != nil {
		logger.Errorf("cache.LastMyNewsKey: %v", cmd.Err())
	}
	if len(cmd.Val()) != 1 {
		return nil, errors.New(fmt.Sprintf("Empty or not valid cache for sk: %s", sk))
	}
	nk := NewsKey{Key: cmd.Val()[0]}
	return &nk, nil
}

func (r *Redis) SortMyNewsKeys(p *Profile, offset, limit int64) []string {
	sk := CreateSetKey(p)
	logger.Debugf("cache.SortMyNewsKeys: sk: %s, offset: %d, limit: %d", sk, offset, limit)
	cmd := r.Cache.Sort(ctx, sk.Key, &redis.Sort{Offset: offset, Count: limit, Order: "DESC", Alpha: true})
	if cmd != nil && cmd.Err() != nil {
		logger.Errorf("cache.SortMyNewsKeys: %v", cmd.Err())
	}
	return cmd.Val()
}

//
//func (r *Redis) PutNews(news *News) {
//	if news == nil {
//		return
//	}
//	key := fmt.Sprintf("News|%s|%s", news.PublicAt, news.Id)
//	logger.Debugf("PutNews: key: %s", key)
//	err := r.Cache.Set(ctx, key, news.String(), redis.KeepTTL).Err()
//	if err != nil {
//		logger.Errorf("PutNews: 1: %v", err)
//	}
//	err = r.Cache.SAdd(ctx, "News", key).Err()
//	if err != nil {
//		logger.Errorf("PutNews: 2: %v", err)
//	}
//}

func (r *Redis) SAdd(id uuid.UUID, key string) {
	k := fmt.Sprintf("News|%s", id)
	err := r.Cache.SAdd(ctx, k, key).Err()
	if err != nil {
		logger.Errorf("cache.SAdd: 2: %v", err)
	}
}

func (r *Redis) PutFriendNews(news *News, id *uuid.UUID) {
	if news == nil {
		return
	}
	key := fmt.Sprintf("News|%s|%s", news.PublicAt, news.Id)
	logger.Debugf("PutNews: key: %s", key)
	err := r.Cache.Set(ctx, key, news.String(), redis.KeepTTL).Err()
	if err != nil {
		logger.Errorf("cache.PutFriendNews: 1: %v", err)
	}
	k := fmt.Sprintf("News|%s", id)
	err = r.Cache.SAdd(ctx, k, key).Err()
	if err != nil {
		logger.Errorf("cache.PutFriendNews: 2: %v", err)
	}
}

//func (r *Redis) SortNewsKeys(offset, limit int64) []string {
//	logger.Debugf("cache.SortNewsKeys: %d, %d", offset, limit)
//	cmd := r.Cache.Sort(ctx, "News", &redis.Sort{Offset: offset, Count: limit, Order: "DESC", Alpha: true})
//	if cmd != nil && cmd.Err() != nil {
//		logger.Errorf("cache.SortNewsKeys: %v", cmd.Err())
//	}
//	return cmd.Val()
//}

func (r *Redis) SortFriendsNewsKeys(offset int64, limit int64, id uuid.UUID) []string {
	key := fmt.Sprintf("News|%s", id)
	logger.Debugf("cache.SortFriendsNewsKeys: %d, %d", offset, limit)
	cmd := r.Cache.Sort(ctx, key, &redis.Sort{Offset: offset, Count: limit, Order: "DESC", Alpha: true})
	if cmd != nil && cmd.Err() != nil {
		logger.Errorf("cache.SortFriendsNewsKeys: %v", cmd.Err())
	}
	return cmd.Val()
}

//func (r *Redis) SetOfNews(offset, limit int64) []string {
//	var result []string
//	keys := r.SortNewsKeys(offset, limit)
//	for _, key := range keys {
//		cmd := r.Cache.Get(ctx, key)
//		if cmd != nil && cmd.Err() != nil {
//			logger.Errorf("SetOfNews: key: %s %v", key, cmd.Err())
//		} else {
//			result = append(result, cmd.Val())
//		}
//	}
//	return result
//}

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

func (r *Redis) SCardFriendsNews(id uuid.UUID) *redis.IntCmd {
	key := fmt.Sprintf("News|%s", id)
	return r.Cache.SCard(ctx, key)
}
