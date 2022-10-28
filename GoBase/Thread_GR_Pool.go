package main

import (
	"fmt"
	"github.com/gogf/gf/os/grpool"
	"github.com/gogf/gf/os/gtimer"
	"time"
)

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
