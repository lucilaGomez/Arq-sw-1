package dto

type UserDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Name     string `json:"name" validate:"required"`
	LastName string `json:"lastname" validate:"required"`
	Password string `json:"password" validate:"required"`
	Admin    bool   `json:"admin"`
}

type UsersDTO []UserDTO
