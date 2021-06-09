package pubsub

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/savsgio/go-logger/v2"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app/queue"
)

var ctx = context.Background()

type Redis struct {
	PubSub *redis.Client
}

func NewRedis(client *redis.Client) *Redis {
	return &Redis{PubSub: client}
}

func (r *Redis) Publish(channel string, message interface{}) {
	cmd := r.PubSub.Publish(ctx, channel, message)
	if cmd.Err() != nil {
		logger.Errorf("pubsub.Redis.Publish: %v", cmd.Err())
	}
}

func (r *Redis) Subscribe(channel string, blockingQueue *queue.BlockingQueue) {

	pubsub := r.PubSub.Subscribe(ctx, channel)
	ch := pubsub.Channel()

	for msg := range ch {
		if blockingQueue.Closed() {
			logger.Debugf("blockingQueue: %s closed", blockingQueue)
			return
		}
		blockingQueue.Put(msg.Payload)
	}
}
