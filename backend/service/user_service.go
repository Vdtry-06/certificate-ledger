package service

import (
	"fmt"

	"certificate-ledger/domain"
	"certificate-ledger/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) CreateUser(req domain.UserRequest) (*domain.User, error) {
	_, err := s.repo.FindByEmail(req.Email)
	if err == nil {
		return nil, fmt.Errorf("user with email %s already exists", req.Email)
	}

	user := &domain.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		Role:     req.Role,
	}

	if err := s.repo.Save(user); err != nil {
		return nil, fmt.Errorf("failed to save user: %v", err)
	}

	return user, nil
}

func (s *UserService) GetUser(id string) (*domain.User, error) {
	return s.repo.FindByID(id)
}

func (s *UserService) GetUserByEmail(email string) (*domain.User, error) {
	return s.repo.FindByEmail(email)
}

func (s *UserService) GetAllUsers() ([]*domain.User, error) {
	return s.repo.FindAll()
}

func (s *UserService) DeleteUser(id string) error {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return fmt.Errorf("user with ID %s not found", id)
	}

	if user.Role == "admin" {
		return fmt.Errorf("cannot delete admin user")
	}

	return s.repo.Delete(id)
}