package database

import (
	"context"

	"github.com/jmoiron/sqlx"
	uuid "github.com/satori/go.uuid"
	"github.com/xsicx/highload/internal/domain"
)

func NewUsersGateway(db *sqlx.DB) *UsersGateway {
	return &UsersGateway{
		db: db,
	}
}

type UsersGateway struct {
	db *sqlx.DB
}

func (g *UsersGateway) GetUserByID(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	return nil, nil
}

func (g *UsersGateway) CreateUser(ctx context.Context, user domain.User) error {
	return nil
}
