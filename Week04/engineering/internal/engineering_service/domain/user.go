package domain

type User struct {

}

func NewUser() *User {
	return &User{}
}

func (r *User) Create() (int64, error) {
	return 1, nil
}
