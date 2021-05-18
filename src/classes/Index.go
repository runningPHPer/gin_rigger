package classes

import (
	"gin_rigger/src/rigger"
	"github.com/gin-gonic/gin"
)

type IndexClass struct {
}

func NewIndexClass() *IndexClass {
	return &IndexClass{}
}

//业务方法
func (this *IndexClass) GetIndex() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(200, gin.H{"result": "success"})
	}
}

func (this *IndexClass) Build(rigger *rigger.Rigger) { //rigger传进来
	rigger.Handle("GET", "/", this.GetIndex()) //把内容隐藏在这里，main就很干净
}
