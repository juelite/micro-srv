package logics

import (
	"time"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	toolbox "github.com/astaxie/beego/toolbox"
)

func init()  {
	register_task()
}

func register_task() {
	//call_count_pop := toolbox.NewTask("call_count_pop", "* */5 * * * *", call_count_pop)
	call_count_pop := toolbox.NewTask("call_count_pop", "*/10 * * * * *", call_count_pop)
	toolbox.AddTask("call_count_pop", call_count_pop)
	toolbox.StartTask()
}

func call_count_pop() error {
	CallIncrDb()
	fmt.Println("调用数据同步成功！ " + fmt.Sprint(time.Now()))
	return nil
}