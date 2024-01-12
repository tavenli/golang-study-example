package main

import (
	"context"
	"fmt"
)

func main() {
	appInst, err := AppInjector(context.Background())

	fmt.Println(err)
	fmt.Println(appInst)
}

type Foo struct {
	X int
}

type Bar struct {
	X int
}

type Baz struct {
	X int
}
