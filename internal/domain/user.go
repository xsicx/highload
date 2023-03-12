package domain

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type User struct {
	ID uuid.UUID `db:"id" json:"id"`

	FirstName  string    `db:"first_name" json:"first_name"`
	SecondName string    `db:"second_name" json:"second_name"`
	Birthdate  time.Time `db:"birthdate" json:"birthdate"`
	Biography  string    `db:"biography" json:"biography"`
	City       string    `db:"city" json:"city"`
	Password   string    `db:"password" json:"password"`
}
