//go:generate mockgen -source=${GOFILE} -package=${GOPACKAGE} -destination=${GOPACKAGE}_mock.go
package sendgrid

import (
	"errors"

	"github.com/guipassos/golib/logger"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type (
	Options struct {
		ApiKey      string
		From        Email
		DynamicData DynamicData
		Templates   Templates
	}
	Sendgrid interface {
		Send(send SendRequest) (err error)
	}
	sendgridImpl struct {
		apiKey      string
		host        string
		from        *mail.Email
		client      *sendgrid.Client
		templates   Templates
		dynamicData DynamicData
	}
)

func New(opt Options) Sendgrid {
	return &sendgridImpl{
		apiKey:      opt.ApiKey,
		from:        mail.NewEmail(opt.From.Name, opt.From.Address),
		client:      sendgrid.NewSendClient(opt.ApiKey),
		dynamicData: opt.DynamicData,
		templates:   opt.Templates,
	}
}

func (s sendgridImpl) Send(send SendRequest) (err error) {
	v3Mail := mail.NewV3Mail()
	v3Mail.SetFrom(s.from)
	if !s.templates.CheckIDExists(send.Template) {
		return errors.New("template not specified")
	}
	go func() {
		send.SetDefaultTemplateLanguage()
		v3Mail.SetTemplateID(s.templates.GetID(send.Template))
		p := mail.NewPersonalization()
		p.AddTos(send.To.List()...)
		for key, value := range s.dynamicData {
			p.SetDynamicTemplateData(key, value)
		}
		for key, value := range send.Data {
			p.SetDynamicTemplateData(key, value)
		}
		v3Mail.AddPersonalizations(p)
		response, err := s.client.Send(v3Mail)
		if err != nil {
			logger.Error("sendmail request error", err)
			return
		}
		if response.StatusCode < 200 || response.StatusCode > 299 {
			logger.Error("sendmail response with fail, details: ", response)
		}
	}()
	return nil
}
