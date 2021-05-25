package main

import (
	. "gin_rigger/src/classes"
	. "gin_rigger/src/middlewares"
	"gin_rigger/src/rigger"
)

func main() {
	//完成这一步
	rigger.Ignite().
		Beans(rigger.NewGormAdapter(), rigger.NewXormAdapter()).
		Attach(NewUserMid()).
		Mount("v1", NewIndexClass()).
		Mount("v2", NewUserClass()).
		Start()
}
