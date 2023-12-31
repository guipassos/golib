//go:generate mockgen -source=${GOFILE} -package=${GOPACKAGE} -destination=${GOPACKAGE}_mock.go
package consumer

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/guipassos/golib/aws/session"
)

type (
	Options struct {
		QueueURL            *string
		MaxNumberOfMessages *int64
		Session             session.Session
		ReadMessageDelay    time.Duration
	}
	Consumer interface {
		Run(ctx context.Context, chResponse chan Response)
		DeleteMessage(messageReceipt *string) error
	}
	consumerImpl struct {
		queueURL            *string
		maxNumberOfMessages *int64
		sqs                 *sqs.SQS
		readMessageDelay    time.Duration
	}
)

func New(opts Options) Consumer {
	return consumerImpl{
		sqs:                 sqs.New(opts.Session.GetSession()),
		queueURL:            opts.QueueURL,
		maxNumberOfMessages: opts.MaxNumberOfMessages,
		readMessageDelay:    opts.ReadMessageDelay,
	}
}
