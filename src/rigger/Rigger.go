package rigger

import "github.com/gin-gonic/gin"

type Rigger struct {
	*gin.Engine //把engine放在主类里面
}

func Ignite() *Rigger { //所谓的构造函数
	return &Rigger{Engine: gin.New()}
}

func (this *Rigger) Start() { //最终启动函数
	this.Run(":8080") //这里暂时先写死
}

//不定参数，传递class进来
func (this *Rigger) Mount(classes ...IClass) *Rigger { //返回自己方便链式调用
	for _, class := range classes {
		class.Build(this) //这里很关键，这样main里面就不需要调用了
	}
	return this
}
