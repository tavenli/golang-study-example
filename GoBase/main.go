package main

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"math"
	"math/rand"
	"os"
	"reflect"
	"regexp"
	"runtime"
	"runtime/debug"
	"sort"
	"strconv"
	"strings"
	"time"
)

var appname = flag.String("appname", "demo", "app name")
var bShowVersion = flag.Bool("version", false, "show version")
var port = flag.Int("version", 8000, "server port")

func main() {
	// flag.Parse 要在第一行执行
	flag.Parse()

	//设置本应用最大使用CPU的核心个数，默认使用所有CPU
	runtime.GOMAXPROCS(runtime.NumCPU())

	note := `
		这里是定义一段可以换行的字符串内容，
		是不是很方便？
	`

	msg := fmt.Sprint("myage", ":", 30)
	fmt.Println("拼接成字符串：", msg)
	fmt.Println("拼接成字符串2：" + note)
	fmt.Println("---------------")

	ShowSysInf()

	GoLangInitObj()

	demo_string()

	fmt.Println("==========================")
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

	//字符串转byte
	b := []byte("abc1def1")
	fmt.Println(b)

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

	/*
		//string到int
		int, err := strconv.Atoi(string)
		//string到int64
		int64, err := strconv.ParseInt(string, 10, 64)
		//int到string
		string := strconv.Itoa(int)
		//int64到string
		string := strconv.FormatInt(int64, 10)
	*/

	//
	var e interface{}
	e = 10

	switch v := e.(type) {
	case int:
		fmt.Println("整型", v)
		break
	case string:
		fmt.Println("字符串", v)
		break

	}

}

func demo_collection() {
	//集合对象的操作

	var users []*UserData

	user := new(UserData)
	user.UserId = 1
	user.UserName = "Taven"
	users = append(users, user)

	user2 := new(UserData)
	user2.UserId = 2
	user2.UserName = "Taven2"
	users = append(users, user2)

	//或者一次性添加多个
	users = append(users, user, user2)

	fmt.Println(len(users), cap(users))

	users = slice_remove(users, 1)

	//-----------------------------------
	map1 := make(map[string]int)

	map1["one"] = 1
	map1["two"] = 2
	map1["three"] = 3
	map1["four"] = 4

	//取值
	mvalue := map1["one"]
	mvalue, contain := map1["one"]
	fmt.Println(mvalue, contain)

	//判断key是否存在
	if _, ok := map1["one"]; ok {
		//存在
	}

	fmt.Println(map1, len(map1))

	//删除key
	delete(map1, "two")
	fmt.Println(map1, len(map1))

	/*

		for k, v := range itemTagMap {
			fmt.Printf("k=%v, v=%v\n", k, v)
			if v == "person" {
				p++
			}

			if v == "team" {
				t++
			}

		}

		var rawMap map[string]interface{}

		var itemTagMap = make(map[int64]string)
		for i := 0; i < len(itemIds); i++ {
			voteItem := services.GetVoteItem(itemIds[i])
			itemTagMap[voteItem.Id] = voteItem.ItemTag
		}

		//遍历map
		for k, v := range itemTagMap {

		}

		var rateResults []*models.RateResult

		rateResults = append(rateResults, rateResult)

	*/

}

func slice_remove(s []*UserData, i int) []*UserData {
	return append(s[:i], s[i+1:]...)
}

func slice_copy() {
	//使用copy复制切片之前，要保证目标切片有足够的大小，注意是大小，而不是容量

	var sa = make([]string, 0)
	for i := 0; i < 10; i++ {
		sa = append(sa, fmt.Sprintf("%v", i))

	}
	var da = make([]string, 0, 10)
	var cc = 0
	cc = copy(da, sa)
	fmt.Printf("copy to da(len=%d)\t%v\n", len(da), da)
	da = make([]string, 5)
	cc = copy(da, sa)
	fmt.Printf("copy to da(len=%d)\tcopied=%d\t%v\n", len(da), cc, da)
	da = make([]string, 10)
	cc = copy(da, sa)
	fmt.Printf("copy to da(len=%d)\tcopied=%d\t%v\n", len(da), cc, da)

}

