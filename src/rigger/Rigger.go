package rigger

import "github.com/gin-gonic/gin"

type Rigger struct {
	*gin.Engine                  //把engine放在主类里面
	g           *gin.RouterGroup //路由分组
}

func Ignite() *Rigger { //所谓的构造函数
	return &Rigger{Engine: gin.New()}
}

func (this *Rigger) Start() { //最终启动函数
	this.Run(":8080") //这里暂时先写死
}

//中间件方法
func (this *Rigger) Attach(f Fairing) *Rigger {
	this.Use(func(context *gin.Context) {
		err := f.OnRequest()
		if err != nil {
			context.AbortWithStatusJSON(400, gin.H{"error": err.Error()}) //中间件发生错误
		} else {
			context.Next() //继续往下走
		}
	})
	return this
}

//重写Handle方法
func (this *Rigger) Handle(httpMethod, relativePath string, handlers ...gin.HandlerFunc) *Rigger {
	this.g.Handle(httpMethod, relativePath, handlers...)
	return this
}

//不定参数，传递class进来
func (this *Rigger) Mount(group string, classes ...IClass) *Rigger { //返回自己方便链式调用
	this.g = this.Group(group)
	for _, class := range classes {
		class.Build(this) //这里很关键，这样main里面就不需要调用了
	}
	return this
}
