package sieve

import (
	"math/big"
	"testing"
	"fmt"
)

func TestSimple(t *testing.T) {

	var s Sieve = Sieve{*big.NewInt(9999999999999999)}

	var x,y uint
	for x = 0; x < 10; x++ {
	for y = 0; y < 7; y++ {

	fmt.Println(s.CmpSub(x,y))

	}}
}
