package rigger

import (
	"io/ioutil"
	"log"
	"os"
)

//读取配置文件
func LoadConfigFile() []byte {
	dir, _ := os.Getwd() //获取根目录
	file := dir + "/application.yaml"
	b, err := ioutil.ReadFile(file)
	if err != nil {
		log.Println(err)
		return nil
	}
	return b
}
