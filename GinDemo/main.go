package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("------------")

	// 禁用控制台颜色
	gin.DisableConsoleColor()

	// 创建记录日志的文件
	f, _ := os.Create("gin.log")
	//gin.DefaultWriter = io.MultiWriter(f)

	// 如果需要将日志同时写入文件和控制台，请使用以下代码
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	router := gin.Default()
	router.Static("/static", "./static")

	//router.Delims("{.{", "}.}")
	//router.LoadHTMLGlob("./templates/**/*")
	//router.Use(ShowRequestInfo())

	router.GET("/index", func(c *gin.Context) {

		c.JSON(200, gin.H{
			"code": 0,
			"msg":  "ok",
		})
	})

	router.POST("/uploadFile", func(c *gin.Context) {

		//获取表单数据 参数为name值
		f, err := c.FormFile("file")
		//错误处理
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		} else {
			//将文件保存至本项目根目录中
			c.SaveUploadedFile(f, f.Filename)
			//保存成功返回正确的Json数据
			c.JSON(http.StatusOK, gin.H{
				"message": "OK",
			})
		}
	})

	router.POST("/uploadFiles", func(c *gin.Context) {
		//router := gin.Default()
		// 8 MiB 设置最大的上传文件的大小
		//router.MaxMultipartMemory = 8 << 20

		//多文件上传
		form, err := c.MultipartForm()
		files := form.File["files"]
		//错误处理
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		}
		for _, f := range files {
			fmt.Println(f.Filename)
			c.SaveUploadedFile(f, f.Filename)
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	})

	//
	attachWebsocket(router)

	router.Run(":7070")

}

//Gin中加入 websocket
var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func wsHello(c *gin.Context) {
	//升级get请求为webSocket协议
	wsConn, err := upGrader.Upgrade(c.Writer, c.Request, nil)

	if err != nil {
		return
	}

	defer wsConn.Close()

	for {
		//读取ws中的数据
		mt, message, err := wsConn.ReadMessage()
		if err != nil {
			break
		}
		if string(message) == "ping" {
			message = []byte("pong")
		}
		//写入ws数据
		err = wsConn.WriteMessage(mt, message)
		if err != nil {
			break
		}
	}

}

func attachWebsocket(router *gin.Engine) {
	//Gin中加入 websocket
	//ws://127.0.0.1:7070/wsHello

	//在线测试客户端 http://www.websocket-test.com

	router.GET("/wsHello", wsHello)
}

func clientWebsocket() {
	url := "ws://127.0.0.1:7070/wsHello"
	wsConn, _, err := websocket.DefaultDialer.Dial(url, nil)

	if err != nil {
		fmt.Println(err)
	}

	go func() {
		for {
			err := wsConn.WriteMessage(websocket.BinaryMessage, []byte("ping"))
			if err != nil {
				fmt.Println(err)
			}
			time.Sleep(time.Second * 2)
		}
	}()

	for {
		_, data, err := wsConn.ReadMessage()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("receive: ", string(data))
	}

}
