package models

type UserModel struct {
	UserId   int    `json:"user_id" gorm:"column:user_id" uri:"id" binding:"required,gt=0"`
	UserName string `json:"user_name" gorm:"column:user_name"`
}

func NewUserModel() *UserModel {
	return &UserModel{}
}

func (this *UserModel) String() string {
	return "userModel"
}
