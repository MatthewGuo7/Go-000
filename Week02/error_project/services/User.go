package services

import (
	"github.com/error_project/models"
	"github.com/error_project/repos"
	"github.com/pkg/errors"
)

var ErrNoPermission = errors.New("No Permission")

type User struct {
	UserModel *models.UserModel
	repos     repos.IUserRepo
}

func NewUser(userModel *models.UserModel, repos repos.IUserRepo) *User {
	return &User{UserModel: userModel, repos: repos}
}

func (r *User) GetUserByID(id int) (*models.UserModel, error) {
	user, err := r.repos.GetUserByID(id)
	if nil != err {
		return nil, err
	}

	if err := r.IsNoPermission(user.Age); nil != err {
		return nil, errors.WithStack(err)
	}

	return user, nil
}

//age < 10 --> no permission
func (r *User) IsNoPermission(age int) error {
	if age < 10 {
		return ErrNoPermission
	}

	return nil
}
