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

type finderMany struct {
	database *pgxpool.Pool
}

const queryFindMany = `
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

const queryFindManyCount = `
	SELECT COUNT(*)
	FROM main.public.users
`

func newFinderMany(database *pgxpool.Pool) *finderMany {
	return &finderMany{
		database: database,
	}
}

func (fm *finderMany) Apply(spec entities.UsersFindManySpec) ([]entities.User, error) {
	query, args := fm.buildQuery(queryFindMany, spec)
	rows, err := fm.database.Query(context.TODO(), query, args...)
	if err != nil {
		return nil, app_error.New(err)
	}
	defer rows.Close()

	userEntities, err := fm.scanRows(rows)
	if err != nil {
		return nil, err
	}

	return userEntities, nil
}

func (fm *finderMany) ApplyCount(spec entities.UsersFindManySpec) (int64, error) {
	query, args := fm.buildQuery(queryFindManyCount, spec)
	result, err := fm.database.Exec(context.TODO(), query, args...)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return 0, nil
		}
		return 0, app_error.New(err)
	}

	return result.RowsAffected(), nil
}

func (fm *finderMany) scanRows(rows pgx.Rows) ([]entities.User, error) {
	users := make([]entities.User, 0)
	for rows.Next() {
		var userDB userData
		err := rows.Scan(
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

		if err != nil {
			return nil, app_error.New(err)
		}

		users = append(users, mapUserDataToUserEntity(&userDB))
	}

	return users, nil
}

func (fm *finderMany) buildQuery(clearQuery string, spec entities.UsersFindManySpec) (string, []any) {
	var args []any
	argsPosition := 1
	query := clearQuery + " WHERE 1=1"

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

	if spec.CreatedAtFrom != nil {
		query += fmt.Sprintf(" AND created_at >= $%d", argsPosition)
		args = append(args, *spec.CreatedAtFrom)
		argsPosition++
	}

	if spec.CreatedAtTo != nil {
		query += fmt.Sprintf(" AND created_at <= $%d", argsPosition)
		args = append(args, *spec.CreatedAtTo)
		argsPosition++
	}

	if spec.Deleted {
		query += fmt.Sprintf(" AND deleted_at IS NOT NULL")
	} else {
		query += fmt.Sprintf(" AND deleted_at IS NULL")
	}

	if spec.Pagination != nil {
		query += " ORDER BY created_at DESC"
		query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", spec.Pagination.Limit, spec.Pagination.Offset)
	}

	return query, args
}