func IntArrayFind(slice []int, value int) int {
	for p, v := range slice {
		if v == value {
			return p
		}
	}
	return -1
}

func IntArrayContain(slice []int, value int) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func Int64ArrayFind(slice []int64, value int64) int {
	for p, v := range slice {
		if v == value {
			return p
		}
	}
	return -1
}

func Int64ArrayContain(slice []int64, value int64) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func StrArrayContain(slice []string, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

//golang学习之指针、内存分配
func pointer_test() {
	//空指针，输出为nil
	var p *int
	fmt.Printf("p: %v\n", p)
	//p: <nil>

	//指向局部变量，变量值初始为0
	var i int
	p = &i
	fmt.Printf("p: %v,%v\n", p, *p)
	//p: 0xc04203c440,0

	i = 100
	fmt.Println("i:", i)   //i: 100
	fmt.Println("&i:", &i) //&i: 0xc04203c440

	//通过指针修改变量数值
	*p = 8
	fmt.Printf("p: %v,%v\n", p, *p)
	//p: 0xc04200a4f0,8

	//数组的初始化及输出
	m := [3]int{3, 4, 5}
	fmt.Printf("m:%v--%v,%v,%v\n", m, m[0], m[1], m[2])
	//m:[3 4 5]--3,4,5

	//指针数组的初始化及输出
	//j, k, l := 3, 4, 5
	//x := [3]*int{&j, &k, &l}
	x := [3]*int{&m[0], &m[1], &m[2]}
	fmt.Printf("x:%v,%v,%v\n", x[0], x[1], x[2])     //x:0xc04205c860,0xc04205c868,0xc04205c870
	fmt.Printf("*x:%v,%v,%v\n", *x[0], *x[1], *x[2]) //*x:3,4,5

	var n [3]*int
	n = x
	fmt.Printf("n:%v,%v,%v\n", n[0], n[1], n[2])
	//指向数组的指针，也即二级指针的使用
	y := []*[3]*int{&x}
	fmt.Printf("y:%v,%v\n", y, y[0])
	fmt.Printf("*y[0]:%v\n", *y[0])
	fmt.Printf("*y[][]:%v,%v,%v\n", *y[0][0], *y[0][1], *y[0][2])
	/*
	   *X            ->    v
	   [3]X        ->    [v0][v1][v2]
	   []X            ->    [v0]...[vi]
	   [3]*X        ->    [p0][p1][p2]
	                     |   |   |
	                     j   k   l
	   []*[3]*X    ->    [n0]...[ni]
	                     | ... |
	                    [p0]->j
	                    [p1]->k
	                    [p2]->l
	*/
}

type Student struct {
	name  string
	id    int
	score uint
}

func memery_test() {
	//new分配出来的数据是指针形式
	p := new(Student)
	p.name = "China"
	p.id = 63333
	p.score = 99
	fmt.Println(*p)
	//var定义的变量是数值形式
	var st Student
	st.name = "Chinese"
	st.id = 666333
	st.score = 100
	fmt.Println(st)
	//make分配slice、map和channel的空间，并且返回的不是指针
	var ptr *[]Student
	fmt.Println(ptr)     //ptr == nil
	ptr = new([]Student) //指向一个空的slice
	fmt.Println(ptr)
	*ptr = make([]Student, 3, 100)
	fmt.Println(ptr)
	stu := []Student{{"China", 3333, 66}, {"Chinese", 4444, 77}, {"Chince", 5555, 88}}
	fmt.Println(stu)
}

//	显示当前系统基本信息
func ShowSysInf() {

	fmt.Println("★★★★★★★★★★★★★★★★★★★★★★★★")
	fmt.Println("runtime.Version --->", runtime.Version()) //GO的版本
	fmt.Println("runtime.NumCPU --->", runtime.NumCPU())   //CPU核数
	fmt.Println("runtime.GOOS --->", runtime.GOOS)         //运行GO的OS操作系统
	fmt.Println("runtime.GOARCH --->", runtime.GOARCH)     //CPU架构
	fmt.Println("runtime.Version --->", runtime.Version()) //当前GO语言版本
	fmt.Println("time --->", time.Now())                   //系统当前时间
	fmt.Println("★★★★★★★★★★★★★★★★★★★★★★★★")

	//var memStats runtime.MemStats
	//runtime.ReadMemStats(&memStats)
	//fmt.Println("runtime.memStats --->", memStats)

	//获取全部的环境变量
	// data := os.Environ()
	// for _, val := range data {
	//     fmt.Println(val)
	// }

}

//	go不支持三元表达式，可以使用自定义的函数实现
//	例如：max := utils.If(x > y, x, y).(int)
func If(condition bool, trueVal, falseVal interface{}) interface{} {

	if condition {
		return trueVal
	}
	return falseVal
}

/*
	交换int数据：a, b := utils.Swap(2, 9)
	交换字符串数据：A, B := utils.Swap("Li", "Chen")
*/
func Swap(x, y interface{}) (interface{}, interface{}) {
	return y, x
}

//	设置环境变量
func SetEnv(key, value string) error {

	return os.Setenv(key, value)
}

//	取环境变量的值
func GetEnv(key string) string {

	return os.Getenv(key)
}

//取进程ID
func GetPid() int {
	return os.Getpid()
}

func KillByPid(pid int) {
	p, _ := os.FindProcess(pid)
	fmt.Println("KillByPid", p)
	p.Kill()
}

func StartProcessDemo() {
	//例子演示
	attr := &os.ProcAttr{
		Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
	}
	p, _ := os.StartProcess("xxx.exe", []string{"xxx", "1.txt"}, attr)
	p.Release()
	time.Sleep(10000)
	p.Signal(os.Kill)
	os.Exit(10)
}

func ToJson(obj interface{}) ([]byte, error) {
	return json.Marshal(obj)
}

func FromJson(data []byte, t interface{}) error {
	return json.Unmarshal(data, t)
}

func ShowObjAllProp(obj interface{}) {
	value_method := reflect.ValueOf(obj)
	obj_type := value_method.Type()

	fmt.Printf("输出对象的属性和方法\t%v\n", obj)

	fmt.Println("\tMethods...")

	for i := 0; i < value_method.NumMethod(); i++ {
		fmt.Printf("\t%d\t%s\n", i, obj_type.Method(i).Name)
	}

	value_element := reflect.ValueOf(obj).Elem()
	obj_element := value_element.Type()

	fmt.Println("\tFields...")
	for i := 0; i < value_element.NumField(); i++ {
		fmt.Printf("\t%d\t%s\n", i, obj_element.Field(i).Name)

	}
}

var (
	CryptoNumStr  = "0123456789"
	CryptoCharStr = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	CryptoSpecStr = "+=-@#~,.[]()!%^*$"
)

//生成32位md5字串
func GetMd5(input string) string {
	hash := md5.New()
	hash.Write([]byte(input))
	return hex.EncodeToString(hash.Sum(nil))
}

func GetSaltMD5(input, salt string) string {
	hash := md5.New()
	//salt = "salt123456" //盐值
	io.WriteString(hash, input+salt)
	result := fmt.Sprintf("%x", hash.Sum(nil))
	return result
}

func GenRandomSalt(length int) string {
	var result []byte = make([]byte, length, length)

	//随机种子
	rand.Seed(time.Now().UnixNano())

	var seedStr string
	//纯数字
	//seedStr = CryptoNumStr
	//纯字母
	//seedStr = CryptoCharStr
	//数字加字母组合
	seedStr = fmt.Sprint(CryptoNumStr, CryptoCharStr)
	//全家桶组合
	//seedStr = fmt.Sprint(CryptoNumStr, CryptoCharStr, CryptoSpecStr)

	for i := 0; i < length; i++ {
		index := rand.Intn(len(seedStr))
		result[i] = seedStr[index]
	}

	return string(result)
}

//  GO的诞辰
const timeLayout = "2006-01-02 15:04:05"

//  取当前系统时间
func GetTimeNow() time.Time {
	return time.Now()
}

func GetTime(timeStr string) time.Time {
	toTime, _ := ToTime(timeStr)
	return toTime
}

func JavaLongTime(javaLong int64) time.Time {
	//1492566520958	-> 2017-04-19 09:48:40
	//fmt.Println(time.Unix(1492566520958/1000, 0))
	//fmt.Println(time.Unix(0, 1492566520958*1000000))
	return time.Unix(0, javaLong*1000000)
}

func ToTime(timeStr string) (time.Time, error) {
	loc, _ := time.LoadLocation("Local")
	toTime, err := time.ParseInLocation(timeLayout, timeStr, loc)
	//toTime, err := time.Parse(timeLayout, timeStr)
	return toTime, err

}

func ToTimeByFm(timeStr string, format string) (time.Time, error) {
	loc, _ := time.LoadLocation("Local")
	toTime, err := time.ParseInLocation(format, timeStr, loc)
	//toTime, err := time.Parse(timeLayout, timeStr)
	return toTime, err

}

//要想格式化为：yyyyMMddHHmmss
//则 format = "20060102150405"
//要想格式化为：yyyy-MM-dd HH:mm:ss
//则 format = "2006-01-02 15:04:05"
//要想格式化为：yyyy-MM-dd
//则 format = "2006-01-02"
func FormatTimeByFm(t time.Time, format string) string {

	return t.Format(format)
}

func GetCurrentTime() string {
	return FormatTime(time.Now())
}

func GetCurrentDay() string {
	return FormatTimeByFm(time.Now(), "2006-01-02")
}

func FormatTime(t time.Time) string {
	//
	return FormatTimeByFm(t, "2006-01-02 15:04:05")
}

func FormatTimeToNum(t time.Time) string {
	//
	return FormatTimeByFm(t, "20060102150405")
}

//  在当前时间之前
func IsBeforeNow(t time.Time) (result bool) {
	result = false
	if &t != nil && t.Before(time.Now()) {
		result = true
	}
	return
}

//  在当前时间之后
func IsAfterNow(t time.Time) (result bool) {
	result = false
	if &t != nil && t.After(time.Now()) {
		result = true
	}
	return
}

func SubDateTime(firstTime time.Time, secondTime time.Time) (result time.Duration) {
	result = time.Duration(0)
	if &firstTime != nil && &secondTime != nil {
		result = secondTime.Sub(firstTime)
	}
	return
}

func DifferDays(firstTime time.Time, secondTime time.Time) int64 {
	result := SubDateTime(firstTime, secondTime).Hours()
	return int64(math.Abs(result) / 24)
}

func DifferHour(firstTime time.Time, secondTime time.Time) int64 {
	result := SubDateTime(firstTime, secondTime).Hours()
	//return int64(result) 两个时间的先后顺序不一样，可能出现负数
	return int64(math.Abs(result))
}

func DifferMin(firstTime time.Time, secondTime time.Time) int64 {
	result := SubDateTime(firstTime, secondTime).Minutes()
	return int64(math.Abs(result))
}

func DifferSec(firstTime time.Time, secondTime time.Time) int64 {
	result := SubDateTime(firstTime, secondTime).Seconds()
	return int64(math.Abs(result))
}

//  24小时前的时间
func Before24h() time.Time {
	t, _ := time.ParseDuration("-24h")
	return time.Now().Add(t)
}

func AddSecs(_time time.Time, secs int64) time.Time {
	t, _ := time.ParseDuration("1s")
	return time.Now().Add(t * time.Duration(secs))
}

/*
   增加10分钟：utils.AddMins(time.Now(), 10)
   减少5分钟：utils.AddMins(time.Now(), -5)
*/
func AddMins(_time time.Time, mins int64) time.Time {
	t, _ := time.ParseDuration("1m")
	return time.Now().Add(t * time.Duration(mins))
}

func AddHours(_time time.Time, hours int64) time.Time {
	t, _ := time.ParseDuration("1h")
	return time.Now().Add(t * time.Duration(hours))
}

func AddDays(_time time.Time, days int) time.Time {
	return _time.AddDate(0, 0, days)
}

func AddMonths(_time time.Time, months int) time.Time {
	return _time.AddDate(0, months, 0)
}

func GetBeginTime(_time time.Time) time.Time {
	//2017-06-28 00:00:00 +0800 CST
	return GetBeginTimeByLoc(_time, time.Local)
	//return GetBeginTimeByLoc(_time, time.UTC)

}

func GetEndTime(_time time.Time) time.Time {
	//2017-06-28 23:59:59.999999999 +0800 CST
	return GetEndTimeByLoc(_time, time.Local)
}

func GetBeginTimeByLoc(_time time.Time, loc *time.Location) time.Time {
	year, month, day := _time.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, loc)

}

