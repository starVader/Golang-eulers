package main

//Find the largest palindrome made from the product of two 3-digit numbers.
import (
	"strconv"
	"fmt"
)

func checkPallindrome(first, second int) bool {
	var mul string = strconv.Itoa(first * second)
	var flag bool
	for i := 0; i < len(mul); i++ {
		if mul[i] != mul[len(mul)-i-1] {
			flag = true
			break
		}
	}
	if flag {
		return false
	}
	return true
}

func main() {
	var largest int
	largest = 100001
	for j := 999; j > 100; j--{
		for i:= 999;i >= 100; i-- {
			if largest < i*j  && checkPallindrome(i, j) {
				largest = i * j
			}
		}
	}
	fmt.Print(largest)

}
