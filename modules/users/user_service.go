package users

import (
	"errors"
	"financial-tracker-api/utils"
)

type UserService struct {
	repo *UserRepository
}

func NewUserService(repo *UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Register(req RegisterRequest) error {
	// Check if user already exists
	_, err := s.repo.GetUserByEmail(req.Email)
	if err == nil {
		return errors.New("user already exists")
	}

	// Hash the password
	hashedPassword := utils.HashPassword(req.Password)

	// Create the user
	user := User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
	}

	return s.repo.CreateUser(user)
}

func (s *UserService) Login(req LoginRequest) (User, error) {
	user, err := s.repo.GetUserByEmail(req.Email)
	if err != nil {
		return User{}, errors.New("invalid email or password")
	}

	if !utils.CheckPasswordHash(req.Password, user.Password) {
		return User{}, errors.New("invalid email or password")
	}

	return user, nil
}