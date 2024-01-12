package main

import (
	"context"
	"fmt"
)

/**
安装，在本地生成 wire.exe 工具文件：
go install github.com/google/wire/cmd/wire@latest

然后在源码目录下执行 wire.exe ，会自动生成 wire_gen.go 文件
*/

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
