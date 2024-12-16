package message_sender

import (
	"server/internal/config"
	"server/internal/entities"
	"server/pkg/app_error"
)

type MessageSender struct {
	cfg         *config.Config
	emailSender SenderMessages
}

func New(cfg *config.Config, emailSender SenderMessages) *MessageSender {
	return &MessageSender{
		cfg:         cfg,
		emailSender: emailSender,
	}
}

func (ms *MessageSender) SendRegMessage(sendMethod entities.MessageMethod, recipient string, code string) error {
	if err := sendMethod.IsValid(); err != nil {
		return app_error.New(err)
	}

	return ms.emailSender.SendRegMessage(recipient, code)
}

func (ms *MessageSender) SendForgotPasswordMessage(sendMethod entities.MessageMethod, recipient string, code string) error {
	if err := sendMethod.IsValid(); err != nil {
		return app_error.New(err)
	}

	return ms.emailSender.SendForgotPasswordMessage(recipient, code)
}
