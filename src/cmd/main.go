package main

import (
	. "gin++/src/classes"
	"gin++/src/rigger"
)

func main() {
	//完成这一步
	rigger.Ignite().
		Mount(NewIndexClass(),
			NewUserClass()).
		Start()
}
