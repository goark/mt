# [goark/mt][github.com/goark/mt/v2] -- [Mersenne Twister]; Pseudo Random Number Generator, Implemented by [Golang][Go]

[![check vulns](https://github.com/goark/mt/workflows/vulns/badge.svg)](https://github.com/goark/mt/actions)
[![lint status](https://github.com/goark/mt/workflows/lint/badge.svg)](https://github.com/goark/mt/actions)
[![GitHub license](http://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/goark/mt/master/LICENSE)
[![GitHub release](http://img.shields.io/github/release/goark/mt.svg)](https://github.com/goark/mt/releases/latest)

This package is "[Mersenne Twister]" algorithm, implemented by pure [Go].

- required Go 1.22 or later
- Compatible with [math/rand/v2] standard package.
- Concurrency-safe (if it uses [mt][github.com/goark/mt/v2].PRNG type)

**Migrated repository to [github.com/goark/mt][github.com/goark/mt/v2]**

## Usage

### Import

```go
import "github.com/goark/mt/v2"
```

### Usage with [math/rand/v2] Standard Package (not concurrency-safe)

```go
package main

import (
    "fmt"
    "math/rand/v2"

    "github.com/goark/mt/v2/mt19937"
)

func main() {
    fmt.Println(rand.New(mt19937.New(19650218)).Uint64())
    //Output:
    //13735441942630277712
}
```

### Usage of [mt][github.com/goark/mt/v2].PRNG type (concurrency-safe version)

```go
import (
    "fmt"

    "github.com/goark/mt/v2"
    "github.com/goark/mt/v2/mt19937"
)

fmt.Println(mt.New(mt19937.New(19650218)).Uint64())
//Output:
//13735441942630277712
```

#### Use [io].Reader interface

```go
import (
    "encoding/binary"
    "fmt"
    "math/rand/v2"
    "sync"

    "github.com/goark/mt/v2"
    "github.com/goark/mt/v2/mt19937"
)

func main() {
    wg := sync.WaitGroup{}
    prng := mt.New(mt19937.New(rand.Int64()))
    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            r := prng.NewReader()
            buf := [8]byte{}
            for i := 0; i < 10000; i++ {
                ct, err := r.Read(buf[:])
                if err != nil {
                    return
                }
                fmt.Println(binary.LittleEndian.Uint64(buf[:ct]))
            }
        }()
    }
    wg.Wait()
}
```

## Benchmark Test

```
$ go test -bench Random -benchmem ./benchmark
goos: linux
goarch: amd64
pkg: github.com/goark/mt/v2/benchmark
cpu: AMD Ryzen 5 PRO 4650G with Radeon Graphics
BenchmarkRandomPCG-12              	785103775	         2.031 ns/op	       0 B/op	       0 allocs/op
BenchmarkRandomMT19917-12          	338082381	         3.551 ns/op	       0 B/op	       0 allocs/op
BenchmarkRandomPCGRand-12          	359948874	         3.288 ns/op	       0 B/op	       0 allocs/op
BenchmarkRandomMT19917Rand-12      	325159622	         3.687 ns/op	       0 B/op	       0 allocs/op
BenchmarkRandomChaCha8Locked-12    	186311572	         6.443 ns/op	       0 B/op	       0 allocs/op
BenchmarkRandomMT19917Locked-12    	128465040	         9.346 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/goark/mt/v2/benchmark	10.408s
```

## License

This package is licensed under MIT license.

- [Commercial Use of Mersenne Twister](http://www.math.sci.hiroshima-u.ac.jp/~m-mat/MT/MT2002/elicense.html)
- [Mersenne Twisterの商業利用について](http://www.math.sci.hiroshima-u.ac.jp/~m-mat/MT/MT2002/license.html)

[github.com/goark/mt/v2]: https://github.com/goark/mt "goark/mt: Mersenne Twister; Pseudo Random Number Generator, Implemented by Golang"
[Go]: https://go.dev/ "The Go Programming Language"
[math/rand/v2]: https://pkg.go.dev/math/rand/v2 "rand package - math/rand/v2 - Go Packages"
[io]: https://pkg.go.dev/io "io package - io - Go Packages"
[Mersenne Twister]: http://www.math.sci.hiroshima-u.ac.jp/~m-mat/MT/emt.html "Mersenne Twister: A random number generator (since 1997/10)"
