/*
 *The prime factors of 13195 are 5, 7, 13 and 29.
 *
 *What is the largest prime factor of the number 600851475143 ?
 */
package main

import "fmt"
import "github.com/sillyfellow/PES/golib/primemath"

func main() {
	var divisors = primemath.FindDivisors(600851475143)
	fmt.Println(*divisors)
}
