package sns

import "github.com/aws/aws-sdk-go/service/sns"

func (s snsImpl) Publish(input PublishInput) (*PublishOutput, error) {
	res, err := s.sns.Publish((*sns.PublishInput)(&input))
	return (*PublishOutput)(res), err
}
