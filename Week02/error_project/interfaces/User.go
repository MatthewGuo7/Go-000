package interfaces

import (
	"fmt"
	"github.com/error_project/infra/dao"
	"github.com/error_project/infra/db"
	"github.com/error_project/models"
	"github.com/error_project/repos"
	"github.com/error_project/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUserByID() gin.HandlerFunc {
	return func(context *gin.Context) {

		var (
			err error
			retUser *models.UserModel
		)

		defer func() {
			if err != nil {
				fmt.Printf("error = %+v", err)
			}
		}()

		id := struct {
			ID int `uri:"id"`
		}{}

		err = context.ShouldBind(id)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "msg": http.StatusText(http.StatusBadRequest)})
			return
		}

		var userRepo repos.IUserRepo = dao.NewUserImpl(db.GormDB)

		userService := services.NewUser(nil, userRepo)
		retUser, err = userService.GetUserByID(id.ID)
		if err != nil {
			if dao.IsNotFound(err) {
				err = nil
			}
		}

		context.JSON(http.StatusOK, gin.H{"code": 200, "msg": "success", "data": retUser})
	}
}
