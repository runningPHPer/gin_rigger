// Code generated from /Users/yuzhonghua/Documents/Code/Go/gin_rigger/src/BeanExpr/BeanExpr.g4 by ANTLR 4.9.1. DO NOT EDIT.

package FuncExpr // BeanExpr
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseBeanExprVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseBeanExprVisitor) VisitStart(ctx *StartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeanExprVisitor) VisitMethodCall(ctx *MethodCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeanExprVisitor) VisitFuncCall(ctx *FuncCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBeanExprVisitor) VisitFuncArgs(ctx *FuncArgsContext) interface{} {
	return v.VisitChildren(ctx)
}
