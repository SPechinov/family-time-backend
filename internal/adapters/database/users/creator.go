package users

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"server/internal/entities"
	"server/pkg/app_error"
)

type queryCreateParams struct {
	Password        []byte
	FirstName       string
	LastName        *string
	EmailEncrypted  *[]byte
	EmailSearchable *[]byte
	PhoneEncrypted  *[]byte
	PhoneSearchable *[]byte
	CountryCode     string
}

const queryCreate = `
	INSERT INTO main.public.users (
	   	password,
	   	first_name,
	   	last_name,
	   	email_encrypted,
	   	email_searchable,
	   	phone_encrypted,
	   	phone_searchable,
		country_code
	) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
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

type creator struct {
	database *pgxpool.Pool
}

func newCreator(database *pgxpool.Pool) *creator {
	return &creator{
		database: database,
	}
}

func (c *creator) Apply(entity entities.UserCreateSpec) (*entities.User, error) {
	queryParams := c.parseEntityToQueryParams(entity)

	row := c.query(context.TODO(), queryParams)

	var newUser userData
	err := c.scan(row, &newUser)

	if err != nil {
		return nil, err
	}

	userEntity := mapUserDataToUserEntity(&newUser)
	return &userEntity, nil
}

func (c *creator) scan(row pgx.Row, newUser *userData) error {
	err := row.Scan(
		&newUser.UserID,
		&newUser.EmailEncrypted,
		&newUser.PhoneEncrypted,
		&newUser.EmailSearchable,
		&newUser.PhoneSearchable,
		&newUser.FirstName,
		&newUser.LastName,
		&newUser.CreatedAt,
		&newUser.UpdatedAt,
		&newUser.DeletedAt,
	)
	if err != nil {
		return app_error.New(err)
	}
	return err
}

func (c *creator) query(ctx context.Context, params *queryCreateParams) pgx.Row {
	return c.database.QueryRow(
		ctx,
		queryCreate,
		params.Password,
		params.FirstName,
		params.LastName,
		params.EmailEncrypted,
		params.EmailSearchable,
		params.PhoneEncrypted,
		params.PhoneSearchable,
		params.CountryCode,
	)
}

func (c *creator) parseEntityToQueryParams(entity entities.UserCreateSpec) *queryCreateParams {
	params := queryCreateParams{
		Password:    entity.Password,
		FirstName:   entity.FirstName,
		CountryCode: entity.CountryCode,
	}

	switch entity.AuthMethodSpec.Type {
	case entities.AuthMethodEmail:
		params.EmailEncrypted = &entity.AuthMethodSpec.Values.Encrypted
		params.EmailSearchable = &entity.AuthMethodSpec.Values.Searchable
	case entities.AuthMethodPhone:
		params.PhoneEncrypted = &entity.AuthMethodSpec.Values.Encrypted
		params.PhoneSearchable = &entity.AuthMethodSpec.Values.Searchable
	}

	return &params
}
