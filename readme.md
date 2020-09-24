lazy
====

Provides support for lazy initialization for bool/int/string/interface{}

This package is made by [genny](https://github.com/cheekybits/genny/).

example
-------

```go
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
```

```shell
$ go run example.go
*** First ***
s1 initialize
Foo
i1 initialize
1
b1 initialize
true
a1 initialize
(0xf99260,0xfbc1a0)
** Second ***
Foo
1
true
(0xf99260,0xfbc1a0)
```
