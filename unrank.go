package unrank

import (
	"errors"
)

var (
	ErrUnr = errors.New("Unrank: Treshold error")
	ErrNul = errors.New("Unrank: Zeroes underrun")
	ErrOne = errors.New("Unrank: Ones underrun")
)

// CmpSub compares the treshold to the binomial number with parameters xx, yy.
// If the treshold is larger or equal, number is substracted from treshold.
// Returns -1 iff treshold < binomial (not-substracted)
// Returns -2 iff an error occured (not substracted)
// Else it substracts and returns the binomial coefficient in any format.
type Tresholder interface {
	CmpSub(xx, yy uint) (int, interface{})
	Sub(interface{})
}

// Set represents the actual underlying result bit set.
// The n-th bit is set to b == {0,1}
type Set interface {
	SetBit(n uint, b uint)
}

// Unrank implements the combination unrank algorithm
// The unrank treshold is provided as part of the Tresholder argument
// The result is stored to the Set argument
func Unrank(s *Set, t Tresholder, z1 uint, z0 uint) error {
	var n, bit, swap uint

	for z0 + z1 > 0 {
		// a cool swap to minimize subtractions
		if z1 > z0 {
			bit = 1-bit
			swap = z0
			z0 = z1
			z1 = swap
		}

		cmp, _ := t.Treshold(z0, z1 + 1)

		switch cmp {
		case -2:
			return ErrUnr
		case -1:
			s.SetBit(n, bit)

			if z0 == 0 {
				return ErrNul
			}
			z0--
		default:
			s.SetBit(n, 1-bit)

			if z1 == 0 {
				return ErrOne
			}
			z1--
		}

		n++;
	}

	return nil
}
