package kafka

import (
	"context"
	"github.com/Shopify/sarama"
	"github.com/savsgio/go-logger/v2"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app-counter/config"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app-counter/utils"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type ConsumerGroup struct {
	Ctx           context.Context
	Cancel        context.CancelFunc
	ConsumerGroup sarama.ConsumerGroup
}

func NewConsumerGroup(environ *config.Config, config *sarama.Config) (*ConsumerGroup, error) {

	ctx, cancel := context.WithCancel(context.Background())
	client, err := sarama.NewConsumerGroup(split(environ.Kafka.Brokers), environ.Kafka.Group, config)
	utils.PanicCheck(err)
	cg := &ConsumerGroup{
		Ctx:           ctx,
		Cancel:        cancel,
		ConsumerGroup: client,
	}
	return cg, err
}

func (g *ConsumerGroup) WaitConsumerGroup(topics string, consumer *Consumer) {

	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		for {
			// `Consume` следует вызывать внутри бесконечного цикла, когда
			// происходит ребалансировка на стороне сервера, сеанс потребителя должен быть
			// воссоздан, чтобы получить новые claims
			if err := g.ConsumerGroup.Consume(g.Ctx, split(topics), consumer); err != nil {
				log.Panicf("Error from consumer: %v", err)
			}
			// проверить, не был ли контекст отменен, сигнализируя о том, что потребитель должен остановиться
			if g.Ctx.Err() != nil {
				return
			}
			consumer.ready = make(chan bool)
		}
	}()

	<-consumer.ready // Подождать, пока потребитель будет настроен
	logger.Info(" consumer up and running!... ")

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-g.Ctx.Done():
		logger.Info(" terminating: context cancelled ")
	case <-sigterm:
		logger.Info(" terminating: via signal ")
	}
	g.Cancel()
	wg.Wait()

	if err := g.ConsumerGroup.Close(); err != nil {
		logger.Error(" Error closing client: %v ", err)
		panic("Error closing client")
	}
}
