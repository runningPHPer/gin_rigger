package classes

import (
	"gin++/src/models"
	"gin++/src/rigger"
	"github.com/gin-gonic/gin"
)

type UserClass struct {
}

func NewUserClass() *UserClass {
	return &UserClass{}
}

func (this *UserClass) UserList(ctx *gin.Context) string {
	return "用户列表"
}

func (this *UserClass) UserDetail(ctx *gin.Context) rigger.Model {
	userModel := models.NewUserModel()
	err := ctx.BindUri(userModel)
	rigger.Error(err, "用户ID不合法")
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
