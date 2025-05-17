package service

import (
	"fmt"
	"os"
	"time"

	"certificate-ledger/domain"
	"certificate-ledger/repository"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
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

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %v", err)
	}

	role := req.Role
	if role == "" {
		role = "user"
	}

	user := &domain.User{
		Name:      req.Name,
		Email:     req.Email,
		Password:  string(hashedPassword),
		Role:      role,
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

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, fmt.Errorf("invalid email or password")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"role":    user.Role,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %v", err)
	}

	return &domain.AuthResponse{
		Token: tokenString,
		User:  *user,
	}, nil
}