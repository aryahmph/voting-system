package service

import "voting-system/model/domain"

type MailService interface {
	Send(mail domain.SendMail)
	ParseTemplate(fileName string, data domain.TemplateMail) string
}
