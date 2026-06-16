package repository

import (
	"context"
	"time"

	db "go-user-api/db/sqlc/generated"
)

type UserRepository struct {
	Queries *db.Queries
}

func NewUserRepository(q *db.Queries) *UserRepository {
	return &UserRepository{
		Queries: q,
	}
}

func (r *UserRepository) CreateUser(
	name string,
	dob string,
) (db.User, error) {

	parsedDob, err := time.Parse("2006-01-02", dob)
	if err != nil {
		return db.User{}, err
	}

	return r.Queries.CreateUser(
		context.Background(),
		db.CreateUserParams{
			Name: name,
			Dob:  parsedDob,
		},
	)
}

func (r *UserRepository) GetAllUsers(
	limit int32,
	offset int32,
) ([]db.User, error) {

	return r.Queries.ListUsers(
		context.Background(),
		db.ListUsersParams{
			Limit:  limit,
			Offset: offset,
		},
	)
}

func (r *UserRepository) GetUserByID(id int32) (db.User, error) {
	return r.Queries.GetUser(
		context.Background(),
		id,
	)
}

func (r *UserRepository) DeleteUser(
	id int32,
) error {

	return r.Queries.DeleteUser(
		context.Background(),
		id,
	)
}

func (r *UserRepository) UpdateUser(
	id int32,
	name string,
	dob string,
) (db.User, error) {

	parsedDob, err := time.Parse(
		"2006-01-02",
		dob,
	)

	if err != nil {
		return db.User{}, err
	}

	return r.Queries.UpdateUser(
		context.Background(),
		db.UpdateUserParams{
			ID:   id,
			Name: name,
			Dob:  parsedDob,
		},
	)
}