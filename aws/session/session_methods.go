package session

import "github.com/aws/aws-sdk-go/aws/session"

func (s sessionImpl) GetSession() *session.Session {
	return s.awsSession
}
