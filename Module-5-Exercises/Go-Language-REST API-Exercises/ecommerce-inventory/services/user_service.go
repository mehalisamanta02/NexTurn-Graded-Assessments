package services

import (
	"ecommerce-inventory/models"
	"ecommerce-inventory/repositories"
	"errors"
)

type UserService struct {
	repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

// RegisterUser registers a new user.
func (service *UserService) RegisterUser(user *models.User) error {
	// Validate user data
	if user.Username == "" || user.Password == "" {
		return errors.New("invalid user data")
	}

	// Register the user in the database
	return service.repo.RegisterUser(user)
}

// AuthenticateUser checks if the user's credentials are valid.
func (service *UserService) AuthenticateUser(username, password string) (*models.User, error) {
	user, err := service.repo.GetUserByUsername(username)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	if user.Password != password {
		return nil, errors.New("incorrect password")
	}

	return user, nil
}
