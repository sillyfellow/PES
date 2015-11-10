/*
 10001st prime
 Problem 7
 By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see
 that the 6th prime is 13.

 What is the 10 001st prime number?
------------ (ss@10/11/15) */

package main

import (
	"fmt"
	"math"
)

func isPrime(n int, primes []int) bool {
	sqrt := 1 + math.Sqrt(float64(n))
	for _, p := range primes {
		if p > int(sqrt) {
			break
		}

		if n%p == 0 {
			return false
		}
	}
	return true
}

func generatePrimes(n int) []int {
	primes := make([]int, 0, n)
	primes = append(primes, 2)

	number := 3
	for ; len(primes) <= n; number++ {
		if isPrime(number, primes) {
			primes = append(primes, number)
		}
	}

	return primes
}

func main() {
	primes := generatePrimes(10001)
	length := len(primes)
	fmt.Println(primes[length-9 : length-1])
}