func GetEndTimeByLoc(_time time.Time, loc *time.Location) time.Time {
	year, month, day := _time.Date()
	return time.Date(year, month, day, 23, 59, 59, 999999999, loc)
}

// 一行代码计算代码执行时间
// defer utils.TimeCost(time.Now())
func TimeCost(start time.Time) {
	terminal := time.Since(start)
	fmt.Println("TimeCost：", terminal)
}

func AbsInt(num float64) int {
	//result := math.Abs(float64(num))
	result := math.Abs(num)
	return int(result)
}

func AbsInt64(num float64) int64 {
	result := math.Abs(num)
	return int64(result)
}

func CeilInt(num float64) int {
	result := math.Ceil(num)
	return int(result)
}

func CeilInt64(num float64) int64 {
	//CeilInt64(5.9) = 6
	//CeilInt64(5.3) = 6
	//CeilInt64(5) = 5
	result := math.Ceil(num)
	return int64(result)
}

func Float64ToInt64(num float64) int64 {
	return int64(num)
}

func Float64TryToInt64(num interface{}) int64 {
	return int64(num.(float64))
}

//	返回最大值
func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//	返回最小值
func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Pages(total, psize int64) int64 {

	pages := float64(total) / float64(psize)
	result := math.Ceil(pages)
	return int64(result)
}

