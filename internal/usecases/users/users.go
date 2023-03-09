package users

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/xsicx/highload/internal/domain"
	"github.com/xsicx/highload/internal/usecases/dto"
	"github.com/xsicx/highload/pkg/password"
)

type Manager struct {
	usersGateway domain.UsersGateway
}

func NewUsersManager(gateway domain.UsersGateway) *Manager {
	return &Manager{usersGateway: gateway}
}

func (m *Manager) Login(ctx context.Context, dto dto.LoginUserDTO) string {
	// TODO: should implement later, for current task - return stub
	b := make([]byte, 12)
	if _, err := rand.Read(b); err != nil {
		return ""
	}

	return hex.EncodeToString(b)
}

func (m *Manager) Register(ctx context.Context, dto dto.RegisterUserDTO) (*domain.User, error) {
	birthdate, err := time.Parse("2006-01-02", dto.Birthdate)
	if err != nil {
		return nil, err
	}

	userPass, err := password.HashPassword(dto.Password)
	if err != nil {
		return nil, err
	}

	user := domain.User{
		ID:         uuid.NewV4(),
		FirstName:  dto.FirstName,
		SecondName: dto.SecondName,
		Birthdate:  birthdate,
		Biography:  dto.Biography,
		City:       dto.City,
		Password:   userPass,
	}

	if err := m.usersGateway.CreateUser(ctx, user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (m *Manager) GetByID(ctx context.Context, userID uuid.UUID) (*domain.User, error) {
	user, err := m.usersGateway.GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, nil
	}

	return user, nil
}
