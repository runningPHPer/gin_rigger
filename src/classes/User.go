package classes

import (
	"gin++/src/rigger"
	"github.com/gin-gonic/gin"
)

type UserClass struct {
}

func NewUserClass() *UserClass {
	return &UserClass{}
}

func (this *UserClass) UserList() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(200, gin.H{
			"result": "success",
		})
	}
}

func (this *UserClass) Build(rigger *rigger.Rigger) {
	rigger.Handle("GET", "/users", this.UserList())
}
