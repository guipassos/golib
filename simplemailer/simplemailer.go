//go:generate mockgen -source=${GOFILE} -package=${GOPACKAGE} -destination=${GOPACKAGE}_mock.go
package simplemailer

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"net/smtp"
	"path/filepath"
	"strings"

	"github.com/guipassos/golib/logger"
)

var (
	ErrValidateSmtpUser     = errors.New("[mailer] invalid value to smtp user")
	ErrValidateSmtpPassword = errors.New("[mailer] invalid value to smtp password")
	ErrValidateSmtpHost     = errors.New("[mailer] invalid value to smtp host")
	ErrValidateSmtpPort     = errors.New("[mailer] invalid value to smtp port")
	ErrValidateFromName     = errors.New("[mailer] invalid value to from name")
	ErrValidateFromEmail    = errors.New("[mailer] invalid value to from email")
	ErrValidateTemplatePath = errors.New("[mailer] invalid value to templates path")
)

type Mailer interface {
	Validate() error
	Send(to []string, subject, filename string, data map[string]interface{}) error
	SendAttachment(to []string, subject, templateFile string, data map[string]interface{}, attachments map[string][]byte) error
	SetTemplateSettings(templatePath, templateBase string, settings map[string]string)
}

type mailerImpl struct {
	smtpUser     string
	smtpPassword string
	smtpHost     string
	smtpPort     string
	fromName     string
	fromEmail    string
	templatePath string
	templateBase string
	settings     map[string]string
}

func NewMailer(smtpUser, smtpPassword, smtpHost, smtpPort, fromName, fromEmail string) Mailer {
	return &mailerImpl{
		smtpUser:     smtpUser,
		smtpPassword: smtpPassword,
		smtpHost:     smtpHost,
		smtpPort:     smtpPort,
		fromName:     fromName,
		fromEmail:    fromEmail,
		templatePath: ".",
		templateBase: "base.html",
	}
}

func (m *mailerImpl) SetTemplateSettings(templatePath, templateBase string, settings map[string]string) {
	m.templatePath = templatePath
	m.templateBase = templateBase
	m.settings = settings
}

func (m mailerImpl) Validate() error {
	if m.smtpUser == "" {
		logger.Error(ErrValidateSmtpUser)
		return ErrValidateSmtpUser
	}
	if m.smtpPassword == "" {
		logger.Error(ErrValidateSmtpPassword)
		return ErrValidateSmtpPassword
	}
	if m.smtpHost == "" {
		logger.Error(ErrValidateSmtpHost)
		return ErrValidateSmtpHost
	}
	if m.smtpPort == "" {
		logger.Error(ErrValidateSmtpPort)
		return ErrValidateSmtpPort
	}
	if m.fromName == "" {
		logger.Error(ErrValidateFromName)
		return ErrValidateFromName
	}
	if m.fromEmail == "" {
		logger.Error(ErrValidateFromEmail)
		return ErrValidateFromEmail
	}
	return nil
}

func (m mailerImpl) Send(to []string, subject, filename string, data map[string]interface{}) error {
	message := fmt.Sprintf("From: %s <%s>\n", m.fromName, m.fromEmail)
	message += "To: " + strings.Join(to, ", ") + "\n"
	message += "Subject: " + subject + "\n"
	message += "Content-Type: multipart/alternative; boundary=\"-\"\n\n"
	message += "---\n"
	message += "Content-Type: text/html; charset=\"utf-8\"\n"
	message += "Content-Disposition: inline\n\n"
	abs, err := filepath.Abs(m.templatePath)
	if err != nil {
		logger.Error("file abs error. ", err.Error())
		return err
	}
	data["settings"] = m.settings
	tmplt, err := template.ParseFiles(abs+"/"+filename, abs+"/"+m.templateBase)
	if err != nil {
		logger.Error("template ", abs+"/"+filename, " not found ", err.Error())
		return err
	}
	var buffer bytes.Buffer
	if err = tmplt.Execute(&buffer, data); err != nil {
		logger.Error("template execute error. ", err.Error())
		return err
	}
	message += buffer.String()
	message += "-----"
	go func() {
		auth := smtp.PlainAuth("", m.smtpUser, m.smtpPassword, m.smtpHost)
		if err := smtp.SendMail(m.smtpHost+":"+m.smtpPort, auth, m.fromEmail, to, []byte(message)); err != nil {
			logger.Error("unable to send email to: ", to, err.Error())
			return
		}
	}()
	return nil
}

func (m mailerImpl) SendAttachment(
	to []string,
	subject,
	templateFile string,
	data map[string]interface{},
	attachments map[string][]byte,
) error {
	message := fmt.Sprintf("From: %s <%s>\n", m.fromName, m.fromEmail)
	message += "To: " + strings.Join(to, ", ") + "\n"
	message += "Subject: " + subject + "\n"
	message += "MIME-Version: 1.0\n"
	message += "Content-Type: multipart/mixed; boundary=\"MAIN\"\n\n"
	message += "--MAIN\n"
	message += "Content-Type: text/html; charset=\"utf-8\"\n"
	message += "Content-Disposition: inline\n"
	message += "Content-Transfer-Encoding: quoted-printable\n\n"
	abs, err := filepath.Abs(".")
	if err != nil {
		logger.Error("file abs error. ", err.Error())
		return err
	}
	tmplt, err := template.ParseFiles(abs+"/"+templateFile, abs+"/templates/base.html")
	if err != nil {
		logger.Error("template ", abs+"/"+templateFile, " not found. ", err.Error())
		return err
	}
	data["settings"] = m.settings
	var buffer bytes.Buffer
	if err = tmplt.Execute(&buffer, data); err != nil {
		logger.Error("template execute error. ", err.Error())
		return err
	}
	message += buffer.String() + "\n\n"
	go func() {
		auth := smtp.PlainAuth("", m.smtpUser, m.smtpPassword, m.smtpHost)
		if len(attachments) > 0 {
			for fileName, attachment := range attachments {
				message += "--MAIN\n"
				message += fmt.Sprintf("Content-Type: %s; name=\"%s\" \n", http.DetectContentType(attachment), fileName)
				message += fmt.Sprintf("Content-Description: \"%s\"\n", fileName)
				message += fmt.Sprintf("Content-Disposition: attachment; filename=%s \n", fileName)
				message += "Content-Transfer-Encoding: base64\n"
				message += base64.StdEncoding.EncodeToString(attachment) + "\n\n"
			}
			message += "--MAIN--\n"
		}
		if err := smtp.SendMail(m.smtpHost+":"+m.smtpPort, auth, m.fromEmail, to, []byte(message)); err != nil {
			logger.Error("unable to send email to: ", to, err.Error())
			return
		}
	}()
	return nil
}
