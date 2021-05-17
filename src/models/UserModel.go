package models

type UserModel struct {
	UserId   int
	UserName string
}

func (this *UserModel) String() string {
	return "userModel"
}
