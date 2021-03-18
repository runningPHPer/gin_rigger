package classes

import "github.com/gin-gonic/gin"

type IndexClass struct {
	*gin.Engine
}

//所谓的构造函数
func NewIndexClass(engine *gin.Engine) *IndexClass {
	return &IndexClass{Engine: engine} //指针需要赋值
}

//业务方法
func (this *IndexClass) GetIndex() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(200, gin.H{"result": "success"})
	}
}

func (this *IndexClass) Build() {
	this.Handle("GET", "/", this.GetIndex())
}
