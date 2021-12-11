package main

import (
	"fmt"
	"time"
)

// 模拟订单对象
type OrderInfo struct {
	id int
}

// 生产订单--生产者
func PAndC_Demo2_producer(out chan<- OrderInfo) {
	// 业务生成订单
	for i := 0; i < 10; i++ {
		order := OrderInfo{id: i + 1}
		fmt.Println("生成订单，订单ID为：", order.id)
		out <- order // 写入channel
	}
	// 如果不关闭，消费者就会一直阻塞，等待读，正式场景不需要主动关闭
	// 所有测试订单生成完毕，关闭channel
	close(out)
}

// 处理订单--消费者
func PAndC_Demo2_consumer(in <-chan OrderInfo) {
	// 从channel读取订单，并处理
	for order := range in {
		fmt.Println("读取订单，订单ID为：", order.id)
	}
}

func PAndC_Demo2_main() {
	ch := make(chan OrderInfo, 5)
	go PAndC_Demo2_producer(ch)
	go PAndC_Demo2_consumer(ch)
	time.Sleep(time.Second * 2)
}
