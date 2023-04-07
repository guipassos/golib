package kafkamod

import (
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Message struct {
	TopicPartition TopicPartition
	Value          []byte
	Key            []byte
	Headers        []Header
	Timestamp      time.Time
	Metadata       *string
}

func (m Message) KafkaHeaders() []kafka.Header {
	headers := make([]kafka.Header, len(m.Headers))
	for i, h := range m.Headers {
		headers[i] = kafka.Header{
			Key:   h.Key,
			Value: h.Value,
		}
	}
	return headers
}

func (m Message) KafkaTopicPartition() (tp kafka.TopicPartition) {
	tp.Topic = &m.TopicPartition.Topic
	tp.Partition = m.TopicPartition.Partition
	tp.Metadata = m.Metadata
	tp.Offset.Set(m.TopicPartition.Offset)
	return tp
}
