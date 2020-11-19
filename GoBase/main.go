package main

import (
	"errors"
	"fmt"
	"runtime/debug"
	"strings"
)

func main() {

	msg := fmt.Sprint("myage", ":", 30)
	fmt.Println("拼接成字符串：", msg)
	fmt.Println("拼接成字符串2：" + msg)
	fmt.Println("---------------")

	GoLangInitObj()

	demo_string()
}

type UserData struct {
	UserId   int
	UserName string
}

//GO语言变量声明和初始化
func GoLangInitObj() {

	/*
		    make用于内建类型（map、slice 和channel）的内存分配
		    new用于各种类型的内存分配

		    内建函数make(T, args)与new(T)有着不同的功能，make只能创建slice、map和channel，并且返回一个有初始值(非零)的T类型，而不是*T
	*/

	//先声明类型，再赋值
	var num int
	num = 2018
	fmt.Println(num)

	//直接赋值多个
	out, in := 2, 3
	fmt.Println(out, in)

	//使用 make 初始化
	//DataBuf := make([]byte, 10)

	MapData := make(map[string]interface{})

	MapData["my"] = "xxxx"

	fmt.Println(MapData)

	var MapData2 map[string]int = map[string]int{"key": 0}

	fmt.Println(MapData2)

	//创建数组(声明长度)
	var array1 = [5]int{1, 2, 3}
	fmt.Println(array1)

	//创建数组(不声明长度)
	var array2 = [...]int{6, 7, 8}
	fmt.Println(array2)

	//创建数组切片 slice
	var array3 = []int{9, 10, 11, 12}
	fmt.Println(array3)

	//创建数组(声明长度)，并仅初始化其中的部分元素
	var array4 = [5]string{3: "Chris", 4: "Ron"}
	fmt.Println(array4)

	//创建数组(不声明长度)，并仅初始化其中的部分元素，数组的长度将根据初始化的元素确定
	var array5 = [...]string{3: "Tom", 2: "Alice"}
	fmt.Println(array5)

	//创建数组切片，并仅初始化其中的部分元素，数组切片的len将根据初始化的元素确定
	var array6 = []string{4: "Smith", 2: "Alice"}
	fmt.Println(array6)

	//使用 new 初始化
	userData1 := new(UserData) //指针类型

	//使用 {} 直接初始化
	userData2 := UserData{}  //非指针
	userData3 := &UserData{} //指针类型

	//使用slice
	userData4 := []UserData{}

	userData5 := UserData{1, "taven"}

	userData6 := UserData{UserId: 1, UserName: "taven"}

	fmt.Println(userData1, userData2, userData3, userData4, userData5, userData6)

	//制造异常
	err := errors.New("这是一个异常")
	if err != nil {
		//defer, panic, recover
		//panic(err)
	}

	//打印异常信息
	fmt.Println(err.Error())

}

func funcA() error {
	defer func() {
		if p := recover(); p != nil {
			fmt.Printf("panic recover! p: %v", p)
			debug.PrintStack()
		}
	}()
	return funcB()
}

func funcA_2() (err error) {
	defer func() {
		if p := recover(); p != nil {
			fmt.Println("panic recover! p:", p)
			str, ok := p.(string)
			if ok {
				err = errors.New(str)
			} else {
				err = errors.New("panic")
			}
			debug.PrintStack()
		}
	}()
	return funcB()
}

func funcB() error {
	// simulation
	panic("foo")
	return errors.New("success")
}

func test() {
	err := funcA()
	if err == nil {
		fmt.Printf("err is nil\\n")
	} else {
		fmt.Printf("err is %v\\n", err)
	}
}

