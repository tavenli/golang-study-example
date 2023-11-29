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

/*
	var builtins = FuncMap{
	    // 返回第一个为空的参数或最后一个参数。可以有任意多个参数。
	    // "and x y"等价于"if x then y else x"
	    "and": and,
	    // 显式调用函数。第一个参数必须是函数类型，且不是template中的函数，而是外部函数。
	    // 例如一个struct中的某个字段是func类型的。
	    // "call .X.Y 1 2"表示调用dot.X.Y(1, 2)，Y必须是func类型，函数参数是1和2。
	    // 函数必须只能有一个或2个返回值，如果有第二个返回值，则必须为error类型。
	    "call": call,
	    // 返回与其参数的文本表示形式等效的转义HTML。
	    // 这个函数在html/template中不可用。
	    "html": HTMLEscaper,
	    // 对可索引对象进行索引取值。第一个参数是索引对象，后面的参数是索引位。
	    // "index x 1 2 3"代表的是x[1][2][3]。
	    // 可索引对象包括map、slice、array。
	    "index": index,
	    // 返回与其参数的文本表示形式等效的转义JavaScript。
	    "js": JSEscaper,
	    // 返回参数的length。
	    "len": length,
	    // 布尔取反。只能一个参数。
	    "not": not,
	    // 返回第一个不为空的参数或最后一个参数。可以有任意多个参数。
	    // "or x y"等价于"if x then x else y"。
	    "or":      or,
	    "print":   fmt.Sprint,
	    "printf":  fmt.Sprintf,
	    "println": fmt.Sprintln,
	    // 以适合嵌入到网址查询中的形式返回其参数的文本表示的转义值。
	    // 这个函数在html/template中不可用。
	    "urlquery": URLQueryEscaper,
	}
*/
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
