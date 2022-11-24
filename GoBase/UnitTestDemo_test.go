package main

import (
	"fmt"
	"testing"
)

/*
	单元测试
	文件命名规范：xxx_test.go
	常用的结构体：testing.M、testing.T、testing.B、testing.PB
*/

func TestMain(m *testing.M) {
	fmt.Println("TestMain 首先执行，可以不定义该方法，如果要定义，则方法名必须是 TestMain(t *testing.M)")
	m.Run()
}

func TestCode1(t *testing.T) {
	//Go语言官方库不提供 assert 断言，而是采用 Log、Fail、FailNow 方法来决定测试的成功和失败
	t.Log("success")
	fmt.Println("--- i am TestCode1")
}

func TestCode2(t *testing.T) {
	//表示测试失败，但是后续代码可执行
	t.Fail()

	fmt.Println("--- i am TestCode2")
}

func TestCode3(t *testing.T) {
	//表示 当前函数立即停止，后续代码不再执行
	t.FailNow()

	fmt.Println("--- i am TestCode3")
}

func TestCode4(t *testing.T) {

	fmt.Println("--- i am TestCode4")
}

func TestCode5(t *testing.T) {
	//打印失败错误日志之后立即终止当前测试函数的执行并宣告测试失败
	t.Fatal("TestCode5 fail")

	fmt.Println("--- i am TestCode5")
}

func TestCode6(t *testing.T) {
	fmt.Println("--- i am TestCode6")
}

func DemoCode7(t *testing.T) {
	//方法名须以"Test"打头，并且形参为 (t *testing.T)，否则不会在单元测试中自动执行
	fmt.Println("--- i am DemoCode7")
}

func TestCode8(t *testing.T) {
	t.Run("我运行另外一个函数", DemoFunc)
	fmt.Println("--- i am TestCode8")
}

func DemoFunc(t *testing.T) {
	//设置断点，可以看到代码会执行到这里
	t.Log("DemoFunc")
	fmt.Println("--- i am DemoFunc")
}
