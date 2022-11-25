package main

import (
	"fmt"
	"github.com/gosuri/uiprogress"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func ConsoleCmdProcessBar_main() {
	//命令行 进度条
	ProcessBar_Demo5()

}

func ProcessBar_Demo1() {
	uiprogress.Start()            // start rendering
	bar := uiprogress.AddBar(100) // Add a new bar

	// optionally, append and prepend completion and elapsed time
	bar.AppendCompleted()
	bar.PrependElapsed()

	for bar.Incr() {
		time.Sleep(time.Millisecond * 20)
	}
}

func ProcessBar_Demo2() {
	uiprogress.Start()
	var steps = []string{"downloading source", "installing deps", "compiling", "packaging", "seeding database", "deploying", "staring servers"}
	bar := uiprogress.AddBar(len(steps))

	// prepend the current step to the bar
	bar.PrependFunc(func(b *uiprogress.Bar) string {
		return "app: " + steps[b.Current()-1]
	})

	for bar.Incr() {
		time.Sleep(time.Second * 2)
	}

}

func ProcessBar_Demo3() {
	waitTime := time.Millisecond * 100
	uiprogress.Start()

	// start the progress bars in go routines
	var wg sync.WaitGroup

	bar1 := uiprogress.AddBar(20).AppendCompleted().PrependElapsed()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for bar1.Incr() {
			time.Sleep(waitTime)
		}
	}()

	bar2 := uiprogress.AddBar(40).AppendCompleted().PrependElapsed()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for bar2.Incr() {
			time.Sleep(waitTime)
		}
	}()

	time.Sleep(time.Second)
	bar3 := uiprogress.AddBar(20).PrependElapsed().AppendCompleted()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= bar3.Total; i++ {
			bar3.Set(i)
			time.Sleep(waitTime)
		}
	}()

	// wait for all the go routines to finish
	wg.Wait()
}

func ProcessBar_Demo4() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // use all available cpu cores

	// create a new bar and prepend the task progress to the bar and fanout into 1k go routines
	count := 1000
	bar := uiprogress.AddBar(count).AppendCompleted().PrependElapsed()
	bar.PrependFunc(func(b *uiprogress.Bar) string {
		return fmt.Sprintf("Task (%d/%d)", b.Current(), count)
	})

	uiprogress.Start()
	var wg sync.WaitGroup

	// fanout into go routines
	for i := 0; i < count; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(500)))
			bar.Incr()
		}()
	}
	time.Sleep(time.Second) // wait for a second for all the go routines to finish
	wg.Wait()
	uiprogress.Stop()
}

func ProcessBar_Demo5() {

	uiprogress.Start()

	UiBarMax := 1024

	var bar *uiprogress.Bar

	bar = uiprogress.AddBar(UiBarMax)
	bar.AppendFunc(func(b *uiprogress.Bar) string {
		return "file: xxxx.iso"
	})

	_ = bar.Set(0)

	fsize := 88888888
	readSize := 0

	//模拟读取
	for {
		//读取文件数据
		time.Sleep(time.Second)
		readSize += 4_194_304

		percentage := float64(readSize) / float64(fsize)
		done := int(float64(UiBarMax) * percentage)

		if done > UiBarMax {
			done = UiBarMax
		}

		_ = bar.Set(done)

		if done >= UiBarMax {
			//模拟读完了
			time.Sleep(time.Second)
			break
		}
	}
}
