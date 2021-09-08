package strbit

func BinString2Bits(s string, r []byte) {
	i := 0
	for _, c := range s {
		if i%8 == 0 {
			r[i/8] = 0
		}
		if c == '1' {
			r[i/8] |= 1 << uint(7-i%8)
		}
		i++
	}
}

func BinBits2String(b []byte) string {
	result := ""
	for i := 0; i < len(b); i++ {
		for j := 7; j >= 0; j-- {
			if b[i]&(1<<uint(j)) != 0 {
				result += "1"
			} else {
				result += "0"
			}
		}
	}
	return result
}
