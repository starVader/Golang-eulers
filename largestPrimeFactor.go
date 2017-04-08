package main 

import "fmt"

/*
	

*/
func Find(A ,B ,C int) int {
	if A == 1{
		return C
	}
	if A % B == 0 {

		fmt.Println(A,B)
		if B  > C { //If the current Prime is greater than last prime (C) is greater  
			C = B // Replace the current greatest integer with the current one
			A = A/B // reduce A to half
			//fmt.Printf("A %d B %d C %d",A,B,C)
		}
		B = 2 //reset B to 2 and start again
	} else {
			//fmt.Printf("b %d ",B)
			B += 1 //Increment B if A is not divided
	}
	return Find(A,B,C)
}

func main() {
	A := 5666789004566788
	B := 2
	var C int
	fmt.Println(Find(A,B,C))
}