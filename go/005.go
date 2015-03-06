/*
 *2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
 *
 *What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
 */

package main

import (
	"fmt"
	"math"

	"github.com/sillyfellow/PES/golib/primemath"
)

func consolidateFactors(start, end int) *map[int]int {
	var consolidated = make(map[int]int)
	for ; start <= end; start++ {
		prime_factors := primemath.FindDivisors(start)
		for n, count := range *prime_factors {
			consolidated[n] = primemath.Max(count, consolidated[n])
		}
	}
	return &consolidated
}

func lcmOfFirstN(n int) int {
	var factors = consolidateFactors(2, n)
	lcm := 1
	for n, count := range *factors {
		lcm *= int(math.Pow(float64(n), float64(count)))
	}
	return lcm
}

func main() {
	fmt.Println(lcmOfFirstN(10))
	fmt.Println(lcmOfFirstN(20))
}
