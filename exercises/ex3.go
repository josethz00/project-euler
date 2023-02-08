package exercises

import "fmt"

func isPrime(n int64) bool {
	if n < 2 { // 0, 1 and 2 are not prime numbers
		return false
	}

	var i int64

	for i = 2; i <= n/2; i++ {
		if n%i == 0 {
			// if rest of division is equal to zero, it means that the number has
			// a divider that's not 1 and itself
			return false // is not prime
		}
	}

	return true
}

func LargestPrimeFactor() {
	fixedN := 20
	n := 20
	i := 2

	for i < n {
		for n%i == 0 {
			n = n / i
		}
		i++
	}
	fmt.Printf("The largest prime factor for %d is: %d \n", fixedN, i)
}
