package main

import (
	. "gin_rigger/src/classes"
	. "gin_rigger/src/middlewares"
	"gin_rigger/src/rigger"
	"log"
)

func main() {
	//工具类加入
	//rigger.GenTplFunc("src/funcs")  //自动生成一个 funcmap.go.要生存的文件路径
	//return
	//完成这一步
	rigger.Ignite().
		Beans(rigger.NewGormAdapter(), rigger.NewXormAdapter()).
		Attach(NewUserMid()).
		Mount("v1", NewIndexClass(), NewCourseClass()).
		Mount("v2", NewUserClass()).
		Task("0/3 * * * * *", func() {
			log.Println("计划任务，执行")
		}).
		Start()
}
