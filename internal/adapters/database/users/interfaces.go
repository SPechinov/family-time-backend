package users

import "time"

type userData struct {
	UserID          string
	EmailEncrypted  *[]byte
	PhoneEncrypted  *[]byte
	EmailSearchable *[]byte
	PhoneSearchable *[]byte
	Password        []byte
	FirstName       string
	LastName        *string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       *time.Time
}
