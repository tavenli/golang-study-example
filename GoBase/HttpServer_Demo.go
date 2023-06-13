package main

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"golang.org/x/net/http2"
	"net/http"
	//"net/http2"
)

func HttpServer_main() {

	http.HandleFunc("/hello", SayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		logs.Error("ListenAndServe: ", err)
	}
}

func HttpServer2_main() {
	/*
		go语言 支持 HTTP/1.1 和 HTTP/2.0
		分别提供了两个库，net/http  和 net/http2
		从 go 1.6 开始支持 net/http2，默认启动为使用 HTTP/2.0
		低于 1.6 版本的，需要引入 go get golang.org/x/net/http2

		https://http2.github.io/

	*/

	server := http.Server{
		Addr:    ":9090",
		Handler: http.HandlerFunc(SayHello),
	}

	//使用 HTTP/2.0
	http2.VerboseLogs = true
	err1 := http2.ConfigureServer(&server, &http2.Server{})
	fmt.Println(err1)

	//虽然 HTTP/2 并没有规定不支持 http，但是目前业内和主流的浏览器都只支持 HTTPS
	//err := server.ListenAndServe()
	err := server.ListenAndServeTLS("D:\\test.pem", "D:\\test.key")

	if err != nil {
		logs.Error("ListenAndServe: ", err)
	}

}

func SayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}
