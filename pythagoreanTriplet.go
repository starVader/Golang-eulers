package main

import (
	"fmt"
	"math"
)

var total = 1000

func main() {
	result := -1

	for a := 1; a < total; a++ {
		for b, c := a, total-a-a; b < total && c > b; b, c = b+1, c-1 {
			actualC := math.Hypot(float64(a), float64(b))
			isPerfect := math.Pow(actualC, 2) == math.Pow(math.Floor(actualC), 2)
			if isPerfect && int(actualC) == c {
				fmt.Println(a, b, c)
				result = a * b * c
			}

			if result != -1 {
				break
			}
		}

		if result != -1 {
			break
		}
	}

	fmt.Println(result)
}
