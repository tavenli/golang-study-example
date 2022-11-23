package main

import (
	"fmt"
	"time"
)

func Channel_main() {

	/*
		Go协程、并发、信道

		并行 = 多个任务同时执行，靠使用多线程来实现
		并发 = 多个线程同时竞争同一个内存数据或内存对象，则产生并发

		Go 协程使用信道（Channel）来进行通信

	*/

}

func demo1() {
	//创建一个匿名携程
	go func() {
		fmt.Println("do something")
	}()

}

func demo2() {
	//声明信道

	var c chan int // 方式一，为nil，不能发送也不能接受数据
	//c := make(chan int)    // 方式二，可用

	//为了让开发工具不报错，增加下面一行
	fmt.Println(c)
}

func demo3() {
	//信道使用，无缓冲信道

	c := make(chan int) // 写数据
	c <- 123

	// 读数据
	var a int
	a = <-c // 方式一
	<-c     // 方式二,读出来的数据丢弃不使用

	fmt.Println(a)
}

func printHello(c chan bool) {
	fmt.Println("hello world goroutine")
	<-c // 读取信道的数据
}

func demo4() {

	/*
		读/写数据的时候信道会阻塞，调度器会去调度其他可用的协程
		造成死锁的情况：
		只向chan写入数据，不读取，会死锁
		只对chan读取数据，不写入，会死锁
		fatal error: all goroutines are asleep - deadlock!

	*/

	c := make(chan bool)
	go printHello(c)
	c <- true // main 协程阻塞
	fmt.Println("main goroutine")

	/*
		输出：
		hello world goroutine
		main goroutine
	*/
}

func demo5() {
	//关闭信道与 for loop

	/*
		val, ok := <- channel
		val 是接收的值，ok 标识信道是否关闭。为 true 的话，该信道还可以进行读写操作；为 false 则标识信道关闭，数据不能传输。
		使用内置函数 close() 关闭信道。
	*/

	ch := make(chan int)
	go printNums(ch)

	for v := range ch {
		fmt.Println(v)
	}
}

func printNums(ch chan int) {
	for i := 0; i < 4; i++ {
		ch <- i
	}

	//关闭信道
	close(ch)
}

func demo6() {
	//缓冲信道和信道容量

	ch := make(chan int, 3)

	ch <- 7
	ch <- 8
	ch <- 9
	//ch <- 10

	/*
		注释打开的话，协程阻塞，发生死锁
		会发生死锁：信道已满且没有其他可用信道读取数据
	*/

	fmt.Println("main stopped")
}

func demo7() {
	//单向信道
	//这种信道主要用在信道作为参数传递的时候，Go 提供了自动转化，双向转单向

	rch := make(<-chan int)
	sch := make(chan<- int)

	//为了让开发工具不报错，增加下面一行
	fmt.Println(rch)
	fmt.Println(sch)
}

func demo8() {
	//读写nil通道都会死锁

	// 读nil通道:
	var dataStream chan interface{}

	<-dataStream

	// 写nil通道:
	var dataStream2 chan interface{}

	dataStream2 <- struct{}{}

	//因为chan没有初始化，为nil，则都会死锁
}

func demo9() {

	//close 值为nil的channel会panic

	//panic: close of nil channel

	var dataStream chan interface{}

	close(dataStream)

}

func demo10() {

	/*
		select 的使用

		select 可以安全的将channels与诸如取消、超时、等待和默认值之类的概念结合在一起。
		select看起来就像switch 包含了很多case,然而与switch不同的是：
		select块中的case语句没有顺序地进行测试，如果没有满足任何条件，执行不会自动失败,如果同时满足多个条件随机执行一个case。
		select就是用来监听和channel有关的IO操作，当 IO 操作发生时，触发相应的动作。

		1、case分支中必须是一个IO操作。
		2、当case分支不满足监听条件，阻塞当前case分支。
		3、如果同时有多个case分支满足，select随机选定一个执行（select底层实现，case对应一个Goroutine）。
		4、一次select监听，只能执行一个case分支，未执行的分支将被丢弃。通常将select放于for循环中。
		5、default在所有case均不满足时，默认执行的分组，为了防止忙轮询，通常将for中select中的default省略。

	*/

	//select基本用法：
	chan1 := make(chan int)
	chan2 := make(chan int)

	select {
	case <-chan1:
		// 如果chan1成功读到数据，则进行该case处理语句
	case chan2 <- 1:
		// 如果成功向chan2写入数据，则进行该case处理语句
	default:
		// 如果上面都没有成功，则进入default处理流程
	}

	//如果有一个或多个IO操作可以完成，则Go运行时系统会随机的选择一个执行，
	//如果没有一个case条件满足，但是有default分支，则执行default分支语句，
	//如果连default都没有，则select语句会一直阻塞，直到至少有一个IO操作可以进行。

	start := time.Now()
	c := make(chan interface{})
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		time.Sleep(4 * time.Second)
		close(c)
	}()

	go func() {
		time.Sleep(3 * time.Second)
		ch1 <- 3
	}()

	go func() {
		time.Sleep(3 * time.Second)
		ch2 <- 5
	}()

	fmt.Println("Blocking on read...")

	select {
	case <-c:
		fmt.Printf("Unblocked %v later.\n", time.Since(start))

	case <-ch1:
		fmt.Printf("ch1 case...")

	case <-ch2:
		fmt.Printf("ch1 case...")

	default:
		fmt.Printf("default go...")
	}

	//break关键字结束select
	//ch5和ch6两个通道都可以读取到值，所以系统会随机选择一个case执行。
	//我们发现选择执行ch5的case时，由于有break关键字只执行了一句

	ch5 := make(chan int, 1)
	ch6 := make(chan int, 1)

	ch5 <- 3
	ch6 <- 5

	select {
	case <-ch5:
		fmt.Println("ch5 selected.")
		break
		fmt.Println("ch5 selected after break")

	case <-ch6:
		fmt.Println("ch6 selected.")
		fmt.Println("ch6 selected without break")
	}

}

var ch1 chan int
var ch2 chan int
var chs = []chan int{ch1, ch2}
var numbers = []int{1, 2, 3, 4, 5}

func demo11() {
	//所有channel表达式都会被求值、所有被发送的表达式都会被求值。
	// 求值顺序：自上而下、从左到右

	select {
	case getChan(0) <- getNumber(2):
		fmt.Println("1th case is selected.")

	case getChan(1) <- getNumber(3):
		fmt.Println("2th case is selected.")

	default:
		fmt.Println("default!.")
	}

	/*
		输出：
		chs[0]
		numbers[2]
		chs[1]
		numbers[3]
		default!.

	*/
}

func getNumber(i int) int {
	fmt.Printf("numbers[%d]\n", i)
	return numbers[i]
}

func getChan(i int) chan int {
	fmt.Printf("chs[%d]\n", i)
	return chs[i]
}
