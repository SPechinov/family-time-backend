package users

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"server/internal/entities"
	"server/pkg/app_error"
)

type finderOne struct {
	database *pgxpool.Pool
}

const queryFindOne = `
	SELECT
	    id,
	    email_encrypted,
	    phone_encrypted,
	    email_searchable,
	    phone_searchable,
		password,
	    first_name,
	    last_name,
	    created_at,
	    updated_at,
	    deleted_at
	FROM main.public.users
`

func newFinderOne(database *pgxpool.Pool) *finderOne {
	return &finderOne{
		database: database,
	}
}

func (fo *finderOne) Apply(spec entities.UsersFindOneSpec) (*entities.User, error) {
	query, args := fo.buildQuery(queryFindOne, spec)

	row := fo.database.QueryRow(context.TODO(), query, args...)

	var userDB userData
	err := fo.scan(row, &userDB)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	userEntity := mapUserDataToUserEntity(&userDB)
	return &userEntity, nil
}

func (fo *finderOne) scan(row pgx.Row, userDB *userData) error {
	err := row.Scan(
		&userDB.UserID,
		&userDB.EmailEncrypted,
		&userDB.PhoneEncrypted,
		&userDB.EmailSearchable,
		&userDB.PhoneSearchable,
		&userDB.Password,
		&userDB.FirstName,
		&userDB.LastName,
		&userDB.CreatedAt,
		&userDB.UpdatedAt,
		&userDB.DeletedAt,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		return err
	}

	if err != nil {
		return app_error.New(err)
	}
	return err
}

func (fo *finderOne) buildQuery(clearQuery string, spec entities.UsersFindOneSpec) (string, []any) {
	var args []any
	argsPosition := 1

	query := clearQuery + " WHERE 1=1"

	if spec.UserID != nil {
		query += fmt.Sprintf(" AND id = $%d", argsPosition)
		args = append(args, *spec.UserID)
		argsPosition++
	}

	if spec.Email != nil {
		query += fmt.Sprintf(" AND email_searchable = $%d", argsPosition)
		args = append(args, *spec.Email)
		argsPosition++
	}

	if spec.Phone != nil {
		query += fmt.Sprintf(" AND phone_searchable = $%d", argsPosition)
		args = append(args, *spec.Phone)
		argsPosition++
	}

	if spec.Deleted {
		query += fmt.Sprintf(" AND deleted_at IS NOT NULL")
	} else {
		query += fmt.Sprintf(" AND deleted_at IS NULL")
	}

	return query, args
}
