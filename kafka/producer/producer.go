//go:generate mockgen -source=${GOFILE} -package=${GOPACKAGE} -destination=${GOPACKAGE}_mock.go
package producer

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/rhizomplatform/golib/kafka/kafkamod"
)

type (
	Producer interface {
		Produce(msg kafkamod.Message) error
		Delivery(chDelivery chan Delivery)
		Flush(timeoutMs int) int
		Close()
	}
	producerImpl struct {
		kafka *kafka.Producer
	}
)

func New(opts Options) (p Producer, err error) {
	var (
		config = opts.getConfigMap()
		impl   = producerImpl{}
	)
	if impl.kafka, err = kafka.NewProducer(&config); err != nil {
		return nil, err
	}
	return impl, nil
}

func (p producerImpl) Produce(msg kafkamod.Message) error {
	return p.kafka.Produce(
		&kafka.Message{
			TopicPartition: msg.KafkaTopicPartition(),
			Value:          msg.Value,
			Key:            msg.Key,
			Headers:        msg.KafkaHeaders(),
		}, nil)
}

func (p producerImpl) Delivery(chDelivery chan Delivery) {
	defer close(chDelivery)
	for e := range p.kafka.Events() {
		switch ev := e.(type) {
		case *kafka.Message:
			chDelivery <- Delivery{
				Error: ev.TopicPartition.Error,
				TopicPartition: kafkamod.TopicPartition{
					Topic:     *ev.TopicPartition.Topic,
					Partition: ev.TopicPartition.Partition,
					Offset:    int64(ev.TopicPartition.Offset),
				},
			}
		}
	}
}

func (p producerImpl) Flush(timeoutMs int) int {
	return p.kafka.Flush(timeoutMs)
}

func (p producerImpl) Close() {
	p.kafka.Flush(0)
	p.kafka.Close()
}
