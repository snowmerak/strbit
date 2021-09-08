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
	fmt.Println(strbit.HexBits2String(r[:4]))

	b := "1110001010010111"
	fmt.Println([]byte(b))
	r = [32]byte{}
	strbit.BinString2Bits(b, r[:])
	fmt.Println(r)
	fmt.Println(strbit.BinBits2String(r[:2]))

	bs := []bool{true, true, false, false, true, false, true, false, true, true, false, false, false, true, true, false, true, true, false, false, true, false, true, false, true, true, false, false, false, true, true, false, true, true, false, false, true, false, true, false, true, true, false, false, false, true, true, false, true, true, false, false, true, false, true, false, true, true, false, false, false, true, true, false}
	r = [32]byte{}
	strbit.BoolArray2Bits(bs, r[:])
	fmt.Println(r)
	bs = make([]bool, 32)
	strbit.BoolBits2Array(r[:4], bs)
	fmt.Println(bs)
}
```

```bash
[97 98 49 102 48 101 56 51]
[171 31 14 131 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
ab1f0e83
[49 49 49 48 48 48 49 48 49 48 48 49 48 49 49 49]
[226 151 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
1110001010010111
[83 99 83 99 83 99 83 99 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
[true true false false true false true false true true false false false true true false true true false false true false true false true true false false false true true false]
```