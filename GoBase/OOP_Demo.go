package main

import "fmt"

/*

Go也可以支持面向对象编程

定义类，定义接口，实现多态等

*/

func OopDemo() {

	var superMan IPerson

	tavenli := new(Boy)
	tavenli.JJLen = 30

	sexMM := new(Girl)
	sexMM.BBDeep = 9999

	//先变男人
	superMan = tavenli
	superMan.Speak()

	//再变女人
	superMan = sexMM
	superMan.Speak()

	//superMan.(Girl).Speak()

}

type IPerson interface {
	Speak()
	Eat(int) bool
	Sleep()
}

type Boy struct {
	JJLen int
}

type Girl struct {
	BBDeep int
}

func (_self *Boy) Speak() {
	fmt.Println("fucking you everyday")
}

func (_self *Boy) Eat(time int) bool {
	return false
}

func (_self *Boy) Sleep() {
	fmt.Println("Dream of sexy girls")
}

func (_self *Girl) Speak() {
	fmt.Println("who loves me?")
}

func (_self *Girl) Eat(time int) bool {
	return true
}

func (_self *Girl) Sleep() {
	fmt.Println("Dream of a lot of money")
}
