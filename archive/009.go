package main

/*
A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,

a^2 + b^2 = c^2
For example, 3^2 + 4^2 = 9 + 16 = 25 = 5^2.

There exists exactly one Pythagorean triplet for which a + b + c = 1000.
Find the product abc.
*/

import (
	"fmt"
)

const sum = 1000

func main() {

loop:
	for a := 1; a < 500; a++ {
		for b := a + 1; b < 500; b++ {
			var c int
			c = sum - a - b
			if a*a+b*b == c*c {
				fmt.Println(a, b, c, a*b*c)
				break loop
			}
		}
	}
}
