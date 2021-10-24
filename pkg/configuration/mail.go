package configuration

type MailConfig struct {
	SmtpHost, Name, Email, Password string
	SmtpPort                        int
}
