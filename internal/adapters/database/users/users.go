package users

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"server/internal/entities"
)

type Users struct {
	creator    *creator
	finderMany *finderMany
	finderOne  *finderOne
	patcherOne *patcherOne
}

func New(database *pgxpool.Pool) *Users {
	return &Users{
		creator:    newCreator(database),
		finderMany: newFinderMany(database),
		finderOne:  newFinderOne(database),
		patcherOne: newPatcherOne(database),
	}
}

func (u *Users) Create(spec entities.UserCreateSpec) (*entities.User, error) {
	return u.creator.Apply(spec)
}

func (u *Users) FindMany(spec entities.UsersFindManySpec) ([]entities.User, error) {
	return u.finderMany.Apply(spec)
}

func (u *Users) FindManyCount(spec entities.UsersFindManySpec) (int64, error) {
	return u.finderMany.ApplyCount(spec)
}

func (u *Users) FindOne(spec entities.UsersFindOneSpec) (*entities.User, error) {
	return u.finderOne.Apply(spec)
}

func (u *Users) PatchOne(spec entities.UsersPatchOneSpec) (*entities.User, error) {
	return u.patcherOne.Apply(spec)
}
