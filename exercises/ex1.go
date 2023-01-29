package exercises

import (
	"fmt"
	"time"
)

func SumOf3And5Multiples() {
	fmt.Println()
	fmt.Println(
		`If we list all the natural numbers below 10 that are multiples
of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.
Find the sum of all the multiples of 3 or 5 below 1000. (Project Euler #1)`,
	)
	fmt.Println()

	start := time.Now()

	k3 := 999 / 3   // 333
	k5 := 995 / 5   // 199
	k15 := 990 / 15 // 66 - intersection

	sum := 3 * ((k3 * (k3 + 1)) / 2)
	sum += 5 * ((k5 * (k5 + 1)) / 2)
	sum -= 15 * ((k15 * (k15 + 1)) / 2)

	elapsed := time.Since(start)

	fmt.Printf("Sum of 3 and 5 multiples below 1000: %d\n", sum)
	fmt.Print("Execution time: ", elapsed)
}
