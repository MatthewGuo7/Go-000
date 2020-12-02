package models

type UserModel struct {
	UserID   int    `gorm:"column:user_id;primary;auto_increment"`
	UserName string `gorm:"column:user_name"`
	Age      int    `gorm:"column:age"`
}

func NewUserModel() *UserModel {
	return &UserModel{}
}

type UserAttrFunc func(model *UserModel)
type UserAttrFuncs []UserAttrFunc

func (r UserAttrFuncs) Apply(model *UserModel) {
	for _, f := range r {
		f(model)
	}
}

func WithUserID(id int) UserAttrFunc {
	return func(model *UserModel) {
		model.UserID = id
	}
}

func WithUserName(name string) UserAttrFunc {
	return func(model *UserModel) {
		model.UserName = name
	}
}