func ToInt(str string) int {
	_num, _ := strconv.Atoi(str)
	return _num
}

func StrToIntArray(str string) []int {
	var arr []int
	for _, _id := range Split(str, ",") {
		id := ToInt(_id)
		arr = append(arr, id)
	}

	return arr
}

func StrArrayToIntArray(strArray []string) []int {
	var arr []int
	for _, _id := range strArray {
		id := ToInt(_id)
		arr = append(arr, id)
	}

	return arr
}

func ToInt64(str string) int64 {
	_num, _ := strconv.ParseInt(str, 10, 64)
	return _num
}

func ToInteger(str string) (int, error) {
	_num, _err := strconv.Atoi(str)
	return _num, _err
}

func ToLong(str string) (int64, error) {
	_num, _err := strconv.ParseInt(str, 10, 64)
	return _num, _err
}

func ToFloat64(str string) (float64, error) {
	_num, _err := strconv.ParseFloat(str, 64)
	return _num, _err
}

func BinaryToInt(str string) (int64, error) {
	_num, _err := strconv.ParseInt(str, 2, 64)
	return _num, _err
}

func IntToBinary(num int64) string {
	bin := strconv.FormatInt(num, 2)
	return bin
}

func IsBinaryOverInt(binStr string, number int64) bool {
	_num, _ := strconv.ParseInt(binStr, 2, 64)
	return (_num & number) == number
}

