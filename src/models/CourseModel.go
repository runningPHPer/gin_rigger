package models

type CourseModel struct {
	CId     int    `json:"c_id" gorm:"column:c_id" uri:"id" binding:"required,gt=0"`
	Content string `json:"content" gorm:"content"`
}

func NewCourseModel() *CourseModel {
	return &CourseModel{}
}

func (this *CourseModel) String() string {
	return "CourseModel"
}
