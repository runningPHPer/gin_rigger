package classes

import (
	"gin_rigger/src/models"
	"gin_rigger/src/rigger"
	"github.com/gin-gonic/gin"
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
	return course
}

//挂载路由
func (this *CourseClass) Build(rigger *rigger.Rigger) {
	rigger.Handle("GET", "/course/:id", this.Detail)
}