func demo_string() {

	fmt.Println("查找子串是否在指定的字符串中")
	fmt.Println(" Contains 函数的用法")
	fmt.Println(strings.Contains("seafood", "foo")) //true
	fmt.Println(strings.Contains("seafood", "bar")) //false
	fmt.Println(strings.Contains("seafood", ""))    //true
	fmt.Println(strings.Contains("", ""))           //true 这里要特别注意
	fmt.Println(strings.Contains("我是中国人", "我"))     //true
	fmt.Println("")
	fmt.Println(" ContainsAny 函数的用法")
	fmt.Println(strings.ContainsAny("team", "i"))        // false
	fmt.Println(strings.ContainsAny("failure", "u & i")) // true
	fmt.Println(strings.ContainsAny("foo", ""))          // false
	fmt.Println(strings.ContainsAny("", ""))             // false
	fmt.Println("")
	fmt.Println(" ContainsRune 函数的用法")
	fmt.Println(strings.ContainsRune("我是中国", '我')) // true 注意第二个参数，用的是字符
	fmt.Println("")
	fmt.Println(" Count 函数的用法")
	fmt.Println(strings.Count("cheese", "e")) // 3
	fmt.Println(strings.Count("five", ""))    // before & after each rune result: 5 , 源码中有实现
	fmt.Println("")
	fmt.Println(" EqualFold 函数的用法")
	fmt.Println(strings.EqualFold("Go", "go")) //大小写忽略
	fmt.Println("")
	fmt.Println(" Fields 函数的用法")
	fmt.Println("Fields are: %q", strings.Fields("  foo bar  baz   ")) //["foo" "bar" "baz"] 返回一个列表
	//相当于用函数做为参数，支持匿名函数
	for _, record := range []string{" aaa*1892*122", "aaa\taa\t", "124|939|22"} {
		fmt.Println(strings.FieldsFunc(record, func(ch rune) bool {
			switch {
			case ch > '5':
				return true
			}
			return false
		}))
	}
	fmt.Println("")
	fmt.Println(" HasPrefix 函数的用法")
	fmt.Println(strings.HasPrefix("NLT_abc", "NLT")) //前缀是以NLT开头的
	fmt.Println("")
	fmt.Println(" HasSuffix 函数的用法")
	fmt.Println(strings.HasSuffix("NLT_abc", "abc")) //后缀是以NLT开头的
	fmt.Println("")
	fmt.Println(" Index 函数的用法")
	fmt.Println(strings.Index("NLT_abc", "abc")) // 返回第一个匹配字符的位置，这里是4
	fmt.Println(strings.Index("NLT_abc", "aaa")) // 在存在返回 -1
	fmt.Println(strings.Index("我是中国人", "中"))     // 在存在返回 6
	fmt.Println("")
	fmt.Println(" IndexAny 函数的用法")
	fmt.Println(strings.IndexAny("我是中国人", "中")) // 在存在返回 6
	fmt.Println(strings.IndexAny("我是中国人", "和")) // 在存在返回 -1
	fmt.Println("")
	fmt.Println(" Index 函数的用法")
	fmt.Println(strings.IndexRune("NLT_abc", 'b')) // 返回第一个匹配字符的位置，这里是4
	fmt.Println(strings.IndexRune("NLT_abc", 's')) // 在存在返回 -1
	fmt.Println(strings.IndexRune("我是中国人", '中'))   // 在存在返回 6
	fmt.Println("")
	fmt.Println(" Join 函数的用法")
	s := []string{"foo", "bar", "baz"}
	fmt.Println(strings.Join(s, ", ")) // 返回字符串：foo, bar, baz
	fmt.Println("")
	fmt.Println(" LastIndex 函数的用法")
	fmt.Println(strings.LastIndex("go gopher", "go")) // 3
	fmt.Println("")
	fmt.Println(" LastIndexAny 函数的用法")
	fmt.Println(strings.LastIndexAny("go gopher", "go")) // 4
	fmt.Println(strings.LastIndexAny("我是中国人", "中"))      // 6
	fmt.Println("")
	fmt.Println(" Map 函数的用法")
	rot13 := func(r rune) rune {
		switch {
		case r >= 'A' && r <= 'Z':
			return 'A' + (r-'A'+13)%26
		case r >= 'a' && r <= 'z':
			return 'a' + (r-'a'+13)%26
		}
		return r
	}
	fmt.Println(strings.Map(rot13, "'Twas brillig and the slithy gopher..."))
	fmt.Println("")
	fmt.Println(" Repeat 函数的用法")
	fmt.Println("ba" + strings.Repeat("na", 2)) //banana
	fmt.Println("")
	fmt.Println(" Replace 函数的用法")
	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))
	fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1))
	fmt.Println("")
	fmt.Println(" Split 函数的用法")
	fmt.Printf("%q\n", strings.Split("a,b,c", ","))
	fmt.Printf("%q\n", strings.Split("a man a plan a canal panama", "a "))
	fmt.Printf("%q\n", strings.Split(" xyz ", ""))
	fmt.Printf("%q\n", strings.Split("", "Bernardo O'Higgins"))
	fmt.Println("")
	fmt.Println(" SplitAfter 函数的用法")
	fmt.Printf("%q\n", strings.SplitAfter("/home/m_ta/src", "/")) //["/" "home/" "m_ta/" "src"]
	fmt.Println("")
	fmt.Println(" SplitAfterN 函数的用法")
	fmt.Printf("%q\n", strings.SplitAfterN("/home/m_ta/src", "/", 2)) //["/"

}
