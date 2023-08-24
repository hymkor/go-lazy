//go:build ignore
// +build ignore

package main

import (
	"github.com/hymkor/go-lazy"
)

var s1 = lazy.New(func() string {
	println("s1 initialize")
	return "Foo"
})

func main() {
	println("start")
	println(s1.Value())
	println(s1.Value())
}
