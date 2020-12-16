package interfaces

import (
	"engineering/internal/engineering_service/DTO"
)

type IUserService interface {
	CreateUser(dto DTO.UserCreateDTO)(int64,error)
}
