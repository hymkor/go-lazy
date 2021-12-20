[![Buil Status](https://travis-ci.com/zetamatta/go-lazy.svg?branch=master)](https://travis-ci.com/github/zetamatta/go-lazy)

lazy
====

Provides support for lazy initialization by generics in Go1.18beta1

example
-------

```go
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
```

```shell
$ go1.18beta1.exe run example.go
*** First ***
s1 initialize
Foo
i1 initialize
1
b1 initialize
true
[]byte initialize
abcd
** Second ***
Foo
1
true
abcd
```
