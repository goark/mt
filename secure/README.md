# About "secure.Source" type

`secure.Source` type is wrapper of crypto/rand package.

## Usage

```go
//go:build run
// +build run

package main

import (
    "fmt"
    "math/rand"

    "github.com/spiegel-im-spiegel/mt"
    "github.com/spiegel-im-spiegel/mt/secure"
)

func main() {
    fmt.Println(rand.New(secure.Source{}).Uint64())
    fmt.Println(mt.New(secure.Source{}).Uint64())
}
```