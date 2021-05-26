package classes

import (
	"gin_rigger/src/models"
	"gin_rigger/src/rigger"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"time"
)

type CourseClass struct {
	*rigger.GormAdapter
}

func NewCourseClass() *CourseClass {
	return &CourseClass{}
}

//详情
func (this *CourseClass) Detail(context *gin.Context) rigger.Model {
	course := models.NewCourseModel()
	rigger.Error(context.ShouldBindUri(course))
	rigger.Error(this.Table("sc_subject_course").Where("c_id = ?", course.CId).Find(course).Error)
	rigger.Task(this.UpdateFalse, func() { //这里使用匿名函数向回调函数传值
		this.UpdateFalseDown(course.CId)
	}, course.CId) //执行协程异步任务
	return course
}

func (this *CourseClass) UpdateFalse(params ...interface{}) {
	time.Sleep(time.Second * 3)
	this.Table("sc_subject_course").Where("c_id = ?", params[0]).Update("false_num", gorm.Expr("false_num+1"))

}

func (this *CourseClass) UpdateFalseDown(id int) {
	log.Println("错误更改结束！id为：", id)
}

//挂载路由
func (this *CourseClass) Build(rigger *rigger.Rigger) {
	rigger.Handle("GET", "/course/:id", this.Detail)
}
