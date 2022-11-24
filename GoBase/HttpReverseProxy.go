package main

import (
	"log"
	"math/rand"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func ReverseProxy_main() {

	u, err := url.Parse("http://localhost:8080/test")
	if err != nil {
		log.Fatalf("url.Parse: %v", err)
	}
	proxy := httputil.NewSingleHostReverseProxy(u)
	log.Printf("Listening at :8081")
	if err := http.ListenAndServe(":8081", proxy); err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}

}

func NewMultipleHostsReverseProxy(targets []*url.URL) *httputil.ReverseProxy {
	director := func(req *http.Request) {
		target := targets[rand.Int()%len(targets)]
		req.URL.Scheme = target.Scheme
		req.URL.Host = target.Host
		req.URL.Path = target.Path
	}
	//反向代理
	return &httputil.ReverseProxy{Director: director}
}

func ReverseProxy_main2() {
	proxy := NewMultipleHostsReverseProxy([]*url.URL{
		{
			Scheme: "http",
			Host:   "localhost:9091",
		},
		{
			Scheme: "http",
			Host:   "localhost:9092",
		},
	})
	log.Fatal(http.ListenAndServe(":9090", proxy))
}
