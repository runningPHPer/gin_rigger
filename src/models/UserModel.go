package models

type UserModel struct {
	UserId   int `uri:"id" binding:"require,gt>0"`
	UserName string
}

func NewUserModel() *UserModel {
	return &UserModel{}
}

func (this *UserModel) String() string {
	return "userModel"
}
