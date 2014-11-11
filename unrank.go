package unrank

import (
	"errors"
)

var (
	ErrCmp = errors.New("Unrank: comparation failed")
	ErrNul = errors.New("Unrank: zeroes underrun")
	ErrOne = errors.New("Unrank: ones underrun")
)

const (
	debug = 1
)

// Binomialer provides tresholders for the x,y field from the pascal's triangle
// If the x or y == 0 it gives the number 0
// If the x or y == 1 it gives 1
// If the x or y == 2 it gives y or x, etc..
type Binomialer interface {
	Binomial(uint, uint) Tresholder
}

// Tresholder is typically an unsigned number
// The Cmp() compares itself to Tresholder and returns -1 iff this < Tresholder
// Cmp() may return -2 if the comparation failed
// The Sub() makes this smaller by argument
type Tresholder interface {
	Cmp(Tresholder) int
	Sub(Tresholder)
}

// Type set represents the result bit set. The SetBit set's 
type Set interface {
	SetBit(int, b uint)
}

// Unrank implements the combination unrank algorithm
func Unrank(s Set, b Binomialer, tr Tresholder, z1 uint, z0 uint) error {
	var n uint

	for z0 + z1 > 0 {
		bi := b.Binomial(z0, z1 + 1)
		cmp := tr.Cmp(bi)

		if cmp == -2 {
			return ErrCmp
		} else if cmp < 0 {
			s.SetBit(n, 0)

			if z0 == 0 {
				return ErrNul
			}
			z0--
		} else {
			s.SetBit(n, 1)
			tr.Sub(bi)

			if z1 == 0 {
				return ErrOne
			}
			z1--
		}

		n++;
	}

	return nil
}
