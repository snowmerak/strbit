package strbit

func BinString2Bits(s string) []byte {
	result := []byte{}
	i := 0
	for _, c := range s {
		if i%8 == 0 {
			result = append(result, 0)
		}
		if c == '1' {
			result[i/8] |= 1 << uint(7-i%8)
		}
		i++
	}
	return result
}
