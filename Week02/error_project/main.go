package main

import (
	"fmt"
	"github.com/error_project/infra/db"
	"github.com/error_project/interfaces"
	"github.com/error_project/middlewares"
	"github.com/gin-gonic/gin"
)

func main() {

	err := db.InitMysql()
	if err != nil {
		fmt.Printf("err = %+v\n",err)
		panic(err)
	}

	r := gin.New()
	r.Use(middlewares.PanicHandler())
	r.Handle("GET","/user/info/:id", interfaces.GetUserByID())

	if err := r.Run(); err != nil {
		panic(err)
	}
}
