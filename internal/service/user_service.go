package service

import (
	"time"

	db "go-user-api/db/sqlc/generated"
	"go-user-api/internal/repository"
)

type UserService struct {
	Repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		Repo: repo,
	}
}

func (s *UserService) CreateUser(
	name string,
	dob string,
) (db.User, error) {

	return s.Repo.CreateUser(name, dob)
}

func (s *UserService) GetAllUsers() ([]db.User, error) {
	return s.Repo.GetAllUsers()
}

func (s *UserService) GetUserByID(id int32) (db.User, error) {
	return s.Repo.GetUserByID(id)
}

func (s *UserService) DeleteUser(
	id int32,
) error {

	return s.Repo.DeleteUser(id)
}

func (s *UserService) UpdateUser(
	id int32,
	name string,
	dob string,
) (db.User, error) {

	return s.Repo.UpdateUser(
		id,
		name,
		dob,
	)
}

func CalculateAge(dob time.Time) int {
	now := time.Now()

	age := now.Year() - dob.Year()

	if now.YearDay() < dob.YearDay() {
		age--
	}

	return age
}