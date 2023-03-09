package domain

import (
	"context"
	"time"

	uuid "github.com/satori/go.uuid"
)

type User struct {
	ID uuid.UUID

	FirstName  string    `json:"first_name"`
	SecondName string    `json:"second_name"`
	Birthdate  time.Time `json:"birthdate"`
	Biography  string    `json:"biography"`
	City       string    `json:"city"`
	Password   string    `json:"password"`
}

type UsersGateway interface {
	GetUserByID(ctx context.Context, id uuid.UUID) (*User, error)
	CreateUser(ctx context.Context, user User) error
}
