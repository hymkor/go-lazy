lazy
====

[![Build Status](https://travis-ci.com/hymkor/go-lazy.svg?branch=master)](https://travis-ci.com/github/hymkor/go-lazy)
[![Go Reference](https://pkg.go.dev/badge/github.com/hymkor/go-lazy.svg)](https://pkg.go.dev/github.com/hymkor/go-lazy)

Provides support for lazy initialization by generics in Go1.18

example 1
---------

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

example 2
---------

Same as example 1. Light but long

```go
package main

import (
    "github.com/hymkor/go-lazy"
)

var s1 = lazy.Of[string]{
    New: func() string {
        println("s1 initialize")
        return "Foo"
    },
}

func main() {
    println(s1.Value())
    println(s1.Value())
}
```
