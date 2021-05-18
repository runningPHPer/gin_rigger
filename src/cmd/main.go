package main

import (
	"fmt"
	. "gin_rigger/src/classes"
	. "gin_rigger/src/middlewares"
	"gin_rigger/src/rigger"
)

func main() {
	//完成这一步
	fmt.Println(rigger.InitConfig().Server)
	rigger.Ignite().
		Beans(rigger.NewGormAdapter(), rigger.NewXormAdapter()).
		Attach(NewUserMid()).
		Mount("v1", NewIndexClass(), NewUserClass()).
		Mount("v2", NewUserClass()).
		Start()
}
