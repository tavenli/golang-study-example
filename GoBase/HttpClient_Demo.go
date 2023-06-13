package main

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"golang.org/x/net/http2"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func HttpClient_main() {

	client := &http.Client{}

	httpReq, err := http.NewRequest("GET", "https://t.pilicat.com/echo/ReqInfo", strings.NewReader(""))
	response, err := client.Do(httpReq)
	fmt.Println(err)
	bodyBytes, err := ioutil.ReadAll(response.Body)
	result := string(bodyBytes)
	logs.Debug("Http result: ", result)

}

func HttpClient2_main() {

	//使用 HTTP/2.0
	transport := &http2.Transport{}
	client := &http.Client{Transport: transport, Timeout: 10 * time.Second}

	httpReq, err := http.NewRequest("GET", "https://a.tool.lu/ta.js", strings.NewReader(""))
	response, err := client.Do(httpReq)
	fmt.Println(err)
	bodyBytes, err := ioutil.ReadAll(response.Body)
	result := string(bodyBytes)
	logs.Debug("Http result: ", result)
}

func HttpClient2_Proxy_main() {
	proxy := func(_ *http.Request) (*url.URL, error) {
		return url.Parse("https://127.0.0.1:8080")
	}

	//使用 HTTP/2.0 通过代理方式（暂时有问题）
	transport, err1 := http2.ConfigureTransports(&http.Transport{Proxy: proxy})
	fmt.Println(err1)

	client := &http.Client{Transport: transport, Timeout: 10 * time.Second}

	httpReq, err := http.NewRequest("GET", "https://a.tool.lu/ta.js", strings.NewReader(""))
	response, err := client.Do(httpReq)
	fmt.Println(err)
	bodyBytes, err := ioutil.ReadAll(response.Body)
	result := string(bodyBytes)
	logs.Debug("Http result: ", result)
}
