package main

//Find the largest palindrome made from the product of two 3-digit numbers.
import (
	"fmt"
	"strconv"
)

func checkPallindrome(first,second int) bool {
	var mul string = strconv.Itoa(first*second)
	var flag bool
	for i := 0;i < len(mul);i++{
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
	first := 999
	second := 999
	for second > 0{
		if checkPallindrome(first,second) {
			fmt.Println(first * second)
			break
		}
		second--
	}
}

