package dao

import (
	"github.com/error_project/models"
	"github.com/error_project/repos"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

var _ repos.IUserRepo = &UserImpl{}

type UserImpl struct {
	db *gorm.DB
}

func (u *UserImpl) GetUserByID(id int) (*models.UserModel, error) {
	var user *models.UserModel  //此处需要报错

	err := u.db.Raw("select * from user_models where user_id = ? ", id).Find(&user).Error
	if err == gorm.ErrRecordNotFound {
		err = ErrNotFound
	}
	return user, errors.WithStack(err)
}

func NewUserImpl(db *gorm.DB) *UserImpl {
	return &UserImpl{db: db}
}


