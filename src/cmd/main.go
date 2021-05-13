package main

import (
	. "gin++/src/classes"
	. "gin++/src/middlewares"
	"gin++/src/rigger"
	"github.com/gin-gonic/gin"
)

func main() {
	//完成这一步
	rigger.Ignite().Attach(NewUserMid()).
		Mount("v1", NewIndexClass(), NewUserClass()).
		Mount("v2", NewUserClass()).
		Start()
	r := gin.New()
	r.Use()
}
