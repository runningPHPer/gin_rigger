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
			doTask(task)
		}
	}()
}

//执行任务
func doTask(task *TaskExecutor) {
	go func() {
		//回调函数
		defer func() {
			if task.callback != nil {
				task.callback() //执行回调任务
			}
		}()
		task.Exec() //执行任务
	}()
}

//任务的执行者
type TaskExecutor struct {
	function TaskFunc
	params   []interface{}
	callback func()
}

//初始化计划任务
func NewTaskExecutor(function TaskFunc, params []interface{}, callback func()) *TaskExecutor {
	return &TaskExecutor{function: function, params: params, callback: callback}
}

func (this *TaskExecutor) Exec() { //执行任务
	this.function(this.params...)
}

//执行一次任务
func Task(function TaskFunc, callback func(), params ...interface{}) {
	if function == nil {
		return
	}
	//塞入任务时，以协程的方式加入
	go func() {
		getTaskList() <- NewTaskExecutor(function, params, callback) //向chan里面塞入任务
	}()
}
