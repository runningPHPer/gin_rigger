package classes

import (
	"fmt"
	"gin_rigger/src/models"
	"gin_rigger/src/rigger"
	"github.com/gin-gonic/gin"
)

type UserClass struct {
	//*rigger.GormAdapter
	*rigger.XormAdapter
	Age *rigger.Value `prefix:"user.age"`
}

func NewUserClass() *UserClass {
	return &UserClass{}
}

func (this *UserClass) UserList(ctx *gin.Context) string {
	return "用户列表" + this.Age.String()
}

func (this *UserClass) UserDetail(ctx *gin.Context) rigger.Model {
	userModel := models.NewUserModel()
	err := ctx.BindUri(userModel)
	rigger.Error(err, "用户ID不合法")
	has, err := this.Table("users").Where("user_id = ?", userModel.UserId).Get(userModel)
	if !has {
		rigger.Error(fmt.Errorf("没有该用户"))
	}
	rigger.Error(err)
	return userModel
}

func (this *UserClass) Users(ctx *gin.Context) rigger.Models {
	users := []*models.UserModel{
		{101, "yuzhonghua"},
		{102, "pengjun"},
	}
	return rigger.MakeModels(users)
}
func (this *UserClass) Build(rigger *rigger.Rigger) {
	rigger.Handle("GET", "/user1", this.UserList)
	rigger.Handle("GET", "/user2/:id", this.UserDetail)
	rigger.Handle("GET", "/user3", this.Users)
}
