package blind75

func reverseBits(num uint32) uint32 {
	res := uint32(0)
	for i := 0; i < 32; i++ {
		res <<= 1
		if num&1 == 1 {
			res++
		}
		num >>= 1
	}
	return res
}
