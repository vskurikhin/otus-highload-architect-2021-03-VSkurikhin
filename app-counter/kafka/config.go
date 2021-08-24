package kafka

import (
	"github.com/Shopify/sarama"
	"strings"
)

func NewConsumerConfig() *sarama.Config {

	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin
	config.Net.TLS.Enable = false
	config.Net.SASL.Handshake = false

	return config
}

func split(s string) []string {
	return strings.Split(s, ";")
}
