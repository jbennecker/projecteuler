package main

/*
By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.

What is the 10,001st prime number?
*/

import (
	"fmt"
	"math"
)

const N = 10001

func isPrime(x int) bool {
	var i int
	for i = 2; i <= int(math.Sqrt(float64(x))); i++ {
		if x%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	var primeCount int = 1
	var i int = 2

	for {
		if isPrime(i) {
			fmt.Printf("%d: %d\n", primeCount, i)
			primeCount++
		}

		if primeCount > N {
			break
		}

		i++
	}

}
