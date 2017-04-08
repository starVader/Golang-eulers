package main
//Fibonacci sequence whose values do not exceed four million, find the sum of the even-valued terms.
import (
	"fmt"
)

func fibonacciSum() int  {
	var sum int
	a := 0
	b := 0
	c := 1

	for c < 4000000 {
		a = b
		b = c
		c = a+b
		if c%2 == 0 {
			sum += c
		}
	}
	return sum
}

func main() {
	fmt.Println(fibonacciSum())

}

