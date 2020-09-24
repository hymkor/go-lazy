// +build +ignore

package main

import (
	"github.com/zetamatta/go-lazy"
)

var s1 = lazy.OfString{
	New: func() string {
		println("s1 initialize")
		return "Foo"
	},
}

var i1 = lazy.OfInt{
	New: func() int {
		println("i1 initialize")
		return 1
	},
}

var b1 = lazy.OfBool{
	New: func() bool {
		println("b1 initialize")
		return true
	},
}

var a1 = lazy.OfAny{
	New: func() interface{} {
		println("a1 initialize")
		return "a"
	},
}

func main() {
	println("*** First ***")
	println(s1.Value())
	println(i1.Value())
	println(b1.Value())
	println(a1.Value())
	println("** Second ***")
	println(s1.Value())
	println(i1.Value())
	println(b1.Value())
	println(a1.Value())
}
