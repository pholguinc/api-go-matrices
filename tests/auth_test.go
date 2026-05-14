package tests

import (
	"testing"
	"github.com/pholguinc/api-go-matrices/internal/dtos"
	"github.com/pholguinc/api-go-matrices/internal/models"
	"github.com/pholguinc/api-go-matrices/internal/services"
)

type MockUserRepository struct {
	users map[string]*models.User
}

func (m *MockUserRepository) Create(user *models.User) error {
	m.users[user.Email] = user
	return nil
}

func (m *MockUserRepository) FindByEmail(email string) (*models.User, error) {
	if user, ok := m.users[email]; ok {
		return user, nil
	}
	return nil, nil
}

func (m *MockUserRepository) FindByID(id string) (*models.User, error) { return nil, nil }

func TestRegister(t *testing.T) {
	mockRepo := &MockUserRepository{users: make(map[string]*models.User)}
	service := services.NewAuthService(mockRepo)

	req := dtos.RegisterRequest{Email: "test@example.com", Password: "password123"}
	user, err := service.Register(req)

	if err != nil || user.Email != req.Email {
		t.Fatalf("Registro fallido")
	}
}

func TestLogin(t *testing.T) {
	mockRepo := &MockUserRepository{users: make(map[string]*models.User)}
	service := services.NewAuthService(mockRepo)

	email, pass := "login@test.com", "secret123"
	_, _ = service.Register(dtos.RegisterRequest{Email: email, Password: pass})

	token, _, err := service.Login(dtos.LoginRequest{Email: email, Password: pass})

	if err != nil || token == "" {
		t.Errorf("Login fallido: %v", err)
	}
}
