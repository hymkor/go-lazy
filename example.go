//+build ignore

package main

import (
	"github.com/zetamatta/go-lazy"
)

var s1 = lazy.Of[string]{
	New: func() string {
		println("s1 initialize")
		return "Foo"
	},
}

var i1 = lazy.Of[int]{
	New: func() int {
		println("i1 initialize")
		return 1
	},
}

var b1 = lazy.Of[bool]{
	New: func() bool {
		println("b1 initialize")
		return true
	},
}

var B = lazy.Of[[]byte]{
	New: func() []byte {
		println("[]byte initialize")
		return []byte{'a', 'b', 'c', 'd'}
	},
}

func main() {
	println("*** First ***")
	println(s1.Value())
	println(i1.Value())
	println(b1.Value())
	println(string(B.Value()))
	println("** Second ***")
	println(s1.Value())
	println(i1.Value())
	println(b1.Value())
	println(string(B.Value()))
}
