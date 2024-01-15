package main

import (
	"fmt"
	"sync"
	"time"
)

type Data struct {
	Value int
}

type Queue struct {
	mutex      sync.Mutex
	cond       *sync.Cond
	buffer     []Data
	terminated bool
}

func NewQueue() *Queue {
	q := &Queue{}
	q.cond = sync.NewCond(&q.mutex)
	return q
}

func (q *Queue) Produce(data Data) {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	q.buffer = append(q.buffer, data)
	fmt.Printf("Produced: %d\n", data.Value)

	// 唤醒等待的消费者
	q.cond.Signal()
}

func (q *Queue) Consume() Data {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	// 等待数据可用
	for len(q.buffer) == 0 && !q.terminated {
		q.cond.Wait()
	}

	if len(q.buffer) > 0 {
		data := q.buffer[0]
		q.buffer = q.buffer[1:]
		fmt.Printf("Consumed: %d\n", data.Value)
		return data
	}

	return Data{}
}

func (q *Queue) Terminate() {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	q.terminated = true

	// 唤醒所有等待的消费者
	q.cond.Broadcast()
}

func PAndC_Demo4_main() {
	//https://juejin.cn/post/7235609253058904119
	//使用互斥锁和条件变量实现生产者消费者模式的示例代码

	queue := NewQueue()

	// 启动生产者
	for i := 1; i <= 3; i++ {
		go func(id int) {
			for j := 1; j <= 5; j++ {
				data := Data{Value: id*10 + j}
				queue.Produce(data)
				time.Sleep(time.Millisecond * 500) // 模拟生产时间
			}
		}(i)
	}

	// 启动消费者
	for i := 1; i <= 2; i++ {
		go func(id int) {
			for {
				data := queue.Consume()
				if data.Value == 0 {
					break
				}
				// 处理消费的数据
				time.Sleep(time.Millisecond * 1000) // 模拟处理时间
			}
		}(i)
	}

	// 等待一定时间后终止消费者
	time.Sleep(time.Second * 6)
	queue.Terminate()

	// 等待生产者和消费者完成
	time.Sleep(time.Second * 1)
}
