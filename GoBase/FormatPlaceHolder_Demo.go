package main

import (
	"fmt"
)

type BigBoy struct {
	Name string
	Age  int
}

func FormatPlaceHolder_main() {
	//
	fmt.Println("格式化 占位符")

	//错误的使用方式，单元测试会报错：call has possible formatting directive %v
	//fmt.Println("示范：%v", "错误的使用方式")
	fmt.Printf("示范：%v \n", "正确的使用方式")

	bigBoy := &BigBoy{Name: "TavenLi", Age: 18}

	fmt.Printf("%v", bigBoy)
	//输出：&{TavenLi 18}
	fmt.Println("")

	//打印结构体时，会添加字段名
	fmt.Printf("%+v", bigBoy)
	//输出：&{Name:TavenLi Age:18}
	fmt.Println("")

	fmt.Printf("%#v", bigBoy)
	//输出：&main.BigBoy{Name:"TavenLi", Age:18}
	fmt.Println("")

	//字面上的百分号，并非值的占位符
	fmt.Printf("%%", bigBoy)
	//输出：%%!(EXTRA *main.BigBoy=&{TavenLi 18})
	fmt.Println("")

	//指针、指针的指针、指针的值
	fmt.Printf("\n对象指针地址：%p", bigBoy)
	fmt.Printf("\n对象指针地址：%p", &bigBoy)
	fmt.Printf("\n对象指针地址：%p", *bigBoy)

	fmt.Printf("\n这是布尔值：%t", 1 == 1)
	fmt.Printf("\n二进制表示：%b", 129)
	fmt.Printf("\nUnicode对应字符：%c", 0x4E2D)
	fmt.Printf("\n十进制表示：%d", 0x12)
	fmt.Printf("\n八进制表示：%o", 10)
	fmt.Printf("\n单引号围绕：%q", 0x4E2D)
	fmt.Printf("\n十六进制表示，字母形式为小写 a-f：%x", 13)
	fmt.Printf("\n十六进制表示，字母形式为大写 A-F：%X", 13)
	fmt.Printf("\nUnicode格式：%U", 0x4E2D)
	fmt.Printf("\nASCII编码：%+q", "远哥")
	fmt.Printf("\nUnicode编码：%#U", '远')

	fmt.Printf("\n科学计数法：%e", 10.2)
	fmt.Printf("\n科学计数法：%E", 10.2)
	fmt.Printf("\n有小数点而无指数：%f", 10.2)
	fmt.Printf("\n根据情况选择无末尾的0 小写：%g", 10.20)
	fmt.Printf("\n根据情况选择无末尾的0 小写：%G", 10.20+2i)

	fmt.Printf("\n输出对应字符的值：%s", []byte("犀利的远哥"))
	fmt.Printf("\n加双引号包裹：%q", "犀利的远哥")
	fmt.Printf("\n转十六进制，小写：%x", "TavenLi")
	fmt.Printf("\n转十六进制，大写：%X", "TavenLi")

}
