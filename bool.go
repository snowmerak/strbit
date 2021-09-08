package strbit

func BoolArray2Bits(b []bool, r []byte) {
	for i, v := range b {
		if v {
			r[i/8] |= 1 << uint(i%8)
		}
	}
}

func BoolBits2Array(r []byte, b []bool) {
	for i, v := range r {
		for j := uint(0); j < 8; j++ {
			b[i*8+int(j)] = (v & (1 << j)) != 0
		}
	}
}
