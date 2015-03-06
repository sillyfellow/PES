package primemath

import (
	"math"

	"github.com/sillyfellow/PES/golib/setutils"
)

//if % is absent, use this.
func is_k_divisor(n int, k int) bool {
	return n == (n/k)*k
}

func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//gets any one prime divisor of n
//during it does it, the primes will be filled up with the primes which would be tested for suitable candidacy
func get_prime_divisor(n int, primes *setutils.IntSet) int {
	if primes.Contains(n) {
		return 0
	}

	//we don't need to go beyond this
	max_to_go := int(math.Ceil(math.Sqrt(float64(n))))

	//for squares of primes, we go one extra step
	if max_to_go*max_to_go == n {
		max_to_go += 1
	}

	//go through the set of primes first. Note the largest one we encountered
	largest := int(2)
	for k := range primes.Set {
		if k > max_to_go {
			continue
		}

		if n%k == 0 {
			return k
		}

		largest = Max(k, largest)
	}

	// prevent the loophole in case the primes was initially filled with some
	// random primes, which are not the smallest primes.

	//from that largest one onwards, we need to check each of the numbers until
	//we reach the max to go
	for ; largest < max_to_go; largest++ {
		//if the number is not a prime, we could let it go
		if get_prime_divisor(largest, primes) != 0 {
			continue
		}
		//otherwise, add it to the list, and check for being a divisor
		primes.Add(largest)
		if n%largest == 0 {
			return largest
		}
	}
	return 0
}

func update_divisor_count(divisor int, divisors *map[int]int) {
	count := (*divisors)[divisor]
	(*divisors)[divisor] = count + 1
}

func find_primes_and_divisors(n int, primes *setutils.IntSet, divisors *map[int]int) {
	pd := get_prime_divisor(n, primes)
	if pd == 0 { // the input is indeed a prime
		update_divisor_count(n, divisors)
		return
	}

	//print the one we found, and look for the divisors of the remaining quotient.
	update_divisor_count(pd, divisors)
	find_primes_and_divisors(n/pd, primes, divisors)
}

func FindDivisors(n int) *map[int]int {
	divisors := make(map[int]int)
	find_primes_and_divisors(n, setutils.NewIntSet(), &divisors)
	return &divisors
}
