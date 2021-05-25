package funcs

import "html/template"

//自定义模版函数
func Strong(text string) template.HTML {
	return template.HTML("<strong>" + text + "<strong>")
}

func Test() string {
	return "TEST"
}
