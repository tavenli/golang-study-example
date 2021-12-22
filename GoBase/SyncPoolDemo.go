package main

import (
	"fmt"
	"sync"
)

func SyncPool_Demo1_main() {
	//sync.Pool 的作用：
	//复用已经使用过的对象，减少对象的创建和回收的时间

	//创建
	namePool := sync.Pool{
		New: func() interface{} {
			return "Taven"
		},
	}

	//借
	name := namePool.Get()
	fmt.Println(name)

	//还
	namePool.Put(name)


}

func SyncPool_Demo2_main()  {

	//创
	bufPoolMax := sync.Pool{
		New: func() interface{} {
			return make([]byte, 65535)
		},
	}

	//借
	buf := bufPoolMax.Get().([]byte)
	fmt.Println(buf)

	//还
	bufPoolMax.Put(buf)

}