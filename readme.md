[![Buil Status](https://travis-ci.com/zetamatta/go-lazy.svg?branch=master)](https://travis-ci.com/github/zetamatta/go-lazy)

lazy
====

Provides support for lazy initialization by generics in Go1.18

example
-------

```go
package main

import (
    "github.com/hymkor/go-lazy"
)

var s1 = lazy.New(func() string {
    println("s1 initialize")
    return "Foo"
})

func main() {
    println(s1.Value())
    println(s1.Value())
}
```

```shell
$ go run example.go
s1 initialize
Foo
Foo
```
