//go:generate mockgen -source=${GOFILE} -package=${GOPACKAGE} -destination=${GOPACKAGE}_mock.go
package session

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

type (
	Options struct {
		Region          string
		AccessKeyID     string
		SecretAccessKey string
	}
	Session interface {
		GetSession() *session.Session
	}
	sessionImpl struct {
		awsSession *session.Session
	}
)

func New(opts Options) (Session, error) {
	credentials := credentials.NewCredentials(
		&credentials.StaticProvider{
			Value: credentials.Value{
				AccessKeyID:     opts.AccessKeyID,
				SecretAccessKey: opts.SecretAccessKey,
			},
		},
	)
	configs := []*aws.Config{
		aws.NewConfig().WithCredentials(credentials),
		aws.NewConfig().WithRegion(opts.Region),
	}
	sess, err := session.NewSession(configs...)
	if err != nil {
		return nil, err
	}
	return sessionImpl{awsSession: sess}, nil
}
