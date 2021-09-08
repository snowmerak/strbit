package strbit

import "fmt"

func HexString2Bits(s string, r []byte) {
	f := true
	i := 0
	for _, c := range s {
		x := byte(c)
		switch {
		case '0' <= c && c <= '9':
			x -= '0'
		case 'a' <= c && c <= 'f':
			x -= 'a' - 10
		case 'A' <= c && c <= 'F':
			x -= 'A' - 10
		default:
			panic("HaxString2Bits: invalid hex char")
		}
		if f {
			r[i] |= x << 4
			f = false
		} else {
			r[i] |= x
			i++
			f = true
		}
	}
}

func HexBits2String(bits []byte) string {
	result := ""
	for _, b := range bits {
		result += fmt.Sprintf("%02x", b)
	}
	return result
}
