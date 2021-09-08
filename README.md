# strbit

## hex

`strbit.HexString2Bits()` is a function that converts a hex string to a byte array.

`strbit.HexBits2String()` is a function that converts a byte array to a hex string.

## bin

`strbit.BinString2Bits()` is a function that converts a binary string to a byte array.

`strbit.BinBits2String()` is a function that converts a byte array to a binary string.

## example

```go
package main

import (
	"fmt"

	"github.com/snowmerak/strbit"
)

func main() {
	x := "ab1f0e83"
	fmt.Println([]byte(x))
	r := [32]byte{}
	strbit.HexString2Bits(x, r[:])
	fmt.Println(r)

	b := "1110001010010111"
	fmt.Println([]byte(b))
	r = [32]byte{}
	strbit.BinString2Bits(b, r[:])
	fmt.Println(r)
}

```

```bash
[97 98 49 102 48 101 56 51]
[171 31 14 131 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
[49 49 49 48 48 48 49 48 49 48 48 49 48 49 49 49]
[226 151 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
```