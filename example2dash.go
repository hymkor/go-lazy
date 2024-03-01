//go:build ignore

package main

import (
	"github.com/hymkor/go-lazy"
)

var s1 = (&lazy.Of[string]{
	New: func() string {
		println("s1 initialize")
		return "Foo"
	},
}).Value

func main() {
	println("start")
	println(s1())
	println(s1())
}
