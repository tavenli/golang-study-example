package main

import "fmt"

func CallbackFunc_main() {

	CallbackFunc_demo1()

	CallbackFunc_demo2()
}

type Callback1 func(int) string

func Callback1_Impl(age int) string {
	return fmt.Sprint("i am ", age)
}

func sayHello1(msg string, callback Callback1) {

	fmt.Println(msg)
	fmt.Println(callback(1))
}

func sayHello2(msg string, callback func(int) string) {

	fmt.Println(msg)
	fmt.Println(callback(2))
}

func CallbackFunc_demo1() {
	sayHello1("No 1", Callback1_Impl)
}

func CallbackFunc_demo2() {

	sayHello2("No 2", func(age int) string {
		return fmt.Sprint("i am ", age)
	})
}
