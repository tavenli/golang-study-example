package main

import (
	"fmt"
	resty "github.com/go-resty/resty/v2"
	colly "github.com/gocolly/colly/v2"
	"os"
	"strconv"
	"time"
)

func GithubLibs_main() {

	//
	//colly_main()

	//
	resty_main()

}

func colly_main() {
	//自动爬虫框架

	c := colly.NewCollector()

	// Find and visit all links
	c.OnHTML("a", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	//c.Visit("http://go-colly.org/")
	c.Visit("https://t.pilicat.com")
}

func resty_main() {

	// Create a Resty Client
	client := resty.New()

	resp, err := client.R().EnableTrace().Get("https://httpbin.org/get")

	// Explore response object
	fmt.Println("Response Info:")
	fmt.Println("  Error      :", err)
	fmt.Println("  Status Code:", resp.StatusCode())
	fmt.Println("  Status     :", resp.Status())
	fmt.Println("  Proto      :", resp.Proto())
	fmt.Println("  Time       :", resp.Time())
	fmt.Println("  Received At:", resp.ReceivedAt())
	fmt.Println("  Body       :\n", resp)
	fmt.Println()

	// Explore trace info
	fmt.Println("Request Trace Info:")
	ti := resp.Request.TraceInfo()
	fmt.Println("  DNSLookup     :", ti.DNSLookup)
	fmt.Println("  ConnTime      :", ti.ConnTime)
	fmt.Println("  TCPConnTime   :", ti.TCPConnTime)
	fmt.Println("  TLSHandshake  :", ti.TLSHandshake)
	fmt.Println("  ServerTime    :", ti.ServerTime)
	fmt.Println("  ResponseTime  :", ti.ResponseTime)
	fmt.Println("  TotalTime     :", ti.TotalTime)
	fmt.Println("  IsConnReused  :", ti.IsConnReused)
	fmt.Println("  IsConnWasIdle :", ti.IsConnWasIdle)
	fmt.Println("  ConnIdleTime  :", ti.ConnIdleTime)
	fmt.Println("  RequestAttempt:", ti.RequestAttempt)
	fmt.Println("  RemoteAddr    :", ti.RemoteAddr.String())

	result := ""

	resp, err = client.R().
		SetQueryParams(map[string]string{
			"page_no": "1",
			"limit":   "20",
			"sort":    "name",
			"order":   "asc",
			"random":  strconv.FormatInt(time.Now().Unix(), 10),
		}).
		SetHeader("Accept", "application/json").
		SetAuthToken("BC594900518B4F7EAC75BD37F019E08FBC594900518B4F7EAC75BD37F019E08F").
		Get("/search_result")

	// Sample of using Request.SetQueryString method
	resp, err = client.R().
		SetQueryString("productId=232&template=fresh-sample&cat=resty&source=google&kw=buy a lot more").
		SetHeader("Accept", "application/json").
		SetAuthToken("BC594900518B4F7EAC75BD37F019E08FBC594900518B4F7EAC75BD37F019E08F").
		Get("/show_product")

	// If necessary, you can force response content type to tell Resty to parse a JSON response into your struct
	resp, err = client.R().
		SetResult(result).
		ForceContentType("application/json").
		Get("v2/alpine/manifests/latest")

	// POST JSON string
	// No need to set content type, if you have client level setting
	resp, err = client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(`{"username":"testuser", "password":"testpass"}`).
		SetResult(&UserData{}). // or SetResult(AuthSuccess{}).
		Post("https://myapp.com/login")

	// POST []byte array
	// No need to set content type, if you have client level setting
	resp, err = client.R().
		SetHeader("Content-Type", "application/json").
		SetBody([]byte(`{"username":"testuser", "password":"testpass"}`)).
		SetResult(&UserData{}). // or SetResult(AuthSuccess{}).
		Post("https://myapp.com/login")

	// POST Struct, default is JSON content type. No need to set one
	resp, err = client.R().
		SetBody(UserData{UserId: 1111, UserName: "testpass"}).
		SetResult(&UserData{}). // or SetResult(AuthSuccess{}).
		SetError(&UserData{}). // or SetError(AuthError{}).
		Post("https://myapp.com/login")

	// POST Map, default is JSON content type. No need to set one
	resp, err = client.R().
		SetBody(map[string]interface{}{"username": "testuser", "password": "testpass"}).
		SetResult(&UserData{}). // or SetResult(AuthSuccess{}).
		SetError(&UserData{}). // or SetError(AuthError{}).
		Post("https://myapp.com/login")

	// POST of raw bytes for file upload. For example: upload file to Dropbox
	fileBytes, _ := os.ReadFile("/Users/jeeva/mydocument.pdf")

	// See we are not setting content-type header, since go-resty automatically detects Content-Type for you
	resp, err = client.R().
		SetBody(fileBytes).
		SetContentLength(true). // Dropbox expects this value
		SetAuthToken("<your-auth-token>").
		SetError(&UserData{}). // or SetError(DropboxError{}).
		Post("https://content.dropboxapi.com/1/files_put/auto/resty/mydocument.pdf") // for upload Dropbox supports PUT too

}
