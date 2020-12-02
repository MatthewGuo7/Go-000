package repos

import "github.com/error_project/models"

type IUserRepo interface {
	GetUserByID(id int)(*models.UserModel, error)
}
