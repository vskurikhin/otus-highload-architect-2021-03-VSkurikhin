package kafka

import (
	"github.com/Shopify/sarama"
	"github.com/savsgio/go-logger/v2"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app-counter/srv"
)

// Consumer представляет потребителя группы потребителей
type Consumer struct {
	ready   chan bool
	service *srv.Service
}

func NewConsumer(srv *srv.Service) *Consumer {
	consumer := Consumer{
		ready:   make(chan bool),
		service: srv,
	}
	return &consumer
}

// Setup запускается в начале нового сеанса, прежде чем ConsumeClaim
func (consumer *Consumer) Setup(sarama.ConsumerGroupSession) error {
	// Mark the consumer as ready
	close(consumer.ready)
	return nil
}

// Cleanup запускается в конце сеанса, когда все горутины ConsumeClaim завершились
func (consumer *Consumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim должен запустить потребительский цикл сообщений ConsumerGroupClaim's Messages().
func (consumer *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {

	// NOTE:
	// Не переносите приведенный ниже код в goroutine.
	// `ConsumeClaim` сам вызывается внутри горутины, см.:
	// https://github.com/Shopify/sarama/blob/master/consumer_group.go#L27-L29
	for message := range claim.Messages() {
		err := consumer.service.Do(message)
		if err == nil {
			session.MarkMessage(message, "")
		} else {
			// Записываем ERROR в лог
			logger.Error(err)
		}
	}
	return nil
}
