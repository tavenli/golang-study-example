package main

import (
	"fmt"
	"github.com/gogf/gf/os/grpool"
	"github.com/gogf/gf/os/gtimer"
	"time"
)

/*
GoFrame的grpool通过协程复用，能够节省内存。

结合我们的需求：
如果你的服务器内存不高或者业务场景对内存占用的要求更高，那就使用grpool。
如果服务器的内存足够，但是对耗时有较高的要求，就用原生的goroutine。

*/

func GoRoutinePool_main() {
	pool := grpool.New(100)

	//添加1千个任务
	for i := 0; i < 1000; i++ {
		_ = pool.Add(job)
	}

	fmt.Println("worker:", pool.Size()) //当前工作的协程数量
	fmt.Println("jobs:", pool.Jobs())   //当前池中待处理的任务数量

	gtimer.SetInterval(time.Second, func() {
		fmt.Println("worker:", pool.Size()) //当前工作的协程数
		fmt.Println("jobs:", pool.Jobs())   //当前池中待处理的任务数
	})

	//阻止进程结束，否则线程池的作业还没有跑完，程序就退出了
	select {}

}

//任务方法
func job() {
	time.Sleep(time.Second)
}
