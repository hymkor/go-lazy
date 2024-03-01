//go:build ignore
// +build ignore

package main

import (
	"github.com/hymkor/go-lazy"
)

var counter = 0

var s1 = lazy.Two[string, int]{
	New: func() (string, int) {
		println("s1 initialize")
		counter++
		return "Foo", counter
	},
}

func main() {
	println("start")
	println(s1.Values())
	println(s1.Values())
}
