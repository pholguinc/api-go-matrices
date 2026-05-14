package services

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/pholguinc/api-go-matrices/internal/dtos"
	"github.com/pholguinc/api-go-matrices/internal/models"
	"github.com/pholguinc/api-go-matrices/internal/repositories"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(req dtos.RegisterRequest) (*models.User, error)
	Login(req dtos.LoginRequest) (string, *models.User, error)
}

type authService struct {
	repo repositories.UserRepository
}

func NewAuthService(repo repositories.UserRepository) AuthService {
	return &authService{repo: repo}
}

func (s *authService) Register(req dtos.RegisterRequest) (*models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := models.User{
		Email:    req.Email,
		Password: string(hashedPassword),
	}

	if err := s.repo.Create(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *authService) Login(req dtos.LoginRequest) (string, *models.User, error) {
	user, err := s.repo.FindByEmail(req.Email)
	if err != nil {
		return "", nil, errors.New("credenciales inválidas")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return "", nil, errors.New("credenciales inválidas")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", nil, err
	}

	return tokenString, user, nil
}
