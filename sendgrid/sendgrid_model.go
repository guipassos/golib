package sendgrid

import "github.com/sendgrid/sendgrid-go/helpers/mail"



type SendRequest struct {
	Template string
	To       Emails
	Data     DynamicData
}

func (s *SendRequest) SetDefaultTemplateLanguage() {
	switch s.Data["lang"] {
	case TEMPLATE_LANGUAGE_ENUS:
	case TEMPLATE_LANGUAGE_PTBR:
	default:
		s.Data["lang"] = TEMPLATE_LANGUAGE_ENUS
	}
}

type (
	Templates   map[string]string
	DynamicData map[string]interface{}
)

func (t Templates) CheckIDExists(id string) bool {
	if _, ok := t[id]; ok {
		return true
	}
	return false
}

func (t Templates) GetID(key string) string {
	if id, ok := t[key]; ok {
		return id
	}
	return ""
}

type (
	Email struct {
		Address string
		Name    string
	}
	Emails []Email
)

func (e Emails) List() (list []*mail.Email) {
	for _, val := range e {
		if val.Name == "" {
			val.Name = val.Address
		}
		list = append(list, mail.NewEmail(val.Name, val.Address))
	}
	return list
}
