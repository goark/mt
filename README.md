# [spiegel-im-spiegel/mt] -- [Mersenne Twister]; Pseudo Random Number Generator, Implemented by [Golang]

[![Build Status](https://travis-ci.org/spiegel-im-spiegel/mt.svg?branch=master)](https://travis-ci.org/spiegel-im-spiegel/mt)
[![GitHub license](http://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/spiegel-im-spiegel/mt/master/LICENSE)
[![GitHub release](http://img.shields.io/github/release/spiegel-im-spiegel/mt.svg)](https://github.com/spiegel-im-spiegel/mt/releases/latest)

This package is "[Mersenne Twister]" algorithm, implemented by pure [Go].

- Compatible with [math/rand] standard package.

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

### Usage with [mt].PRNG type (goroutine-safe version)

```go
import (
    "encoding/binary"
    "fmt"

    "github.com/spiegel-im-spiegel/mt"
    "github.com/spiegel-im-spiegel/mt/mt19937"
)

func ExampleMT() {
    prng := mt.New(mt19937.NewSource(19650218))
    r := prng.Open()
    defer r.Close()

    buf := [8]byte{}
    ct, err := r.Read(buf[:])
    if err != nil {
        return
    }
    fmt.Println(binary.LittleEndian.Uint64(buf[:ct]))
    //Output:
    //13735441942630277712
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
[Mersenne Twister]: http://www.math.sci.hiroshima-u.ac.jp/~m-mat/MT/emt.html "Mersenne Twister: A random number generator (since 1997/10)"
