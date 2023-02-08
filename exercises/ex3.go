package exercises

func isPrime(n int) bool {
	if n <= 2 { // 0, 1 and 2 are not prime numbers
		return false
	}

	for i := 2; i < n/2; i++ {
		if n%i == 0 {
			// if rest of division is equal to zero, it means that the number has
			// a divider that's not 1 and itself
			return false // is not prime
		}
	}

	return true
}

func LargestPrimeFactor() {
}
