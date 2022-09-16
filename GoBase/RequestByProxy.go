package main

import (
	"crypto/tls"
	"fmt"
	"golang.org/x/net/proxy"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

func ReqByProxy_Main() {
	//通过代理发起请求

	ReqByHttpsProxy_Demo()

	//ReqBySocks5Proxy_Demo()
}

func ReqByHttpProxy_Demo() {

	req_url := "http://httpbin.org/get"
	fmt.Printf("url: %s", req_url)

	proxy := func(_ *http.Request) (*url.URL, error) {
		return url.Parse("http://127.0.0.1:8080")
	}

	transport := &http.Transport{Proxy: proxy}

	//或者 通过环境变量方式 ProxyFromEnvironment
	//支持 http、https、socks5
	//transport := &http.Transport{Proxy: http.ProxyFromEnvironment}

	c := &http.Client{Transport: transport}

	req, err := http.NewRequest("GET", req_url, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, err := c.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", body)
}

func ReqByHttpsProxy_Demo() {
	//

	req_url := "https://httpbin.org/get"
	fmt.Printf("url: %s", req_url)

	proxy := func(_ *http.Request) (*url.URL, error) {
		return url.Parse("http://127.0.0.1:8080")
	}

	//跳过证书验证
	transport := &http.Transport{Proxy: proxy, TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}

	c := &http.Client{Transport: transport}

	req, err := http.NewRequest("GET", req_url, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, err := c.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", body)
}

func ReqByEnvProxy_Demo() {

	req_url := "http://httpbin.org/get"
	fmt.Printf("url: %s", req_url)

	//可以设置在系统的环境变量中，也可以在程序执行时通过代码设置
	os.Setenv("HTTP_PROXY", "http://127.0.0.1:8080")
	os.Setenv("HTTPS_PROXY", "https://127.0.0.1:8080")

	//os.Setenv("HTTP_PROXY", "socks5://1.1.1.1:1080")
	//os.Setenv("HTTPS_PROXY", "socks5://1.1.1.1:1080")

	//或者 通过环境变量方式 ProxyFromEnvironment
	//支持 http、https、socks5
	transport := &http.Transport{Proxy: http.ProxyFromEnvironment}

	c := &http.Client{Transport: transport}

	req, err := http.NewRequest("GET", req_url, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, err := c.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", body)
}

func ReqBySocks5Proxy_Demo() {

	dialer, _ := proxy.SOCKS5("tcp", "127.0.0.1:10808", nil, proxy.Direct)

	httpTransport := &http.Transport{}
	httpTransport.Dial = dialer.Dial

	httpClient := &http.Client{Transport: httpTransport}
	// set our socks5 as the dialer

	if resp, err := httpClient.Get("http://httpbin.org/get"); err != nil {
		log.Fatalln(err)
	} else {
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Printf("%s", body)
	}

}

func ReqBySocks5Proxy_Demo2() {
	//参考使用内部包实现 net/http/socks_bundle.go 的 http.socksNewDialer

	transport := &http.Transport{Proxy: http.ProxyFromEnvironment}

	fmt.Println(transport)
}

func ReqByTcpProxy_Demo() {
	//使用代理，发起 net.Dial 请求
	//实际本质就是实现一个sockes5的客户端，每次请求要通过 sockes5的客户端 发给sockes5代理服务端

}
