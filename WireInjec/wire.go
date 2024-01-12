//go:build wireinject
// +build wireinject

package main

import (
	"context"
	"github.com/google/wire"
)

/**
安装，在本地生成 wire.exe 工具文件：
go install github.com/google/wire/cmd/wire@latest

然后在源码目录下执行 wire.exe ，会自动生成 wire_gen.go 文件
*/

/**

var PFoo = wire.NewSet(NewFoo) //将 NewFoo 创建的对象设置为可以被别人依赖

*/

var PFoo = wire.NewSet(NewFoo, NewBar, NewBaz) //将有相互依赖关系的对象组合在一起，减少Injector的生成

func AppInjector(ctx context.Context) (Baz, error) {
	wire.Build(NewFoo, NewBar, NewBaz)
	return Baz{}, nil
}
