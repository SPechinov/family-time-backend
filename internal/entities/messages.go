package entities

import "errors"

type MessageMethod string

const (
	MessageMethodEmail MessageMethod = "email"
	MessageMethodPhone MessageMethod = "phone"
)

func (m MessageMethod) IsValid() error {
	switch m {
	case MessageMethodEmail, MessageMethodPhone:
		return nil
	}
	return errors.New("smtp method is not valid")
}
