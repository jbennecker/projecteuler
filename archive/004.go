package main

/*
A palindromic number reads the same both ways.
The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.

Find the largest palindrome made from the product of two 3-digit numbers.
*/

import (
	"fmt"
	"strconv"
)

func isPalindrome(n int) bool {

	s := strconv.Itoa(n)
	sLen := int(len(s))

	for i := 0; i < sLen/2; i++ {
		if s[i] != s[sLen-i-1] {
			return false
		}
	}

	// was ist mit ungeraden stellen? 333
	return true
}

func main() {
	var a, b, c int

loop:
	for a = 999; a > 800; a-- {
		for b = 999; b > 800; b-- {
			c = a * b
			if isPalindrome(c) {
				// found it
				break loop
			}
		}
	}

	fmt.Println(c)
}
