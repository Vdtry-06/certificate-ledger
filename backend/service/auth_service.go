package service

import (
	"fmt"
	"time"

	"certificate-ledger/domain"
	"certificate-ledger/repository"
)

type AuthService struct {
	repo *repository.UserRepository
}

func NewAuthService(repo *repository.UserRepository) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (s *AuthService) Register(req domain.UserRequest) (*domain.User, error) {
	_, err := s.repo.FindByEmail(req.Email)
	if err == nil {
		return nil, fmt.Errorf("user with email %s already exists", req.Email)
	}

	user := &domain.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password, 
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := s.repo.Save(user); err != nil {
		return nil, fmt.Errorf("failed to save user: %v", err)
	}

	return user, nil
}

func (s *AuthService) Login(req domain.LoginRequest) (*domain.AuthResponse, error) {

	user, err := s.repo.FindByEmail(req.Email)
	if err != nil {
		return nil, fmt.Errorf("invalid email or password")
	}

	if user.Password != req.Password {
		return nil, fmt.Errorf("invalid email or password")
	}

	token := "dummy-token-" + user.ID

	return &domain.AuthResponse{
		Token: token,
		User:  *user,
	}, nil
}
