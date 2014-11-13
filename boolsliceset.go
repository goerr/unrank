package unrank

type BoolSliceSet []bool

func (s *BoolSliceSet) SetBit(n int, b uint) {
	for n >= len(*s) {
		*s = append(*s, false)
	}
	(*s)[n] = (b == 1)
}
