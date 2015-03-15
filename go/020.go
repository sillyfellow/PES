/*
 *
 *n! means n × (n − 1) × ... × 3 × 2 × 1
 *
 *For example, 10! = 10 × 9 × ... × 3 × 2 × 1 = 3628800,
 *and the sum of the digits in the number 10! is 3 + 6 + 2 + 8 + 8 + 0 + 0 = 27.
 *
 *Find the sum of the digits in the number 100!
 */

package main

import (
	"fmt"
	// I'm using this opportunity to learn to use some imported, new type!
	"math/big"
)

func biggy(n int64) *big.Int {
	return big.NewInt(n)
}

// Actually, unnecessary when we use bigInt  -- (ss@03/15/2015) --
func removeTrailingZeros(n *big.Int) *big.Int {
	for biggy(10).Cmp(new(big.Int).Rem(n, biggy(10))) == 0 {
		n = new(big.Int).Quo(n, biggy(10))
	}
	return n
}

func factorialNSansZero(n *big.Int) *big.Int {
	if n.Cmp(biggy(0)) == 0 {
		return biggy(1)
	}
	return new(big.Int).Mul(n, removeTrailingZeros(factorialNSansZero(new(big.Int).Sub(n, biggy(1)))))
}

func sumOfDigits(n *big.Int) int64 {
	sum := int64(0)
	for n.Cmp(biggy(0)) != 0 {
		sum += new(big.Int).Rem(n, biggy(10)).Int64()
		n = n.Quo(n, biggy(10))
	}
	return sum
}

func main() {
	fmt.Println(sumOfDigits(factorialNSansZero(biggy(100))))
	fmt.Println(factorialNSansZero(biggy(100)))
}
