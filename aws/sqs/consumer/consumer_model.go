package consumer

import (
	"github.com/aws/aws-sdk-go/service/sqs"
)

type (
	Response struct {
		Error       error
		Message     *Message
		MessageBody *MessageBody
	}
	Message     sqs.Message
	MessageBody struct {
		Type           string `json:"Type"`
		TopicARN       string `json:"TopicArn"`
		Message        string `json:"Message"`
		MessageID      string `json:"MessageId"`
		Timestamp      string `json:"Timestamp"`
		UnsubscribeURL string `json:"UnsubscribeURL"`
	}
)
