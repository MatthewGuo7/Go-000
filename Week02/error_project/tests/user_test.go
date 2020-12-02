package tests

import (
	"fmt"
	"github.com/error_project/infra/dao"
	"github.com/error_project/infra/db"
	"github.com/error_project/repos"
	"github.com/error_project/services"
	"testing"
)

func init()  {
	err := db.InitMysql()
	if err != nil {
		fmt.Printf("err = %+v\n",err)
		panic(err)
	}
}

func TestGetUser(t *testing.T) {
	var userRepo repos.IUserRepo = dao.NewUserImpl(db.GormDB)

	userService := services.NewUser(nil, userRepo)
	retUser, err := userService.GetUserByID(1)
	fmt.Println(retUser, err)
}
