package composites

import (
	"server/internal/config"
	"server/internal/pkg/email_sender"
	"server/internal/services/message_sender"
)

func NewMessageSender(cfg *config.Config) *message_sender.MessageSender {
	emailSender := email_sender.New(cfg)
	return message_sender.New(cfg, emailSender)
}
