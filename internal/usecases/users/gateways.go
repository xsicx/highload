package users

import (
	"context"

	uuid "github.com/satori/go.uuid"
	"github.com/xsicx/highload/internal/domain"
)

type Gateway interface {
	GetUserByID(ctx context.Context, id uuid.UUID) (*domain.User, error)
	CreateUser(ctx context.Context, user domain.User) error
}
