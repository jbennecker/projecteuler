package main

/*
The sum of the squares of the first ten natural numbers is,

1^2 + 2^2 + ... + 10^2 = 385
The square of the sum of the first ten natural numbers is,

(1 + 2 + ... + 10)^2 = 552 = 3025
Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025 − 385 = 2640.

Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.
*/

import (
	"fmt"
)

const N = 100

func main() {

	sum1, sum2, i := 0, 0, 1

	sum1 = 0
	for i = 1; i <= N; i++ {
		sum1 += i * i
	}

	fmt.Println(sum1)

	sum2 = 0
	for i = 1; i <= N; i++ {
		sum2 += i
	}
	sum2 = sum2 * sum2

	diff := sum2 - sum1

	fmt.Println(diff)
}
