package models

type UserLogin struct {
	Email    string
	Password string
}
type UserOtp struct {
	Email string `json:"email" validate:"required,email" example:"user@example.com"`
}
type UserVerificationRequest struct {
	Email string `json:"email" validate:"required,email"`
}
type UserDetails struct {
	ID        uint
	Firstname string
	Lastname  string
	Email     string
	Phone     string
}
type AdminDetailsResponse struct {
	ID        uint   `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"Email"`
}
type Users struct {
	ID        uint   `json:"id" gorm:"uniquekey; not null"`
	Firstname string `json:"firstname" gorm:"validate:required"`
	Lastname  string `json:"lastname" gorm:"validate:required"`
	Email     string `json:"email" gorm:"validate:required"`
	Password  string `json:"password" gorm:"validate:required"`
	Phone     string `json:"phone" gorm:"validate:required"`
	Blocked   bool   `json:"blocked" gorm:"validate:required"`
	Username  string `json:"user_name"`
	GenderID  uint   `json:"gender_id"`
	Age       uint   `json:"age"`
}
type UserSignup struct {
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=6"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"user_name"`
	GenderID  uint   `json:"gender_id"`
	Age       uint   `json:"age"`
	Number    uint   `json:"number"`
}

type UserUpdateDetails struct {
	Email     string `json:"email" validate:"required,email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"user_name"`
	Password  string `json:"password" validate:"required,min=6"`
}

type UserProfilePhoto struct {
	UserID   uint     `bson:"userId"`
	ImageURL []string `bson:"imageUrl"`
}
type UserDetail struct {
	ID        uint
	Firstname string
	Lastname  string
	Email     string
	Phone     string
	Password  string
}
type TokenUser struct {
	User         UserDetails
	AccessToken  string
	RefreshToken string
}
type UserDetailsResponse struct {
	Id        uint   `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

type Otp struct {
	Email string `json:"email" validate:"required,email"`
	Otp   uint
}
