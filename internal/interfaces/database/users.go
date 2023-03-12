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
	user := domain.User{}

	query := `SELECT id, first_name, second_name, birthdate, biography, city, password FROM users WHERE id = $1`

	err := g.db.GetContext(ctx, &user, query, id)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (g *UsersGateway) CreateUser(ctx context.Context, user domain.User) error {
	query := `INSERT INTO users (id, first_name, second_name, birthdate, biography, city, password) 
			VALUES ($1, $2, $3, $4, $5, $6, $7)`

	_, err := g.db.ExecContext(
		ctx,
		query,
		user.ID, user.FirstName, user.SecondName, user.Birthdate, user.Biography, user.City, user.Password,
	)
	if err != nil {
		return err
	}

	return nil
}
