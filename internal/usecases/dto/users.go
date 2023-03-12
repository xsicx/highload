package dto

type LoginUserRequest struct {
	ID       string `json:"id"`
	Password string `json:"password"`
}

type RegisterUserRequest struct {
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
	Birthdate  string `json:"birthdate"`
	Biography  string `json:"biography"`
	City       string `json:"city"`
	Password   string `json:"password"`
}

type UserResponse struct {
	ID         string `json:"id"`
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
	Age        int    `json:"age"`
	Birthdate  string `json:"birthdate"`
	Biography  string `json:"biography"`
	City       string `json:"city"`
}

type UserTokenResponse struct {
	Token string `json:"token"`
}

type UserIDResponse struct {
	UserID string `json:"user_id"`
}
