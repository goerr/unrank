package sieve

import (
	"math/big"

	"github.com/soniakeys/integer/binomial"
)

/*
This file is the tresholder using the fast sieve Binomial function
*/

type Sieve struct {
	big.Int
}

func (b *Sieve) CmpSub(xx, yy uint) (int, interface{}) {

	var c big.Int
	binomial.Binomial(&c, xx + yy, yy)

	cmp := b.Cmp(&c)

	if cmp != -1 {
		b.Sub(&c)
		return 1, c
	}

	return -1, 0
}

func (b *Sieve) Sub(c interface{}) {
	b.Sub(&c)
}
