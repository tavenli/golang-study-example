package main

import (
	"context"
	"errors"
)

// ProvideBaz returns a value if Bar is not zero.
func NewBaz(ctx context.Context, bar Bar) (Baz, error) {
	if bar.X == 0 {
		return Baz{}, errors.New("cannot provide baz when bar is zero")
	}
	return Baz{X: bar.X}, nil
}

// ProvideBar returns a Bar: a negative Foo.
func NewBar(foo Foo) Bar {
	return Bar{X: -foo.X}
}

// ProvideFoo returns a Foo.
func NewFoo() Foo {
	return Foo{X: 42}
}
