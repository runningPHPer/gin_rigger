package rigger

import "sync"

//rigger的任务处理
//任务方法的定义
type TaskFunc func(params ...interface{})

var taskList chan *TaskExecutor //任务列表,用chan表示

var once sync.Once //golang的单例模式

//获取任务列表
func getTaskList() chan *TaskExecutor {
	once.Do(func() {
		taskList = make(chan *TaskExecutor) //初始化chan
	})
	return taskList
}

//包的初始化
func init() {
	taskList := getTaskList() //得到任务列表
	go func() {
		for task := range taskList {
			task.Exec() //执行任务
		}
	}()
}

//任务的执行者
type TaskExecutor struct {
	function TaskFunc
	params   []interface{}
}

//初始化任务执行者
func NewTaskExecutor(function TaskFunc, params []interface{}) *TaskExecutor {
	return &TaskExecutor{function: function, params: params}
}

func (this *TaskExecutor) Exec() { //执行任务
	this.function(this.params...)
}

//执行一次任务
func Task(function TaskFunc, params ...interface{}) {
	getTaskList() <- NewTaskExecutor(function, params) //向chan里面塞入任务
}
