package rigger

import "github.com/gin-gonic/gin"

//统一错误处理中间件
func ErrorHandle() gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			if e := recover(); e != nil { //捕获错误
				context.AbortWithStatusJSON(400, gin.H{"error": e}) //捕获错误，返回
			}
		}()
		context.Next()
	}
}

//错误方法
func Error(err error, msg ...string) {
	if err == nil {
		return
	} else {
		errMsg := err.Error()
		if len(msg) > 0 {
			errMsg = msg[0] //只取第一个
		}
		panic(errMsg)
	}
}
