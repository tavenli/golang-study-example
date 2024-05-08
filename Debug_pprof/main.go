package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	_ "net/http/pprof"
	"strings"
	"time"
)

/*
	分析工具
	go tool pprof http://127.0.0.1:8070/debug/pprof/heap

	开启图形化web
	go tool pprof -http=:8088 http://127.0.0.1:8070/debug/pprof/heap

	开启成功后，可打开ui
	http://localhost:8088/ui/

*/

func init() {
	//异步协程开启 pprof
	go open_pprof()
}
func main() {
	engine := gin.Default()
	engine.GET("/show", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"msg": "success",
		})
	})
	testPprofHeap()
	_ = engine.Run(":8060")
}

func open_pprof() {
	if err := http.ListenAndServe(":8070", nil); err != nil {
		fmt.Println("pprof err:", err)
	}
}

// 模拟内存使用增加
func testPprofHeap() {
	go func() {
		var stringSlice []string
		for {
			time.Sleep(time.Second * 2)
			repeat := strings.Repeat("hello,world", 50000)
			stringSlice = append(stringSlice, repeat)
			fmt.Printf("time:%d length:%d \n", time.Now().Unix(), len(stringSlice))
		}
	}()
}
