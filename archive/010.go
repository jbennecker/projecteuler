package main

/*
The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.

Find the sum of all the primes below two million.
*/

import (
	"fmt"
	"math"
)

const N = 2000000

func isPrime(x int) bool {
	if x == 1 {
		return false
	}
	var i int
	for i = 2; i <= int(math.Sqrt(float64(x))); i++ {
		if x%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	sum := 0
	for i := 0; i < N; i++ {
		if isPrime(i) {
			sum += i
		}
	}

	fmt.Println(sum)
}
