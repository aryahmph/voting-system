package service

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
	"strings"
	"voting-system/model/domain"
	"voting-system/pkg/configuration"
	"voting-system/pkg/exception"
)

type MailServiceImpl struct {
	MailConfig configuration.MailConfig
}

func NewMailServiceImpl(mailConfig configuration.MailConfig) *MailServiceImpl {
	return &MailServiceImpl{MailConfig: mailConfig}
}

func (service *MailServiceImpl) Send(sendMail domain.SendMail) {
	body := "From: " + service.MailConfig.Name + "\n" +
		"To: " + strings.Join(sendMail.To, ",") + "\n" +
		"Subject: " + sendMail.Subject + "\n" + "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n" +
		sendMail.Message

	auth := smtp.PlainAuth("", service.MailConfig.Email, service.MailConfig.Password, service.MailConfig.SmtpHost)
	smtpAddr := fmt.Sprintf("%s:%d", service.MailConfig.SmtpHost, service.MailConfig.SmtpPort)

	smtp.SendMail(smtpAddr, auth, service.MailConfig.Email, sendMail.To, []byte(body))
}

func (service *MailServiceImpl) ParseTemplate(fileName string, data domain.TemplateMail) string {
	parseFiles, err := template.ParseFiles(fileName)
	exception.PanicIfError(err)

	buff := new(bytes.Buffer)
	err = parseFiles.Execute(buff, data)
	exception.PanicIfError(err)

	return buff.String()
}