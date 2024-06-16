package models

type User struct {
	ID           uint   `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	PasswordHash string `json:"-"`
}

type LoginRequest struct {
	Email        string `json:"email" validate:"required,email"`
	PasswordHash string `json:"password" validate:"required"`
}

type LoginResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

type RegisterRequest struct {
	Name         string `json:"name" validate:"required,min=3,max=20"`
	Email        string `json:"email" validate:"required,email"`
	PasswordHash string `json:"password" validate:"required,min=6,max=30"`
}

type UpdateUserRequest struct {
	Name string `json:"name" validate:"required,min=3,max=20"`
}

type RefreshResponse struct {
	RefreshToken string `json:"refresh_token"`
}