func IsBinNumOverInt(binNum int64, number int64) bool {

	return (binNum & number) == number
}

func ToStr(_num int) string {
	return strconv.Itoa(_num)
}

func FormatInt(_num int) string {
	return strconv.FormatInt(int64(_num), 10)
}

func FormatInt64(_num int64) string {
	return strconv.FormatInt(_num, 10)
}

func FormatFloat64(_num float64) string {
	return strconv.FormatFloat(_num, 'f', 2, 64)
}

func IsEmpty(str string) bool {

	return Len(str) <= 0
}

func IsNotEmpty(str string) bool {

	return !IsEmpty(str)
}

func Replace(str string, find string, to string) string {

	return strings.Replace(str, find, to, 1)
}

func ReplaceAll(str string, find string, to string) string {

	return strings.Replace(str, find, to, -1)
}

func Split(str string, spChar string) []string {

	return strings.Split(str, spChar)
}

func Contains(str string, find string) bool {

	return strings.Contains(str, find)
}

func TrimSpace(str string) string {

	return strings.TrimSpace(str)
}

func TrimPrefix(str string, find string) string {

	return strings.TrimPrefix(str, find)
}

//	strings.HasPrefix("ABC_xyz", "ABC")
func StartsWith(str string, find string) bool {

	return strings.HasPrefix(str, find)
}

