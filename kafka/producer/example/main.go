package example

import (
	"encoding/json"

	"github.com/guipassos/golib/kafka/kafkamod"
	"github.com/guipassos/golib/kafka/producer"
	"github.com/guipassos/golib/logger"
)

const TopicExample = "EXAMPLE"

func main() {
	p, err := producer.New(producer.Options{
		Configs: map[string]interface{}{
			"bootstrap.servers": "localhost:9092",
			"acks":              "all",
		},
	})
	if err != nil {
		logger.Fatal("failed to create producer ", err)
	}
	msg, err := exampleMessage()
	if err != nil {
		logger.Fatal("failed to generate msg ", err)
	}
	if err = p.Produce(*msg); err != nil {
		logger.Fatal("failed to produce msg ", err)
	}
	chDelivery := make(chan producer.Delivery)
	go p.Delivery(chDelivery)
	for delivery := range chDelivery {
		if delivery.Error != nil {
			logger.Error("DELIVERY ERROR: ", delivery.Error)
			continue
		}
		logger.Info("DELIVERY: ", delivery.TopicPartition)
	}
	p.Close()
}

func exampleMessage() (*kafkamod.Message, error) {
	msgBytes, err := json.Marshal(map[string]string{
		"name":  "xpik",
		"email": "contact@xpik.me",
	})
	if err != nil {
		logger.Error("failed to convert json ", err)
		return nil, err
	}
	msgHeader := []kafkamod.Header{
		{Key: "Content-Type", Value: []byte("application/json")},
	}
	msg := kafkamod.Message{
		TopicPartition: kafkamod.TopicPartition{
			Topic: TopicExample,
		},
		Value:   msgBytes,
		Headers: msgHeader,
	}
	return &msg, nil
}
