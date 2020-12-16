package web

import (
	"engineering/apis/interfaces"
	"engineering/internal/engineering_service/DTO"
	"engineering/pkg/http_server"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserAPI struct {
	userService interfaces.IUserService
}

func NewUserAPI(userService interfaces.IUserService) *UserAPI {
	return &UserAPI{userService: userService}
}

func (r *UserAPI) Create(ctx *gin.Context) {

	createData := DTO.UserCreateDTO{}
	_, _ = r.userService.CreateUser(createData)

	ctx.JSON(http.StatusOK,gin.H{"code":http.StatusOK, "msg":http.StatusText(http.StatusOK)})
}

func (u *UserAPI) AddRouter(server *http_server.HttpServer) {
	server.Handle("POST","/user/create", u.Create)
}


