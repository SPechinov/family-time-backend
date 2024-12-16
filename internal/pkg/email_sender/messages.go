package email_sender

import "fmt"

func (es *EmailSender) buildRegMessage(code string) string {
	return fmt.Sprintf("<html><body>Email - your confirmation code: %s</body></html>", code)
}

func (es *EmailSender) buildForgotPasswordMessage(code string) string {
	return fmt.Sprintf("<html><body>Email - your forgot password confirmation code: %s</body></html>", code)
}
