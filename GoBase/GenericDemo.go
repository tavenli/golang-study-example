package main

import (
	"fmt"
	"strconv"
)

//泛型集合
type List[T any] []T
type HashMap[K string, V any] map[K]V
type MapChan[T any] chan T

//泛型约束
type NumStr interface {
	Num | Str
}
type Num interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64 | ~complex64 | ~complex128
}
type Str interface {
	string
}

type ShowPrice interface {
	String() string
	~int | ~string
}

//带~的约束，只要是同类型约定的interface都可以，不带~的表示严格约束，必须是同一个泛型类型
type ShowPrice2 interface {
	String() string
	int | string
}

//Cannot use comparable in union
/*

type ShowPrice3 interface {
	int | string | comparable
}

*/

func GenericDemo() {
	//Go 1.18 开始正式支持泛型

	PrintSlice[int]([]int{5, 6, 7, 8, 100})
	PrintSlice[float64]([]float64{1.1, 2.2, 5.5})
	PrintSlice[string]([]string{"张三", "李四", "王五", "曾麻子"})

	//省略显示类型
	PrintSlice([]int64{20, 30, 33, 40, 50})

	v := List[int]{58, 1881}
	PrintSlice(v)
	v2 := List[string]{"张三", "李四", "王五"}
	PrintSlice(v2)

	m1 := HashMap[string, int]{"key": 1}
	m1["key"] = 2

	m2 := HashMap[string, string]{"key": "value"}
	m2["key"] = "kiss my ass"
	fmt.Println(m1, m2)

	c1 := make(MapChan[int], 10)
	c1 <- 1
	c1 <- 2

	c2 := make(MapChan[string], 10)
	c2 <- "hello"
	c2 <- "world"

	fmt.Println(<-c1, <-c2)
	fmt.Println(<-c1, <-c2)

	//泛型约束使用例子
	fmt.Println(AutoAdd(3, 4))
	fmt.Println(AutoAdd("Hello", "World"))

	//
	fmt.Println(ShowPriceList([]Price{1, 2}))
	fmt.Println(ShowPriceList([]Price2{"a", "b"}))

	//下面的使用方式编译器会报错
	//Type does not implement constraint 'ShowPrice2' because type is not included in type set ('int', 'string')
	//fmt.Println(ShowPriceList2([]Price{1, 2}))

	fmt.Println(FindIndex([]int{1, 2, 3, 4, 5, 6}, 5))
	fmt.Println(FindIndex([]string{"张三", "李四", "王五", "曾麻子"}, "李四"))

}

func PrintSlice[T any](s []T) {
	for _, v := range s {
		fmt.Printf("%v\n", v)
	}
}
func AutoAdd[T NumStr](a, b T) T {
	return a + b
}

func ShowPriceList[T ShowPrice](s []T) (ret []string) {
	for _, v := range s {
		ret = append(ret, v.String())
	}
	return
}

func ShowPriceList2[T ShowPrice2](s []T) (ret []string) {
	for _, v := range s {
		ret = append(ret, v.String())
	}
	return
}

type Price int

func (i Price) String() string {
	return strconv.Itoa(int(i))
}

type Price2 string

func (i Price2) String() string {
	return string(i)
}

func FindIndex[T comparable](a []T, v T) int {
	for i, e := range a {
		if e == v {
			return i
		}
	}
	return -1
}
