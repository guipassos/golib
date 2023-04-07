package example

import (
	"github.com/rhizomplatform/golib/kafka/consumer"
	"github.com/rhizomplatform/golib/logger"
)

const TopicExample = "EXAMPLE"

func main() {
	c, err := consumer.New(consumer.Options{
		Topics: []string{TopicExample},
		Configs: map[string]interface{}{
			"bootstrap.servers": "localhost:9092",
			"auto.offset.reset": "earliest",
		},
	})
	if err != nil {
		logger.Fatal("failed to create consumer ", err)
	}
	chResponse := make(chan consumer.Response)
	go c.Run(chResponse)
	for res := range chResponse {
		if res.Error != nil {
			logger.Error("MESSAGE ERROR: ", res.Error)
			continue
		}
		logger.Info("MESSAGE: ", res.Message)
	}
	c.Close()
}
