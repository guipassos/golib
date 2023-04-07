package producer

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/rhizomplatform/golib/kafka/kafkamod"
)

type (
	Options struct {
		Configs map[string]interface{}
	}
	Delivery struct {
		TopicPartition kafkamod.TopicPartition
		Error          error
	}
)

func (o *Options) getConfigMap() kafka.ConfigMap {
	var cm = kafka.ConfigMap{}
	for key, val := range o.Configs {
		cm.SetKey(key, val)
	}
	return cm
}
