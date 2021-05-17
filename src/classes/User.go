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
	return &models.UserModel{UserId: 1, UserName: "yuzhonghua"}
}

func (this *UserClass) Build(rigger *rigger.Rigger) {
	rigger.Handle("GET", "/user1", this.UserList)
	rigger.Handle("GET", "/user2", this.UserDetail)
}
