package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
	Login(input LoginInput) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{
		repository,
	}
}

func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	user := User{}
	user.Name = input.Name
	user.Email = input.Email
	user.Occupation = input.Occupation
	passwordhash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}

	user.PasswordHash = string(passwordhash)
	user.Role = "user"

	userData, err := s.repository.Save(user)
	if err != nil {
		return user, err
	}

	return userData, nil
}

func (s *service) Login(input LoginInput) (User, error) {
	email := input.Email
	password := input.Password

	userData, err := s.repository.FindByEmail(email)
	if err != nil {
		return userData, err
	}

	if userData.ID == 0 {
		return userData, errors.New("No data user")
	}

	err = bcrypt.CompareHashAndPassword([]byte(userData.PasswordHash), []byte(password))
	if err != nil {
		return userData, err
	}

	return userData, nil
}