//	strings.HasSuffix("ABC_xyz", "xyz")
func EndsWith(str string, find string) bool {

	return strings.HasSuffix(str, find)
}

//  strings.Count("cheese", "e") = 3
func Count(str string, find string) int {

	return strings.Count(str, find)
}

//  返回第一个匹配字符的位置，返回-1为未找到
//  strings.Index("ABC_xyz", "xyz") = 4
//  strings.Index("ABC_xyz", "B") = 1
func Index(str string, find string) int {

	return strings.Index(str, find)
}

//strings.Join(arrays, ",") = "foo, bar, bas"
func Join(strs []string, spChar string) string {

	return strings.Join(strs, spChar)
}

func IntArrayJoin(array []int, spChar string) string {
	var buffer bytes.Buffer
	for _i, _id := range array {
		if _i == len(array)-1 {
			buffer.WriteString(ToStr(_id))
		} else {
			buffer.WriteString(fmt.Sprint(_id, spChar))
		}

	}

	return buffer.String()
}

//  字母转为小写
//  strings.ToLower("Love GoLang") = "love golang"
func ToLower(str string) string {

	return strings.ToLower(str)
}

//  字母转为大写
//  strings.ToTitle("love 中国") = "LOVE 中国"
func ToUpper(str string) string {
	return strings.ToUpper(str)
	//return strings.ToTitle(str)
}

func Len(str string) int {

	return len(str)
}

func Print(str string) {
	//var show = fmt.Println
	//show(str)
	fmt.Println(str)
}

func FilterByRegex(expr, input, placeTo string) string {
	regx, _ := regexp.Compile(expr)
	return regx.ReplaceAllString(input, placeTo)
}

func FilterStyle(input string) string {
	//regx, _ := regexp.Compile("<style((?:.|\\n)*?)</style>")
	regx, _ := regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	return regx.ReplaceAllString(input, "")
}

func FilterScript(input string) string {
	//regx, _ := regexp.Compile("<script((?:.|\\n)*?)</script>")
	regx, _ := regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	return regx.ReplaceAllString(input, "")
}

func FilterHtml(input string) string {
	regx, _ := regexp.Compile("<.+?>")
	//regx, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	return regx.ReplaceAllString(input, "")
}

