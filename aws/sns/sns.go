//go:generate mockgen -source=${GOFILE} -package=${GOPACKAGE} -destination=${GOPACKAGE}_mock.go
package sns

import (
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/guipassos/golib/aws/session"
)

type (
	Options struct {
		Session session.Session
	}
	SNS interface {
		Publish(input PublishInput) (*PublishOutput, error)
	}
	snsImpl struct {
		sns *sns.SNS
	}
)

func New(opts Options) SNS {
	return snsImpl{
		sns: sns.New(opts.Session.GetSession()),
	}
}
