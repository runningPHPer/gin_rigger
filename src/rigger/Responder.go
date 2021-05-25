package rigger

import (
	"github.com/gin-gonic/gin"
	"reflect"
)

var ResponderList []Responder

func init() {
	ResponderList = []Responder{new(StringResponder),
		new(ModelResponder),
		new(ModelsResponder),
		new(ViewResponder)} //返回的是指针
}

//输出格式。返回gin.handlerFunc
type Responder interface {
	RespondTo() gin.HandlerFunc
}

//类型转换函数
func Convert(handler interface{}) gin.HandlerFunc {
	h_ref := reflect.ValueOf(handler)
	for _, r := range ResponderList {
		r_ref := reflect.ValueOf(r).Elem()            //reflect.Valueof()获取的是Responder的指针，所以需要用Elem()方法获取指针对应的值
		if h_ref.Type().ConvertibleTo(r_ref.Type()) { //判断h_ref的类型能否转换成r_ref的类型
			r_ref.Set(h_ref)                                 //用反射的方式向handler赋值
			return r_ref.Interface().(Responder).RespondTo() //断言成Responder 并调用RespondTo方法
		}
	}
	return nil
}

//返回字符类型
type StringResponder func(context *gin.Context) string

func (this StringResponder) RespondTo() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.String(200, this(context))
	}
}

//返回Model类型，所以自定义的struct 都视为Model
type ModelResponder func(context *gin.Context) Model

func (this ModelResponder) RespondTo() gin.HandlerFunc { //返回实体类
	return func(context *gin.Context) {
		context.JSON(200, this(context))
	}
}

//返回Models类型，Model切片
type ModelsResponder func(context *gin.Context) Models

func (this ModelsResponder) RespondTo() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Writer.Header().Set("Content-Type", "application/json")
		context.Writer.WriteString(string(this(context)))
	}
}

//返回模版文件
type View string
type ViewResponder func(context *gin.Context) View

func (this ViewResponder) RespondTo() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.HTML(200, string(this(context))+".html", nil)
	}
}
