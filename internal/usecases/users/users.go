package users

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"time"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/xsicx/highload/internal/domain"
	"github.com/xsicx/highload/internal/usecases/dto"
	"github.com/xsicx/highload/pkg/password"
)

type Manager struct {
	usersGateway Gateway
}

func NewUsersManager(gateway Gateway) *Manager {
	return &Manager{usersGateway: gateway}
}

func (m *Manager) Login(ctx context.Context, dto dto.LoginUserRequest) (string, error) {
	userID, err := uuid.FromString(dto.ID)
	if err != nil {
		return "", errors.Wrap(err, "uuid parser error")
	}

	user, err := m.usersGateway.GetUserByID(ctx, userID)
	if err != nil {
		return "", err
	}

	if user == nil {
		return "", errors.New("user not found")
	}

	if !password.CheckPasswordHash(dto.Password, user.Password) {
		return "", errors.New("wrong password")
	}

	// TODO: should implement later, for current task - return stub
	b := make([]byte, 12)
	if _, err := rand.Read(b); err != nil {
		return "", errors.Wrap(err, "generating token error")
	}

	return hex.EncodeToString(b), nil
}

func (m *Manager) Register(ctx context.Context, dto dto.RegisterUserRequest) (*domain.User, error) {
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
