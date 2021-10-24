package domain

type SendMail struct {
	To               []string
	Subject, Message string
}

type TemplateMail struct {
	Name  string
	Token string
}
