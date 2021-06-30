package main

import (
	"fmt"
	"sync"
)

type Person struct {
}

//方式1
var once sync.Once
var instance *Person
func GetInstance1() *Person {
	once.Do(func(){
		instance = new(Person)
	})
	return instance
}

//方式2
var mutex sync.Mutex
func GetInstance2() *Person {
	mutex.Lock()
	defer mutex.Unlock()

	if instance == nil {
		instance = new(Person)
	}
	return instance
}

//方式3
var instance3 *Person = new(Person)
func GetInstance3() *Person {
	return instance
}

func (_self *Person) Eat() {
	fmt.Println("Person.Eat")
}

func (_self *Person) Talk() {
	fmt.Println("Person.Talk")
}