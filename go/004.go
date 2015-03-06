/*
 *A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.
 *
 *Find the largest palindrome made from the product of two 3-digit numbers.
 */

package main

import "fmt"
import "math"

func reverseNumber(n int) int {
	reverse := 0
	for n > 0 {
		reverse = (reverse * 10) + n%10
		n /= 10
	}
	return reverse
}

func isNumberPalindrome(n int) bool {
	if n < 0 {
		return isNumberPalindrome(-n)
	}
	return n == reverseNumber(n)
}

func nDigitBoundaries(n int) (int, int) {
	return int(math.Pow(10, float64(n))) - 1, int(math.Pow(10, float64(n-1)))
}

func largestNDigitProductPalindromes(n int) int {
	largest2NDigitNumber, smallest2NDigitNumber := nDigitBoundaries(n * 2)

	for i := largest2NDigitNumber; i > smallest2NDigitNumber; i-- {
		if !isNumberPalindrome(i) {
			continue
		}
		l, s := nDigitBoundaries(n)
		for ; s <= l; s++ {
			if i%s == 0 && (i/s) <= l {
				return i
			}
		}
	}
	return 0
}

func main() {
	fmt.Println(largestNDigitProductPalindromes(2))
	fmt.Println(largestNDigitProductPalindromes(3))
}
