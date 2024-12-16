package users

import (
	"server/internal/entities"
)

type database interface {
	Create(spec entities.UserCreateSpec) (user *entities.User, err error)
	FindMany(spec entities.UsersFindManySpec) (users []entities.User, err error)
	FindOne(spec entities.UsersFindOneSpec) (user *entities.User, err error)
	PatchOne(spec entities.UsersPatchOneSpec) (*entities.User, error)
}
