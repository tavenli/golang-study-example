package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func HtmlTemplate_main() {

	//
	//HtmlTemplate_Demo1()

	//
	HtmlTemplate_Demo2()

}

func HtmlTemplate_Demo1() {

	type Inventory struct {
		Material string
		Count    uint
	}
	sweaters := Inventory{"wool", 17}

	data := make(map[string]interface{})
	data["sweaters"] = sweaters
	data["currentTime"] = time.Now()

	tmpl := template.New("test")

	tmpl.Funcs(template.FuncMap{
		"FormatAsDate": HtmlTemplate_FormatAsDate,
		"GetMapVal":    HtmlTemplate_GetMapVal,
	})

	tmpl, err := tmpl.Parse("{{.sweaters.Count}} items are made of {{.sweaters.Material}} {{FormatAsDate .currentTime}}")
	if err != nil {
		panic(err)
	}

	//err = tmpl.Execute(os.Stdout, data)

	buf := new(bytes.Buffer)
	err = tmpl.Execute(buf, data)
	_html := buf.String()
	fmt.Println(_html)

	if err != nil {
		panic(err)
	}

}

func HtmlTemplate_Demo2() {

	var users []*UserData
	for i := 0; i < 10; i++ {
		users = append(users, &UserData{UserId: i, UserName: "taven"})
	}

	data := make(map[string]interface{})
	data["currentTime"] = time.Now()
	data["users"] = users
	data["score"] = 66

	tmpl := template.New("dailyReport.html")

	tmpl.Funcs(template.FuncMap{
		"FormatAsDate": HtmlTemplate_FormatAsDate,
		"GetMapVal":    HtmlTemplate_GetMapVal,
		"AppName":      HtmlTemplate_AppName,
		"Sum":          HtmlTemplate_Sum,
		"CurTimeByFm":  HtmlTemplate_CurTimeByFm,
	})

	//workPath, _ := os.Getwd()
	tmpl, err := tmpl.ParseFiles("./views/dailyReport.html")
	if err != nil {
		panic(err)
	}

	//err = tmpl.Execute(os.Stdout, data)

	buf := new(bytes.Buffer)
	err = tmpl.Execute(buf, data)
	_html := buf.String()
	fmt.Println(_html)

	if err != nil {
		panic(err)
	}

}

func HtmlTemplate_AppName() string {
	return "GoExample"
}

func HtmlTemplate_Sum(num1, num2 int) int {
	return num1 + num2
}

func HtmlTemplate_CurTimeByFm(format string) string {

	return FormatTimeByFm(time.Now(), format)
}

func HtmlTemplate_FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%04d-%02d-%02d", year, month, day)
}

func HtmlTemplate_GetMapVal(mapData map[string]interface{}, key string) interface{} {
	if val, ok := mapData[key]; ok {
		//存在
		return val
	}
	return nil
}

func HtmlTemplateStartServe() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/404", page_not_found)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println("ListenAndServer: ", err)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	//解析模板文件
	t, _ := template.ParseFiles("hello.html")
	//执行模板
	t.Execute(w, "Hello world")

}

func page_not_found(rw http.ResponseWriter, r *http.Request) {
	//t, _ := template.New("404").ParseFiles("d:\\404.html")
	content, _ := ioutil.ReadFile(fmt.Sprintf("%s/%s", "d:\\temp", "404.html"))
	t, _ := template.New("404").Parse(string(content))
	data := make(map[string]interface{})
	data["content"] = "page not found"
	t.Execute(rw, data)
}
