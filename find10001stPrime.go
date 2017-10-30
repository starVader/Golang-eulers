package main

import "fmt"

const maxSieve = 150000
const targetPrime = 10001

func main() {
	var primes []int
	nonPrimes := make(map[int]bool)

	for i := 2; i <= maxSieve && len(primes) < targetPrime; i++ {
		if nonPrimes[i] != true {
			primes = append(primes, i)
			for j := i; j <= maxSieve; j = j + i {
				nonPrimes[j] = true
			}
		}
	}

	if len(primes) < targetPrime {
		fmt.Println("Target prime not found yet, please increase maxSieve")
		fmt.Printf("Target Prime: %d\n", targetPrime)
		fmt.Printf("Found Primes: %d\n", len(primes))
	} else {
		fmt.Printf("Result: %d\n", primes[targetPrime-1])
	}
}
