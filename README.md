lazy
====

[![Build Status](https://travis-ci.com/hymkor/go-lazy.svg?branch=master)](https://travis-ci.com/github/hymkor/go-lazy)
[![Go Reference](https://pkg.go.dev/badge/github.com/hymkor/go-lazy.svg)](https://pkg.go.dev/github.com/hymkor/go-lazy)

Provides support for lazy initialization by generics in Go1.18

example 1
---------

```example1.go
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
```

```go run example1.go|
start
s1 initialize
Foo
Foo
```

example 2
---------

Same as example 1. Light but long

```example2.go
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
    println("start")
    println(s1.Value())
    println(s1.Value())
}
```

example 3
---------

Two values version like `"sync".OnceValues`

```example3.go
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
```

```go run example3.go|
start
s1 initialize
Foo 1
Foo 1
```

