package email_sender

import (
	"fmt"
	"net/smtp"
	"server/internal/config"
	"server/pkg/app_error"
)

type EmailSender struct {
	address       string
	fromEmail     string
	fromEmailName string
}

func New(cfg *config.Config) *EmailSender {
	return &EmailSender{
		address:       cfg.SMTP.Host + ":" + cfg.SMTP.Port,
		fromEmail:     cfg.SMTP.FromEmail,
		fromEmailName: cfg.SMTP.FromEmailName,
	}
}

func (es *EmailSender) SendRegMessage(recipient string, code string) error {
	message := es.buildRegMessage(code)
	result := es.emailBuilder("Confirmation code", recipient, message)

	err := smtp.SendMail(es.address, nil, es.fromEmail, []string{recipient}, []byte(result))
	if err != nil {
		return app_error.New(err)
	}

	return nil
}

func (es *EmailSender) SendForgotPasswordMessage(recipient string, code string) error {
	message := es.buildForgotPasswordMessage(code)
	result := es.emailBuilder("Forgot password confirmation code", recipient, message)

	err := smtp.SendMail(es.address, nil, es.fromEmail, []string{recipient}, []byte(result))
	if err != nil {
		return app_error.New(err)
	}

	return nil
}

func (es *EmailSender) emailBuilder(subject, recipient, content string) string {
	email := fmt.Sprintf("From: %s <%s>\n", es.fromEmailName, es.fromEmail)
	email += fmt.Sprintf("To: %s\n", recipient)
	email += fmt.Sprintf("Subject: %s\n", subject)
	email += "Content-Type: text/html; charset=\"utf-8\"\n\n" // Добавлено: заголовок для HTML
	email += content

	return email
}
