package dto

type LoginUserDTO struct {
	ID       string
	Password string
}

type RegisterUserDTO struct {
	FirstName  string
	SecondName string
	Birthdate  string
	Biography  string
	City       string
	Password   string
}
