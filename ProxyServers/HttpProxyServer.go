package main

import (
	"fmt"
	"github.com/elazarl/goproxy"
	"net/http"
)

func main() {

	//startHttpProxyServer()

	//startHttpProxyServer2()

	startSocks5ProxyServer()

}

func startHttpProxyServer() {

	proxy := goproxy.NewProxyHttpServer()
	//打印转发日志
	proxy.Verbose = true

	server := &http.Server{Addr: ":8080", Handler: proxy}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("ListenAndServe error", err)
	}

}

func startSocks5ProxyServer() {
	//不带密码
	NewSocks5Server(":1080", 0, "", "").Start()

	//有密码的代理服务器
	//NewSocks5Server(":1080", 2, "test", "123456").Start()
}
