package main

import (
	"fmt"
	"os"
	"sync"
)

func HangUp_Main_Demo1() {

	c := make(chan os.Signal)

	fmt.Println("start a server or do something")

	s := <-c
	fmt.Println(s)

}

func HangUp_Main_Demo2() {
	//使用将近0%的CPU，因为它会导致goroutine阻塞，
	//这意味着调度器会将其置于睡眠状态，并且它永远不会被唤醒
	select {}

}

func HangUp_Main_Demo3() {
	//使用100%的CPU，因为它连续执行循环迭代。
	for {
	}

}

func HangUp_Main_Demo4() {

	//WaitGroup总共有三个方法：Add(delta int), Done(), Wait()
	//Add：添加或者减少等待 goroutine 的数量
	//Done：相当于Add(-1)
	//Wait：执行阻塞，直到所有的WaitGroup数量变成 0

	wg := sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()

}

func HangUp_Main_Demo5() {

	//使用等待組
	var wg sync.WaitGroup
	doFunc := func(num int) {
		fmt.Println("doing", num)
		defer wg.Done()
	}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go doFunc(i)
	}
	wg.Wait()

}