func FilterA(input string) string {

	regx, _ := regexp.Compile("<.?a(.|\n)*?>")
	return regx.ReplaceAllString(input, "")
}

func FilterImage(input string) string {

	regx, _ := regexp.Compile("<img(.|\\n)*?>")
	return regx.ReplaceAllString(input, "")
}

func FilterSpecialChar(input string) string {

	regx, _ := regexp.Compile("[+=|{}':;',]")
	return regx.ReplaceAllString(input, "")
}

func FilterUrlPrefix(input string) string {

	regx, _ := regexp.Compile("\\w+://")
	return regx.ReplaceAllString(input, "")
}

func IsNumber(input string) bool {

	match, _ := regexp.MatchString("^\\d+$", input)
	return match
}

func IsIP(input string) bool {

	match, _ := regexp.MatchString("^((2[0-4]\\d|25[0-5]|[01]?\\d\\d?)\\.){3}(2[0-4]\\d|25[0-5]|[01]?\\d\\d?)$", input)
	return match
}

func IsEMail(input string) bool {

	match, _ := regexp.MatchString("^([a-z0-9A-Z]+[-|\\.]?)+[a-z0-9A-Z]@([a-z0-9A-Z]+(-[a-z0-9A-Z]+)?\\.)+[a-zA-Z]{2,}$", input)
	return match
}

//高效拼接字符串
func LinkStrs(inputs ...string) string {
	var buf bytes.Buffer
	for _, v := range inputs {
		buf.WriteString(v)
	}
	return buf.String()
}

func LinkInputs(inputs ...interface{}) string {
	var buf bytes.Buffer
	for _, v := range inputs {
		switch t := v.(type) {
		case string:
			buf.WriteString(t)
		default:
			buf.WriteString(fmt.Sprint(t))

		}
	}
	return buf.String()
}

// func LinkInputs(inputs ...interface{}) string {
// 	var buf bytes.Buffer
// 	for _, v := range inputs {
// 		switch t := v.(type) {
// 		case string:
// 			buf.WriteString(t)
// 		//case int, int64:
// 		case int64:
// 			buf.WriteString(FormatInt64(t))
// 		case int:
// 			buf.WriteString(FormatInt(t))
// 		case float64:
// 			buf.WriteString(FormatFloat64(t))
// 		default:
// 			buf.WriteString(fmt.Sprint(t))

// 		}

// 		fmt.Println("v:", v)

// 	}
// 	return buf.String()
// }

func demo_order_sort() {
	//排序
	intList := []int{2, 4, 3, 5, 7, 6, 9, 8, 1, 0}
	float8List := []float64{4.2, 5.9, 12.3, 10.0, 50.4, 99.9, 31.4, 27.81828, 3.14}
	// float4List := [] float32 {4.2, 5.9, 12.3, 10.0, 50.4, 99.9, 31.4, 27.81828, 3.14}    // no function : sort.Float32s
	stringList := []string{"a", "c", "b", "d", "f", "i", "z", "x", "w", "y"}

	sort.Ints(intList)
	sort.Float64s(float8List)
	sort.Strings(stringList)

	fmt.Printf("%v\n%v\n%v\n", intList, float8List, stringList)

}

func demo_order_sort2() {
	//反向排序
	intList := []int{2, 4, 3, 5, 7, 6, 9, 8, 1, 0}
	float8List := []float64{4.2, 5.9, 12.3, 10.0, 50.4, 99.9, 31.4, 27.81828, 3.14}
	stringList := []string{"a", "c", "b", "d", "f", "i", "z", "x", "w", "y"}

	sort.Sort(sort.Reverse(sort.IntSlice(intList)))
	sort.Sort(sort.Reverse(sort.Float64Slice(float8List)))
	sort.Sort(sort.Reverse(sort.StringSlice(stringList)))

	fmt.Printf("%v\n%v\n%v\n", intList, float8List, stringList)
}
