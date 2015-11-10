/*
 Problem 6
 Published on Friday, 14th December 2001, 07:00 pm; Solved by 286585;
 Difficulty rating: 5%

 The sum of the squares of the first ten natural numbers is,

 1^2 + 2^2 + ... + 10^2 = 385

 The square of the sum of the first ten natural numbers is,

 (1 + 2 + ... + 10)^2 = 55^2 = 3025 Hence the difference between the sum of the
 squares of the first ten natural numbers and the square of the sum is 3025 −
 385 = 2640.

 Find the difference between the sum of the squares of the first one hundred
 natural numbers and the square of the sum.
------------ (ss@10/11/15) */

package main

import "fmt"

func sumOfSquares(n int) int {
	sum := 0
	for ; n > 0; n-- {
		sum += (n * n)
	}
	return sum
}

func squareOfSums(n int) int {
	sum := (n * (n + 1)) / 2
	return sum * sum
}

func main() {
	fmt.Println(squareOfSums(10) - sumOfSquares(10))
	fmt.Println(squareOfSums(100) - sumOfSquares(100))
}
