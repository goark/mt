# [spiegel-im-spiegel/mt] -- [Mersenne Twister]; Pseudo Random Number Generator, Implemented by [Golang]

[![Build Status](https://travis-ci.org/spiegel-im-spiegel/mt.svg?branch=master)](https://travis-ci.org/spiegel-im-spiegel/mt)
[![GitHub license](http://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/spiegel-im-spiegel/mt/master/LICENSE)
[![GitHub release](http://img.shields.io/github/release/spiegel-im-spiegel/mt.svg)](https://github.com/spiegel-im-spiegel/mt/releases/latest)

This package is "[Mersenne Twister]" algorithm, implemented by pure [Go].

- Compatible with [math/rand] standard package.
- Goroutine-safe (if it uses [mt].PRNG type)

## Usage

### Usage with [math/rand] Standard Package (not goroutine-safe)

```go
import (
    "fmt"
    "math/rand"

    "github.com/spiegel-im-spiegel/mt/mt19937"
)

func ExampleMT19937() {
    fmt.Println(rand.New(mt19937.NewSource(19650218)).Uint64())
    //Output:
    //13735441942630277712
}
```

### Usage of [mt].PRNG type (goroutine-safe version)

```go
import (
    "fmt"

    "github.com/spiegel-im-spiegel/mt"
    "github.com/spiegel-im-spiegel/mt/mt19937"
)

func ExampleMT() {
    fmt.Println(mt.New(mt19937.NewSource(19650218)).Uint64())
    //Output:
    //13735441942630277712
}
```

#### Use [io].Reader interface

```go
package main

import (
    "fmt"
    "sync"

    "github.com/spiegel-im-spiegel/mt"
    "github.com/spiegel-im-spiegel/mt/mt19937"
)

func main() {
    prng := mt.New(mt19937.NewSource(19650218))
    wg := sync.WaitGroup{}
    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            r := prng.NewReader()
            for i := 0; i < 10000; i++ {
                buf := [8]byte{}
                _, err := r.Read(buf[:])
                if err != nil {
                    return
                }
            }
        }()
    }
    wg.Wait()
}
```

## License

This paclage is licensed under MIT license.

- [Commercial Use of Mersenne Twister](http://www.math.sci.hiroshima-u.ac.jp/~m-mat/MT/MT2002/elicense.html)
- [Mersenne Twisterの商業利用について](http://www.math.sci.hiroshima-u.ac.jp/~m-mat/MT/MT2002/license.html)

[spiegel-im-spiegel/mt]: https://github.com/spiegel-im-spiegel/mt "spiegel-im-spiegel/mt: Mersenne Twister; Pseudo Random Number Generator, Implemented by Golang"
[mt]: https://github.com/spiegel-im-spiegel/mt "spiegel-im-spiegel/mt: Mersenne Twister; Pseudo Random Number Generator, Implemented by Golang"
[Go]: https://golang.org/ "The Go Programming Language"
[Golang]: https://golang.org/ "The Go Programming Language"
[math/rand]: https://golang.org/pkg/math/rand/ "rand - The Go Programming Language"
[io]: https://golang.org/pkg/io/ "io - The Go Programming Language"
[Mersenne Twister]: http://www.math.sci.hiroshima-u.ac.jp/~m-mat/MT/emt.html "Mersenne Twister: A random number generator (since 1997/10)"
