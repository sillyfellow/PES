/*
 *The prime factors of 13195 are 5, 7, 13 and 29.
 *
 *What is the largest prime factor of the number 600851475143 ?
 */
package main

import "fmt"
import "math"
import "github.com/sillyfellow/PES/golib"

//if % is absent, use this.
func is_k_divisor(n int, k int) bool {
	return n == (n/k)*k
}

func my_max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func my_min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//gets any one prime divisor of n
//during it does it, the primes will be filled up with the primes which would be tested for suitable candidacy
func get_prime_divisor(n int, primes *golib.IntSet) int {
	if primes.Contains(n) {
		return 0
	}

	//we don't need to go beyond this
	max_to_go := int(math.Ceil(math.Sqrt(float64(n))))

	//go through the set of primes first. Note the largest one we encountered
	largest := int(2)
	for k := range primes.Set {
		if k > max_to_go {
			continue
		}

		if n%k == 0 {
			return k
		}

		largest = my_max(k, largest)
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

func find_divisors(n int, primes *golib.IntSet) {
	pd := get_prime_divisor(n, primes)
	if pd == 0 { // the input is indeed a prime
		fmt.Println(n)
		return
	}

	//print the one we found, and look for the divisors of the remaining quotient.
	fmt.Println(pd)
	find_divisors(n/pd, primes)
}

func main() {
	find_divisors(600851475143, golib.NewIntSet())
}
