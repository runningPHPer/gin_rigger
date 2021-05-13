package middlewares

import "fmt"

type UserMid struct {
}

func NewUserMid() *UserMid {
	return &UserMid{}
}

func (this *UserMid) OnRequest() error {
	fmt.Println("这是中间件")
	return fmt.Errorf("这这是强制错误！")
}
