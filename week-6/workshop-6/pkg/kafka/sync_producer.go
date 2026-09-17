package kafka

import (
	"github.com/Shopify/sarama"
	"github.com/pkg/errors"
)

func NewSyncProducer(brokers []string, opts ...Option) (sarama.SyncProducer, error) {
	config := prepareProducerSaramaConfig(opts...)

	syncProducer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		return nil, errors.Wrap(err, "error with sync kafka-producer")
	}

	return syncProducer, nil
}
