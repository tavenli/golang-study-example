package main

import (
	"fmt"
	"time"
)

func producer(ch chan<- int) {
	for i := 1; i <= 5; i++ {
		ch <- i // 将数据发送到通道
		fmt.Println("生产者生产：", i)
		time.Sleep(time.Second) // 模拟生产过程
	}
	close(ch) // 关闭通道
}

func consumer(ch <-chan int, done chan<- bool) {
	for num := range ch {
		fmt.Println("消费者消费：", num)
		time.Sleep(2 * time.Second) // 模拟消费过程
	}
	done <- true // 通知主线程消费者已完成
}

func PAndC_Demo3_main() {
	ch := make(chan int, 3) // 创建带缓冲的通道
	done := make(chan bool) // 用于通知主线程消费者已完成

	go producer(ch)       // 启动生产者goroutine
	go consumer(ch, done) // 启动消费者goroutine

	// 主线程等待消费者完成
	<-done
	fmt.Println("消费者已完成")

	// 主线程结束，程序退出
}
