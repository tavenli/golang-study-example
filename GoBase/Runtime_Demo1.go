package main

import (
	"fmt"
	"runtime"
	"time"
)

/**
runtime.Caller 返回函数栈信息

func Caller(skip int) (pc uintptr, file string, line int, ok bool)
参数：skip是要提升的堆栈帧数，0-当前函数，1-上一层函数，….

返回值：

pc 是 uintptr 这个返回的是函数指针
file 是函数所在文件名目录
line 所在行号
ok 是否可以获取到信息
*/
func Runtime_Demo1_main() {
	for i := 0; i < 4; i++ {
		rtest(i)
	}
}

func rtest(skip int) {
	call(skip)
}

func call(skip int) {
	pc, file, line, ok := runtime.Caller(skip)
	pcName := runtime.FuncForPC(pc).Name() //获取函数名
	fmt.Println(fmt.Sprintf("%v   %s   %d   %t   %s", pc, file, line, ok, pcName))
}

func call2() {
	//另一个类似的函数 runtime.Callers

	// 取出调用栈50个数据
	pcSlice := make([]uintptr, 50)
	count := runtime.Callers(2, pcSlice)

	pcSlice = pcSlice[:count] // 可能没有 50 个

	// 返回的 *Frames 类似一个迭代器
	frames := runtime.CallersFrames(pcSlice)
	var frame runtime.Frame
	more := count > 0
	for more {
		frame, more = frames.Next()
		fmt.Println(fmt.Sprintf("%v   %s   %d   %s", frame.PC, frame.File, frame.Line, frame.Function))
	}
}

func Gosched_Demo() {
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("goroutine....")
		}
	}()

	for i := 0; i < 4; i++ {
		//让出时间片，先让别的协议执行，它执行完，再回来执行此协程
		runtime.Gosched()
		fmt.Println("main....")
	}
}

func Goexit_Demo() {
	//创建新建的协程
	go func() {
		fmt.Println("goroutine开始...")

		//调用了别的函数
		fun()

		fmt.Println("goroutine结束...")
	}() //别忘了()

	//睡一会儿，不让主协程结束
	time.Sleep(3 * time.Second)
}

func fun() {
	defer fmt.Println("defer...")

	//return           //终止此函数
	runtime.Goexit() //终止所在的协程

	fmt.Println("fun函数....")
}
