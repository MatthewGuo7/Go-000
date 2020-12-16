package DTO

type UserCreateDTO struct {
	UserName string `json:"user_name" validate:"required,min=1"`
}
