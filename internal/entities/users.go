package entities

import (
	"time"
)

type User struct {
	UserID          string
	Email           *string
	Phone           *string
	EmailEncrypted  *[]byte
	PhoneEncrypted  *[]byte
	EmailSearchable *[]byte
	PhoneSearchable *[]byte
	Password        []byte
	FirstName       string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       *string
}

func (u *User) Deleted() bool {
	return u.DeletedAt != nil
}

type UserCreateSpec struct {
	AuthMethodSpec
	Password    []byte
	FirstName   string
	CountryCode string
}

type AuthMethodSpec struct {
	Type   AuthMethodType
	Values AuthMethodValues
}

type UsersFindOneSpec struct {
	Email   *[]byte
	Phone   *[]byte
	UserID  *string
	Deleted bool
}

type UsersFindManySpec struct {
	Email         *[]byte
	Phone         *[]byte
	CreatedAtFrom *time.Time
	CreatedAtTo   *time.Time
	Pagination    *Pagination
	Deleted       bool
}

type PrivateDataSpec struct {
	Encrypted  []byte
	Searchable []byte
}

type UsersPatchDataSpec struct {
	Email     *PrivateDataSpec
	Phone     *PrivateDataSpec
	Password  *[]byte
	FirstName *string
	LastName  *string
	Deleted   *bool
}

type UsersPatchOneSpec struct {
	Email  *[]byte
	Phone  *[]byte
	UserID *string
	Data   UsersPatchDataSpec
}
