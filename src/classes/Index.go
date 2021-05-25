package classes

import (
	"fmt"
	"gin_rigger/src/rigger"
	"github.com/gin-gonic/gin"
)

type IndexClass struct {
}

func NewIndexClass() *IndexClass {
	return &IndexClass{}
}

//业务方法
func (this *IndexClass) GetIndex(context *gin.Context) rigger.View {
	fmt.Println(123)
	return "index"
}

func (this *IndexClass) Build(rigger *rigger.Rigger) { //rigger传进来
	rigger.Handle("GET", "/index", this.GetIndex) //把内容隐藏在这里，main就很干净
}
