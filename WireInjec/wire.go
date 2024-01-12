//go:build wireinject
// +build wireinject

package main

import (
	"context"
	"github.com/google/wire"
)

/**

var PFoo = wire.NewSet(NewFoo) //将 NewFoo 创建的对象设置为可以被别人依赖

*/

var PFoo = wire.NewSet(NewFoo, NewBar, NewBaz) //将有相互依赖关系的对象组合在一起，减少Injector的生成

func AppInjector(ctx context.Context) (Baz, error) {
	wire.Build(NewFoo, NewBar, NewBaz)
	return Baz{}, nil
}
