package user

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrEmailExists  = errors.New("email already exists")
	ErrUserNotFound = errors.New("user not found")
)

type Service interface {
	GetAllUsers() ([]User, error)
	CreateUser(email, password string) (*User, error)
	UpdateUser(id string, email, password *string) (*User, error)
	DeleteUser(id string) error
	GetUserByID(id string) (*User, error)
}

type service struct {
	repo UserRepository
}

func NewService(repo UserRepository) Service {
	return &service{repo: repo}
}

func (s *service) CreateUser(email, password string) (*User, error) {
	user := &User{
		ID:       uuid.NewString(),
		Email:    email,
		Password: password,
	}
	return s.repo.Create(user)
}

func (s *service) DeleteUser(id string) error {
	return s.repo.Delete(id)
}

func (s *service) GetAllUsers() ([]User, error) {
	return s.repo.GetAll()
}

func (s *service) GetUserByID(id string) (*User, error) {
	return s.repo.GetByID(id)
}

func (s *service) UpdateUser(id string, email *string, password *string) (*User, error) {
	user, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if email != nil {
		user.Email = *email
	}
	if password != nil {
		user.Password = *password
	}
	return s.repo.Update(user)
}
