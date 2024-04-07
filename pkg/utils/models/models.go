package models

type UserLogin struct {
	Email    string
	Password string
}

type UserDetails struct {
	ID        uint
	Firstname string
	Lastname  string
	Email     string
	Phone     string
}

type UserSignup struct {
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=6"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"user_name"`
	GenderID  int    `json:"gender_id"`
	Age       int    `json:"age`
	Number    string `json:"number"`
}
type UserDetail struct {
	ID        uint
	Firstname string
	Lastname  string
	Email     string
	Phone     string
	Password  string
}
