package users

import (
	"crypto/sha512"
	"golang.org/x/crypto/bcrypt"
	"server/internal/config"
	"server/internal/entities"
	"server/pkg/app_error"
	"server/pkg/crypto"
)

type Users struct {
	cfg      *config.Config
	database database
}

func New(cfg *config.Config, database database) *Users {
	return &Users{
		cfg:      cfg,
		database: database,
	}
}

func (u *Users) FindOne(spec entities.UsersFindOneSpec) (*entities.User, error) {
	spec.Email = u.hashPersonalInfoValuePtr(spec.Email)
	spec.Phone = u.hashPersonalInfoValuePtr(spec.Phone)

	return u.database.FindOne(spec)
}

func (u *Users) Exists(spec entities.UsersFindOneSpec) (bool, error) {
	user, err := u.FindOne(spec)
	if err != nil {
		return false, err
	}

	return user != nil, nil
}

func (u *Users) FindMany(spec entities.UsersFindManySpec) ([]entities.User, error) {
	spec.Email = u.hashPersonalInfoValuePtr(spec.Email)
	spec.Phone = u.hashPersonalInfoValuePtr(spec.Phone)

	return u.database.FindMany(spec)
}

func (u *Users) Create(spec entities.UserCreateSpec) (*entities.User, error) {
	hashedPassword, err := u.hashPassword(spec.Password)
	if err != nil {
		return nil, err
	}

	encryptedAuthMethod, err := u.encryptPersonalInfoValue(spec.AuthMethodSpec.Values.Encrypted)
	if err != nil {
		return nil, err
	}

	spec.Password = hashedPassword
	spec.AuthMethodSpec.Values.Encrypted = encryptedAuthMethod
	spec.AuthMethodSpec.Values.Searchable = u.hashPersonalInfoValue(spec.AuthMethodSpec.Values.Searchable)

	return u.database.Create(spec)
}

func (u *Users) PatchOne(spec entities.UsersPatchOneSpec) (*entities.User, error) {
	if spec.Email != nil {
		spec.Email = u.hashPersonalInfoValuePtr(spec.Email)
	}

	if spec.Phone != nil {
		spec.Phone = u.hashPersonalInfoValuePtr(spec.Phone)
	}

	if spec.Data.Email != nil {
		encryptedEmail, err := u.encryptPersonalInfoValue(spec.Data.Email.Encrypted)
		if err != nil {
			return nil, err
		}
		spec.Data.Email.Encrypted = encryptedEmail
		spec.Data.Email.Searchable = u.hashPersonalInfoValue(spec.Data.Email.Searchable)
	}

	if spec.Data.Phone != nil {
		encryptedPhone, err := u.encryptPersonalInfoValue(spec.Data.Phone.Encrypted)
		if err != nil {
			return nil, err
		}
		spec.Data.Phone.Encrypted = encryptedPhone
		spec.Data.Phone.Searchable = u.hashPersonalInfoValue(spec.Data.Phone.Searchable)
	}

	if spec.Data.Password != nil {
		hashedPassword, err := u.hashPassword(*spec.Data.Password)
		if err != nil {
			return nil, err
		}
		spec.Data.Password = &hashedPassword
	}

	return u.database.PatchOne(spec)
}

func (u *Users) CompareHashAndPassword(hash, password []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, password)
	return err == nil
}

func (u *Users) hashPersonalInfoValuePtr(value *[]byte) *[]byte {
	if value == nil || len(*value) == 0 {
		return value
	}
	hashedEmail := u.hashPersonalInfoValue(*value)
	return &hashedEmail
}

func (u *Users) hashPersonalInfoValue(value []byte) []byte {
	hashedValue := sha512.Sum512(value)
	return hashedValue[:]
}

func (u *Users) hashPassword(password []byte) ([]byte, error) {
	bytes, err := bcrypt.GenerateFromPassword(password, 8)
	if err != nil {
		return nil, app_error.New(err)
	}
	return bytes, nil
}

func (u *Users) encryptPersonalInfoValue(authMethodPlain []byte) ([]byte, error) {
	encryptedMethodType, err := crypto.AESEncrypt([]byte(u.cfg.Crypto.AuthCredentialsKey), authMethodPlain)
	if err != nil {
		return nil, app_error.New(err)
	}
	return encryptedMethodType, nil
}
