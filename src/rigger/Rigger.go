package rigger

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"reflect"
)

type Rigger struct {
	*gin.Engine                  //把engine放在主类里面
	g           *gin.RouterGroup //路由分组
	//dba interface{}
	props []interface{} //加入的其他对象
}

func Ignite() *Rigger { //所谓的构造函数
	rigger := &Rigger{Engine: gin.New(), props: make([]interface{}, 0)} //返回指针对象需要赋值
	rigger.Use(ErrorHandle())                                           //强制绑定错误处理中间件。不需要修改
	return rigger
}

func (this *Rigger) Start() { //最终启动函数
	config := InitConfig()
	this.Run(fmt.Sprintf(":%d", config.Server.Port)) //这里暂时先写死
}

//中间件方法
func (this *Rigger) Attach(f Fairing) *Rigger {
	this.Use(func(context *gin.Context) {
		err := f.OnRequest(context)
		if err != nil {
			context.AbortWithStatusJSON(400, gin.H{"error": err.Error()}) //中间件发生错误
		} else {
			context.Next() //继续往下走
		}
	})
	return this
}

//beans简单的依赖注入
func (this *Rigger) Beans(dba ...interface{}) *Rigger {
	this.props = append(this.props, dba...) //
	return this
}

//重写Handle方法
func (this *Rigger) Handle(httpMethod, relativePath string, handler interface{}) *Rigger {
	//	if h,ok:=handler.(func(context *gin.Context) string);ok{
	//		this.g.Handle(httpMethod, relativePath, func(context *gin.Context) {
	//			context.String(200,h(context))
	//		})
	//	}
	if h := Convert(handler); h != nil {
		this.g.Handle(httpMethod, relativePath, h)
	}
	return this
}

//不定参数，传递class进来
func (this *Rigger) Mount(group string, classes ...IClass) *Rigger { //返回自己方便链式调用
	this.g = this.Group(group)
	for _, class := range classes {
		class.Build(this)    //这里很关键，这样main里面就不需要调用了
		this.setProps(class) //设置属性
	}
	return this
}

//获取属性
func (this *Rigger) getProps(t reflect.Type) interface{} {
	for _, p := range this.props {
		if t == reflect.TypeOf(p) {
			return p
		}
	}
	return nil
}

//实例化，传进来的对象
func (this *Rigger) setProps(class IClass) {
	vClass := reflect.ValueOf(class).Elem() //通过反射获取class的对象
	for i := 0; i < vClass.NumField(); i++ {
		f := vClass.Field(i)
		if !f.IsNil() || f.Kind() != reflect.Ptr { //如果已经实例化，或者嵌套的对象不是指针类型不做任何处理
			continue
		}
		if p := this.getProps(f.Type()); p != nil { //如果传入的属性与结构体中的属性相等
			f.Set(reflect.New(f.Type().Elem()))     //通过反射设置值
			f.Elem().Set(reflect.ValueOf(p).Elem()) //通过反射初始化对象
		}
	}
}
