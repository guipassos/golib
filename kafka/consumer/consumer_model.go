package consumer

import (
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/rhizomplatform/golib/kafka/kafkamod"
)

type (
	Options struct {
		Configs            map[string]interface{}
		Topics             []string
		ReadMessageTimeout time.Duration
		ReadMessageDelay   time.Duration
	}
	Response struct {
		Error   error
		Message *kafkamod.Message
	}
)

func (o *Options) getConfigMap() (kafka.ConfigMap, error) {
	var cm = kafka.ConfigMap{}
	for key, val := range o.Configs {
		if err := cm.SetKey(key, val); err != nil {
			return nil, err
		}
	}
	return cm, nil
}
