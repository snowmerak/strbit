package strbit

import "fmt"

func HexString2Bits(s string) []byte {
	result := []byte{}
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
		}
		if f {
			result = append(result, x<<4)
			f = false
		} else {
			result[i] |= x
			i++
			f = true
		}
	}
	return result
}

func HexBits2String(bits []byte) string {
	result := ""
	for _, b := range bits {
		result += fmt.Sprintf("%02x", b)
	}
	return result
}
