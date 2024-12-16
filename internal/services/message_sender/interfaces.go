package message_sender

type SenderMessages interface {
	SendRegMessage(recipient string, code string) error
	SendForgotPasswordMessage(recipient string, code string) error
}
