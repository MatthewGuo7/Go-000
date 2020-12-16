//+build wireinject

package main

import (
	"engineering/internal/engineering_service/domain"
	"engineering/internal/engineering_service/services"
	"github.com/google/wire"
)

//经测试，构造函数里传递的接口 没有办法被创建，因为没有 接口的new方法???
func InitUserAPI() *services.UserServiceImpl {
	wire.Build(services.NewUserServiceImpl, domain.NewUser, domain.NewUserLog)
	return &services.UserServiceImpl{}
}
