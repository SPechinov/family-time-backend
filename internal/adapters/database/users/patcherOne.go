package users

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"server/internal/entities"
	"server/pkg/app_error"
	"strings"
	"time"
)

type patcherOne struct {
	database *pgxpool.Pool
}

const queryPatchOneReturning = `
	RETURNING
	    id,
	    email_encrypted,
	    phone_encrypted,
	    email_searchable,
	    phone_searchable,
	    first_name,
	    last_name,
	    created_at,
	    updated_at,
	    deleted_at
`

func newPatcherOne(database *pgxpool.Pool) *patcherOne {
	return &patcherOne{
		database: database,
	}
}

func (p *patcherOne) Apply(spec entities.UsersPatchOneSpec) (*entities.User, error) {
	query, args, err := p.buildQuery(spec)
	if err != nil {
		return nil, err
	}

	row := p.database.QueryRow(context.TODO(), query, args...)

	var userDB userData
	err = p.scan(row, &userDB)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	userEntity := mapUserDataToUserEntity(&userDB)
	return &userEntity, nil
}

func (p *patcherOne) scan(row pgx.Row, userDB *userData) error {
	err := row.Scan(
		&userDB.UserID,
		&userDB.EmailEncrypted,
		&userDB.PhoneEncrypted,
		&userDB.EmailSearchable,
		&userDB.PhoneSearchable,
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

func (p *patcherOne) buildQuery(spec entities.UsersPatchOneSpec) (string, []any, error) {
	if err := p.isValidSpec(spec); err != nil {
		return "", nil, err
	}

	query := `UPDATE main.public.users`

	var args []any
	argsPosition := 1

	var sets []string

	if spec.Data.Email != nil {
		args = append(args, spec.Data.Email.Searchable)
		sets = append(sets, fmt.Sprintf("email_searchable = $%d", argsPosition))
		argsPosition++

		args = append(args, spec.Data.Email.Encrypted)
		sets = append(sets, fmt.Sprintf("email_encrypted = $%d", argsPosition))
		argsPosition++
	}

	if spec.Data.Phone != nil {
		args = append(args, spec.Data.Phone.Searchable)
		sets = append(sets, fmt.Sprintf("phone_searchable = $%d", argsPosition))
		argsPosition++

		args = append(args, spec.Data.Phone.Encrypted)
		sets = append(sets, fmt.Sprintf("phone_encrypted = $%d", argsPosition))
		argsPosition++
	}

	if spec.Data.Password != nil {
		args = append(args, *spec.Data.Password)
		sets = append(sets, fmt.Sprintf("password = $%d", argsPosition))
		argsPosition++
	}

	if spec.Data.FirstName != nil {
		args = append(args, *spec.Data.FirstName)
		sets = append(sets, fmt.Sprintf("first_name = $%d", argsPosition))
		argsPosition++
	}

	if spec.Data.LastName != nil {
		args = append(args, *spec.Data.LastName)
		sets = append(sets, fmt.Sprintf("last_name = $%d", argsPosition))
		argsPosition++
	}

	if spec.Data.Deleted != nil {
		sets = append(sets, fmt.Sprintf("deleted = $%d", argsPosition))
		if *spec.Data.Deleted {
			args = append(args, time.Now())
		} else {
			args = append(args, nil)
		}
		argsPosition++
	}

	query += " SET " + strings.Join(sets, ", ")

	var wheres []string
	if spec.UserID != nil {
		args = append(args, *spec.UserID)
		wheres = append(wheres, fmt.Sprintf("user_id = $%d", argsPosition))
		argsPosition++
	}

	if spec.Email != nil {
		args = append(args, *spec.Email)
		wheres = append(wheres, fmt.Sprintf("email_searchable = $%d", argsPosition))
		argsPosition++
	}

	if spec.Phone != nil {
		args = append(args, *spec.Phone)
		wheres = append(wheres, fmt.Sprintf("phone_searchable = $%d", argsPosition))
		argsPosition++
	}

	query += " WHERE " + strings.Join(wheres, " AND ")
	query += " " + queryPatchOneReturning

	return query, args, nil
}

func (p *patcherOne) isValidSpec(spec entities.UsersPatchOneSpec) error {
	if spec.UserID != nil || spec.Email != nil || spec.Phone != nil {
		return nil
	}
	return app_error.New(errors.New("UserID or Email or Phone must be provided"))
}
