package rigger

//用户规范中间件的代码和功能的接口
//fairing 意为整流罩有保护之意
type Fairing interface {
	OnRequest() error
}
