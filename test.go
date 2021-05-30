package main

import (
	"fmt"
	"gin_rigger/src/BeanExpr/FuncExpr"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"log"
	"reflect"
	"strconv"
	"strings"
)

type FuncExprListener struct {
	*FuncExpr.BaseBeanExprListener
	funcName   string
	args       []reflect.Value
	methodName string //方法名称
	execType   uint8  //执行类型。0为默认执行函数 1为struct
}

var FuncMap map[string]interface{}

//获取方法名称
func (this *FuncExprListener) ExitFuncCall(ctx *FuncExpr.FuncCallContext) {
	this.funcName = ctx.GetStart().GetText()
}

func (this *FuncExprListener) ExitMethodCall(ctx *FuncExpr.MethodCallContext) {
	this.execType = 1 //类方法执行
	this.methodName = ctx.GetStart().GetText()
}

//获取参数
func (this *FuncExprListener) ExitFuncArgs(ctx *FuncExpr.FuncArgsContext) {
	for i := 0; i < ctx.GetChildCount(); i++ {
		token := ctx.GetChild(i).GetPayload().(*antlr.CommonToken)
		var value reflect.Value
		switch token.GetTokenType() {
		case FuncExpr.BeanExprLexerStringArg:
			stringArg := strings.Trim(token.GetText(), "'")
			value = reflect.ValueOf(stringArg)
			break
		case FuncExpr.BeanExprLexerIntArg:
			v, err := strconv.ParseInt(token.GetText(), 10, 64)
			if err != nil {
				panic("parse int64 error")
			}
			value = reflect.ValueOf(v)
			break
		case FuncExpr.BeanExprLexerFloatArg:
			v, err := strconv.ParseFloat(token.GetText(), 64)
			if err != nil {
				panic("parse float64 error")
			}
			value = reflect.ValueOf(v)
			break
		default:
			continue
		}
		this.args = append(this.args, value)
	}
}

//查询struct对应属性
func (this *FuncExprListener) findField(method string, v reflect.Value) reflect.Value {
	if v.Kind() == reflect.Ptr { //是指针类型
		v = v.Elem() //获取指针指向的类型
	}
	if filed := v.FieldByName(method); filed.IsValid() {
		return filed
	}
	return reflect.Value{}
}

func (this *FuncExprListener) Run() {
	switch this.execType {
	case 0:
		if f, ok := FuncMap[this.funcName]; ok {
			v := reflect.ValueOf(f)
			if v.Kind() == reflect.Func {
				v.Call(this.args)
			}
		}
		break
	case 1:
		ms := strings.Split(this.methodName, ".")
		if obj, ok := FuncMap[ms[0]]; ok {
			objv := reflect.ValueOf(obj)
			currentv := objv
			for i := 1; i < len(ms); i++ {
				if i == len(ms)-1 { //最后一个
					if method := currentv.MethodByName(ms[i]); !method.IsValid() {
						panic("method error:" + ms[i]) //抛出错误
					} else {
						method.Call(this.args) //使用输入参数跳用函数
					}
					break
				} else { //不是最后一个
					field := this.findField(ms[i], currentv)
					if field.IsValid() {
						currentv = field
					} else {
						panic("field error " + ms[i])
					}
				}
			}
		}
		break
	default:
		log.Println("nothing to do")
	}
}

type Admin struct {
}

func (this *Admin) AdminName() {
	fmt.Println("admin")
}

type User struct {
	Adm *Admin
}

func (this *User) Name(name string) {
	fmt.Println("my name is " + name)
}

func main() {
	FuncMap = map[string]interface{}{
		"test": func(name string, age int64) {
			log.Println("this is ", name, " and name is ", age)
		},
		"User": &User{},
	}
	is := antlr.NewInputStream("User.Adm.AdminName()")
	lexer := FuncExpr.NewBeanExprLexer(is)
	ts := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := FuncExpr.NewBeanExprParser(ts)
	lis := &FuncExprListener{}
	antlr.ParseTreeWalkerDefault.Walk(lis, p.Start())

	lis.Run()

}
