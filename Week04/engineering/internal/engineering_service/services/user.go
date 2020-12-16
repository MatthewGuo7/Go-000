package services

import (
	"engineering/apis/interfaces"
	"engineering/internal/engineering_service/DTO"
	"engineering/internal/engineering_service/domain"
)

var _ interfaces.IUserService = &UserServiceImpl{}

type UserServiceImpl struct {
	user *domain.User
	userLog *domain.UserLog
}

func NewUserServiceImpl(user *domain.User, userLog *domain.UserLog) *UserServiceImpl {
	return &UserServiceImpl{user: user, userLog: userLog}
}

func (u *UserServiceImpl) CreateUser(dto DTO.UserCreateDTO) (int64, error) {
	return 1, nil
}
